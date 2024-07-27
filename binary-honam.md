---
title: "비 내리는 호남선 (binary-honam.live)"
date: 2024-07-26T13:29:14.170Z
---

# [[WIP]] 비 내리는 호남선 (binary-honam.live)

[[항해]]중이였다. [[목포]]에서 노래 가사 한 구절을 떠올렸다.

> 비내리는 호남선 남행열차에   
> 흔들리는 차창 너머로...   
> \- 김수희, 『남행열차』

문득 그런 생각이 들었다. 비 내리는 [[호남선]] 남행열차를 보여주는 사이트를 만들 수 있지 않을까 ...

binary-honam.live는 그렇게 시작된 프로젝트이다. [[Coupy]]에게 [[디자인]]을 맡기고자 하였으나, 깔깔 웃기만 하고 도망쳐버려서 나 혼자 만들게 되었다.

---

지금은 [[2024-w30|2024년 07월 26일]]이다.

## 컨셉 잡기

원래 이런 정신나간 사이트는 시각적으로 기깔나야 한다. 그 뒤에 어떤 기술이 쓰이던 상관 없지만, 정신나간 비주얼을 자랑하는 것이 중요하다.

![비 내리는 호남선 서비스의 UI 디자인 목업](../images/binary-honam-live-ui-mock.png)

기획이 너무 재밌던 나머지 디자인을 두시간만에 끝내버렸다. 에셋은 [[Remix Icons]], [[Shade UI]]를 사용하였다. 더 정신나간 비주얼 컨셉을 그리고 싶었지만, 이상한 감성은 아무래도 실력이 먼저 받쳐줘야 하는지라.

## 비 내리는 호남선 찾는 로직 구현하기

2024년 07월 26일 기준, 국내 간선 여객[[철도]]를 국가에서 운영함에 큰 감사를 느낀다. [[일본]]처럼 여객수송이 민영화되어 30개의 회사가 존재했다면 이 프로젝트는 시작하지도 못했으리라.

한국은 [[코레일]]과 [[SR]] 두 회사만이 호남선에서 여객수송을 운행하고 있다. 코레일은 실시간 열차 운행 계획 API를 제공하지만, SR은 그렇지 않다. 공개된 운행계획표를 [[JSON]]으로 재가공해서 정적 데이터로 서빙해야 한다.

그러다 문득 이런 생각이 든다.. [[여수]] 가는 기차는 호남선이 아닌가? 호남선은 서[[대전]]부터 [[목포]]라는데? [[서울]]에서 여수를 가면 호남선을 지난다면 이 또한 "호남선 남행열차"로 볼 수 있지 않을까? 이는.. 대답하기 쉽지 않은 질문일 것이다.. 열심히 찾아보니 서울에서 여수를 갈 때 호남선을 거치긴 하는 것 같다.. 사실 자신이 있지 않다 그냥 대충 알아봤는데

1. [[계룡역]]은 호남선의 역이다. 환승? 가능한 다른 노선은 없다.
2. [[임실역]]은 전라선의 역이다. 환승? 가능한 다른 노선은 없다.
3. [[용산역]]에서 임실역으로 갈 때 계룡역을 지난다
4. 서울에서 전라선의 역을 갈 때는 호남선을 지나게 된다 / 호남선을 지나야만 전라선을 갈 수 있다

그러나 우리는 여수를 갈 때 호남선을 탄다고 말하는가? 이는 마치 서울에서 [[창원]]을 갈 때 [[경부선]]을 탄다고 할 수 있는가?와 동일한 문제이다. 경부선을 타고 가긴 하지만.. 심지어 서울-창원간 철로의 대부분은 경부선이지만.. 그 누구도 이를 두고 "난 경부선을 탄다!"라고 하지 않는다..

[[나무위키]]에서 표현하기로는 호남선 이름을 붙힌 호남선 열차를 "순수 호남선 계통"이라고 칭하는 것 같다. 호남선 이름은 아니지만 호남선에 타는 열차는 "다른 노선 계통의 열차"라고 부른다. 참으로 어려운 문제이다. 로직을 구현할 때 "✅ 순수 호남선 계통만 보기" 옵션을 포함해서 구현해야겠다.

