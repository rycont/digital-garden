---
title: 토이프로젝트를 위한 서비스 추천
layout: ../layouts/article.astro
date: 2023-09-01T04:24:37.235Z
---

# 토이프로젝트를 위한 서비스 추천

## 프론트엔드 배포

[GitHub](깃허브) Pages는 이제 그만 놓아주어요. [[Vercel]], [[Netlify]], [[Cloudflare]] Pages를 사용한다면 깃허브 저장소에 있는 프론트엔드 코드를 쉽게 배포할 수 있습니다.

- [[React]], NextJS, SvelteKit과 같은 빌드툴 / SSR 프레임워크도 사용 가능

- 별다른 설정 없이 클릭 몇번 만에 바로 배포 가능

- (원한다면) 백엔드 코드도 함께 배포 가능

- 대부분 토이프로젝트 수준에서는 무료

- HTTPS 지원

- **중요: PR을 만들면 변경사항을 프리뷰할 수 있도록 추가로 배포해줌**

  - ![](../images/478f29cc-2539-4058-adc3-8d541c2b421d.png)

  - 그래서 PR을 리뷰할 때 로컬에 클론할 필요 없이 브라우저에서 바로 확인 가능

  - 협업 생산성이 대단하게 증가할 수 있음

각 서비스별 차이점은 다음과 같습니다

- **Vercel**

  - 장점

    - ==교육용 / 오픈소스 지원용 무료 플랜==이 있음. 버셀팀과 컨택 필요.

      - 받기 어려움요

    - 디자인이 예쁨

    - ==NextJS와의 궁합이 좋음==

    - 한국리전에 배포 가능

  - 단점

    - ==Organization에 있는 레포를 배포하려면 팀플랜을 결제해야 함==

- **Netlify**

  - 장점

    - ==Organization에 있는 레포도 개인계정에서 배포 가능==

    - Auth, Forms, Database 등의 플러그인 생태계가 잘 되어있음 -> 간단한 게시판 정도의 프로젝트는 Netlify만으로 개발 가능

  - 단점

    - Vercel에 비해 빌드 속도가 조금 느리다

- **Cloudflare Pages**

  - 장점

    - ==클라우드플레어에 등록한 도메인을 바로 연결할 수 있다==

    - Cloudflare Registrar에서 도메인을 구입하고, Cloudflare DNS에 등록하고, Cloudflare Pages에 연결하는 과정이 다른 어떤 서비스들과 비교할 수 없을 정도로 간편함

    - DNS 설정하고 도메인 연결하는게 은근 스트레스인거 알잖아요 ..

    - ==Cloudflare 망에 배포되기 때문에 로딩속도가 무지 빠름==

    - ==팀계정도 무료==

  - 단점

    - 빌드 속도가 약간 느림

    - 백엔드 코드 실행이 NodeJS 런타임이 아님.

      - Cloudflare Workers (Edge Runtime)라는 독자 런타임을 사용함

      - NextJS 등의 SSR 프레임워크에서 호환성 문제가 발생할 수 있음

## 데이터

### **Supabase: PostgreSQL와 File Storage 뽑아먹기**

[[Supabase]]는 데이터베이스에 보안정책을 직접 설정하여 [[프론트엔드]]에서 DB에 안전하게 직접 접근할 수 있도록 하는 서비스임. 그러나 이런 기능을 사용하지 않더라도, Supabase에 가입하면 PostgreSQL 인스턴스 하나를 주기 때문에 일반적인 백엔드용으로 사용해도 좋습니다. File Storage도 주니, 적극적으로 뽑아먹읍시다!

### PlanetScale: MySQL 뽑아먹기

[[PlanetScale]]은 MySQL 데이터베이스의 스키마를 안전하게 구성 / 변경할 수 있도록 DB에 브랜치 형식을 도입한 DBaaS 서비스입니다. 가입하면 [[MySQL]] 인스턴스를 하나 주니 뽑아먹으면 됩니다.

### Upstash: Redis, Kafka 뽑아먹기

[[Upstash]]는 데이터플랫폼 전문 클라우드 서비스입니다. [[Redis]]와 [[Kafka]] 프리티어를 제공합니다.

### Browserless: Headless Chromium 뽑아먹기

그럴 일이 얼마나 있을지는 모르겠지만 Browserless라는 솔루션을 이용하면 리모트서버에 있는 headless chromium [[브라우저]]를 제어할 수 있습니다.

## 백엔드

### [[Cloudtype]]

한국형 클라우드 서비스입니다. 사용 경험이 훌륭하고 프리티어가 넉넉합니다. 깃허브 레포지토리를 연결하고 빌드, 포트 등을 설정하면 자동으로 배포할 수 있습니다. 서버 접근 속도가 다른 클라우드와는 비교할 수 없을 정도로 빠릅니다. [[Docker|Dockerfile]]만 있으면 백엔드를 바로 배포할 수 있습니다.

### [[Railway]]

흔히 "백엔드계의 Vercel"이라고 불릴정도로 미려하고 간편한 UI를 자랑하는 클라우드 서비스 입니다. Railway에서도 Dockerfile만 있으면 배포할 수 있습니다. PostgreSQL, Redis와 같은 데이터 플랫폼들도 호스팅하기 때문에, 한 [[클라우드]] 안에서 해결할 수 있습니다.

Cloudtype과 동일하게 깃허브 레포지토리를 연결하고 빌드, 포트 등을 설정하면 자동으로 배포할 수 있습니다.

### Oracle Cloud (OCI)

다른 서비스들과는 달리 전통적인 클라우드 서비스입니다. [[Oracle]] Cloud의 VM Instance 서비스에서 A1.Flex라는 인스턴스는 성능이 훌륭하면서도 무료인걸로 유명합니다. 24GB의 램과 4코어의 [[ARM]] CPU를 제공합니다. 제가 애용합니다.

### 서버 더 쉽게 관리하기

도커 컨테이너들을 쉽게 관리할 수 있는 서비스들이 있다면 화내지 않고 서비스를 운영하는데 도움이 됩니다.

- [[CapRover]]: 도커에 익숙하지 않은 사람도 쉽게 이용할 수 있습니다

- [[Portainer]]: 저는 포테이토라고 부릅니다. CapRover보다 더 전문적으로 관리할 수 있지만, UI가 복잡하게 느껴질 수 있습니다

- NPM(Nginx Package Manager): [[Nginx]] 설정을 GUI로 할 수 있게 도와줍니다

## Self Hosted Services

원한다면 상용 서비스들 대신 Self Hosted 서비스를 직접 운영할 수 있습니다.

- UseMemo: 구글킵의 대체

- Mastodon / Misskey: 트위터의 대체

- Mattermost: Slack의 대체

- Ghost: 훌륭한 블로깅 솔루션

- Supabase: Firebase의 대체제. 얘도 호스팅이 된답니다.

- Minio: S3의 대체제

- CodeServer: Codespaces의 대체제

- GitLab: GitHub의 대체제

- Outline VPN: 훌륭한 VPN

- WireGuard: 훌륭한 VPN

- Outline Wiki: Notion의 대체제

- Umami: 구글 애널리틱스의 대체제

- ERPNext: 이카운트 ERP의 대체제. 좋음.

- Discourse: 스레드 형식의 포럼을 호스팅할 수 있음. 좋음. 강추.

- Jina: 검색 서비스를 만들 때 유용한 AI 검색엔진 솔루션. 검색 뿐만 아니라 AI 서비스에 종합적으로 연동 가능
