---
title: 실험적인 "한국어" 기반 프로그래밍 언어의 제안
date: 2023-12-10T10:13:33.244Z
---

# 실험적인 "한국어" 기반 프로그래밍 언어의 제안

[[한국어]]의 구조를 차용한 프로그래밍 언어?

한글 [[프로그래밍]]을 넘어서, 한국어 프로그래밍 언어를 제안한다.

```wast
("안녕," + 이름 + "!")는 (이름)에게 인사하기
"지민"에게 인사한 결과를 보여주기

"재현", "지민", "성수"는 (학급에 있는 학생들)이다
각 (학급에 있는 학생들)에게 인사한 결과를 (인삿말)에 저장하기

만약 (학급에 있는 학생들)의 수가 3보다 많으면
  "지금 사람이 너무 많아요"를 보여주기
아니면
  "조금 더 들어가도 될 것 같아요"를 보여주기

(배열)을 (구분자)로 합치기는 (
  (합친 결과)는 ""
  (숫자)는 1

  (
    (합친 결과) 뒤에 배열의 (숫자)번째 요소을 붙히기
    (숫자)에 1을 더하기
  )
  만약 배열의 길이가 숫자보다 크다면 계속 반복하기

  (합친 결과)를 결과로 사용하기
)

(인삿말)을 (",")로 합친 결과를 보여주기
```

```wast
(새 Express)는 (백엔드)이다.
(새 Kysely)는 (디비)이다.
백엔드에 "/login"이라는 요청이 "GET"으로 들어올 때 (
  (탐색 결과)는 디비의 "users" 테이블에서 (
    "이름"이 (요청)의 "이름"
  )인 데이터를 찾은 결과이다

  (요청된 비밀번호의 해시)는 (요청)의 "비밀번호"를 MD5로 해싱한 결과이다
  만약 (탐색 결과)의 (비밀번호 해시)가 (요청된 비밀번호의 해시)와 같다면 (
    결과는 탐색 결과에서 "1234"를 키로 하는 JWT를 생성한 결과이다
  ) 아니라면 (
    결과는 "비밀번호가 일치하지 않습니다"이다
  )
)
```