> (여담) 개인적으로 평소엔 나무위키를 전혀 읽지 않는다. [[NextDNS]] 설정에서 차단해둔 몇 안되는 사이트중 하나이다. 그러나 자료를 찾으면서 기차 관련한 자료는 나무위키에 꽤나 자세하게 설명되어 있다는 사실을 알았다. 이번 프로젝트에서만 조금 참고하고자 한다.

아무튼 기차를 전혀 모르는데 기차 [[도메인]]의 제품을 만드려니 골머리가 아프다.

---

코레일에서는 [[공공데이터포털]]을 통해 여객열차 운행 계획 [[API]]를 제공한다. 이걸 사용하면 코레일 열차는 쉽게 가져올 수 있을 줄 알았는데, 헛된 기대였다.

![여객열차 운행계획 API 고장](../images/korail-train-plan-api-broken.png)

전혀 운행계획 데이터가 아니였다. 이미 운행한 열차의 계획만 반환되고, 앞으로의 운행계획은 나오지 않았다. 내가 "운행 정보" API와 헷갈린줄 알고 여러번 시도해봤는데 동일했다. 용산 - 광주송정 / [[행신]] - [[부산]] 모두 테스트 해보았지만 전혀 다름이 없었다. 전혀 제대로 작동하지 않았다. 이렇다면 코레일과 SR 모두 운행 계획을 정적으로 입력해두고 쿼리해야 하는 수 밖에 없다.

대체 [[레일블루]]는 데이터를 어떻게 가져오는거야.

---

지금은 [[2024-w30|2024년 07월 27일]]이다.

### 코레일 일반열차 운행일정 파싱하기

코레일 홈페이지에서는 열차 시간표를 엑셀 파일로 제공한다. 이를 파싱해서 간단히 열차 운행 정보 데이터베이스를 만들어보자.

[열차운임 및 시간표 - letskorail.com](https://www.letskorail.com/ebizcom/cs/guide/guide/guide11.do)

위 게시판의 "KTX 시간표"와 "일번열차 시간표"를 참고하였다.

![코레일 호남선 일반열차 시간표](korail-normal-train-honam-timetable.png)

위와같이 거대한 시간표를 볼 수 있다. 간단히 파싱해보자. 손으로 직접 입력하는게 파싱하는 것 보다 쉬울 것 같지만 한시간 걸려서 할 일을 5분만에 끝내기 위해 두시간동안 코딩하는게 개발자의 덕목이기에, 이번 기회엔 보편가치를 따르고자 한다.

시발역, 중간 정차역, 종착역이 각 역 정차 시각과 함께 나와있다. 이중 중간 정차역은 필요하지 않기에, 시발역과 종착역, 그리고 운행 시각만 분석해보자. 엑셀에서는 데이터를 클립보드에 복사할 때 각 줄이 `\n`(개행)으로, 각 셀이 `\t`(탭)으로 구분되는 텍스트로 저장한다. 그렇기에 간단한 파서로 쉽게 데이터를 얻어올 수 있다.

전라선도 동일하게 파싱하였으나, 호남선을 경유하지 않고 익산에서 출발하는 일부 열차는 제외하여 용산발 열차만 데이터에 남겼다.

```json
{
    "departure": { "station": "용산", "time": "21:22" },
    "arrival": { "station": "익산", "time": "00:38" },
    "trainName": "무궁화",
    "trainNumber": "1443",
    "mainline": "호남선",
    "runningDay": [ 1, 2, 3, 4, 5, 6, 7 ]
  },
  {
    "departure": { "station": "용산", "time": "05:44" },
    "arrival": { "station": "여수엑스포", "time": "11:05" },
    "trainName": "무궁화",
    "trainNumber": "1501",
    "mainline": "전라선",
    "runningDay": [ 1, 2, 3, 4, 5, 6, 7 ]
  }
```

위와 같은 데이터 형식으로 저장하였다

### 코레일 고속열차 운행일정 파싱하기

고속열차 운행일정은 일반열차와 다른 포맷으로 제공된다. 그러나 구조화되어 엑셀 시트로 제공되기에, 테이블 파서는 동일하게 사용하고 값 분석 로직만 새로 작성하였다.

![코레일 호남선 고속열차 운행 일정](../images/korail-express-train-honam-timetable.png)

- 고속열차의 경우 요일마다 운행하는 열차가 다른 것 같다. 처음 알았다.. 이 또한 데이터에 저장해주었다.
- 시발역, 종착역이 명시되어 있지 않고, 각 역별 정차 시각만 나와있다. 행신발 / 서울발 / 용산발로 나뉘어 있어 첫 정차역으로 시발역을 알아냈다.
- 모든 전라선 하행 고속열차는 서울시/고양시 착발이기에 반드시 호남선을 경유한다. 그렇기에 별도로 데이터 선별은 거치지 않았다.

```json
{
    "departure": { "station": "용산", "time": "7:38" },
    "arrival": { "station": "광주송정", "time": "9:24" },
    "trainName": "KTX",
    "trainNumber": "441",
    "runningDay": [ 1, 2, 3, 4, 5, 6, 7 ],
    "mainline": "호남선"
}
```

동일한 포맷으로 저장하였다.

### SRT 운행일정 파싱하기

SRT는 운행일정을 엑셀로 제공하지 않는다. PDF로만 제공해서, 직접 눈으로 파싱했다. 역시 파서 작성보다 시간이 짧게 걸렸다. 

[열차운임 및 시간표 < 승차권이용안내 < 이용안내 < 승차권 예약/발매 - etk.srail.kr](https://etk.srail.kr/cms/archive.do?pageId=TK0402050000)

## 날씨 데이터 받아오기

만만한 기상청에서 받아오기로 했다. 기상청 단기 예보는 향후 50시간 동안의 기상 예보를 한시간 시격으로 제공한다.

https://www.data.go.kr/data/15084084/openapi.do

API가 저수준으로 제공되어 추상화가 필요하다. 간단한 재구성으로 사용하기 쉽게 가공할 수 있었다.

```
Map(4) {
  '20240727 1300' => { probability: '60', precipitation: '8.0mm' },
  '20240727 1500' => { probability: '60', precipitation: '6.0mm' },
  '20240727 1700' => { probability: '60', precipitation: '3.0mm' },
  '20240727 1800' => { probability: '60', precipitation: '7.0mm' }
}
```

이렇게 향후 50시간 내에 비가 오는 시간대를 받아온다.

```
export const principalStation = {
    용산: {
        x: 60,
        y: 126,
    },
    광주송정: {
        x: 57,
        y: 74,
    },
    서대전: {
        x: 68,
        y: 100,
    },
    목포: {
        x: 50,
        y: 66,
    },
}
```

기상청 API는 위경도 좌표계를 사용하지 않는다. 위경도와 유사한 자체 좌표계를 사용하는데, 행정구역별로 기상청 좌표를 계산해둔 엑셀 시트를 참고해서 개발하여야 한다. 위경도를 그대로 넣으면 안된다.

![행정구역별 기상청 API 좌표를 나열한 엑셀 시트](../images/kma-positioning-system.png)

호남선의 주요 역을 거점으로 설정하여, 위 위치중 한 곳이라도 비가 오는 경우를 "비 오는 시간"으로 설정하고, "비 오는 시간"에 운행이 걸쳐있는 기차 운행을 "비 내리는 호남선"으로 정의한다.

사실 이 정의가 썩 맘에 들진 않는다. 목포에 도착 할 때가 다 돼서야 용산에 비가 내릴 수도 있고, 서대전에 비가 내리고 있지만 서대전역에 가기 전에 비가 그칠 수도 있다. 그러나 중간 정차역과 비 오는 시간대를 모두 고려하려면 API 10여번 가까이 해야 하기 때문에, 일단은 충분히 단순화한 휴리스틱으로 계산 알고리즘을 마무리 하겠다. 

---

이 프로젝트를 시작할 때 까지만 해도 노래 가사가

> 비 내리는 호남선, 완행 열차에~

인줄 알았다. 그래서 당연히 고속선 주행하는 KTX는 완행열차가 아니니까 빼야겠구나..라고 생각하고 있었는데, 오히려 남행열차라서 상행선을 계산하지 않아도 되지 완전히 럭키귀찮은개발자잔하~ 빨리 프론트엔드 개발하고 싶다.