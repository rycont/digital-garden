---
title: Cloudflare Pages에 Flutter 앱 배포하기
layout: ../layouts/article.astro
date: 2023-06-09T00:13:35.808Z
---

# Cloudflare Pages에 Flutter 앱 배포하기

[](Cloudflare) Page는 프론트엔드 프로젝트를 간편하게 배포할 수 있게 도와주는 플랫폼입니다. 빌드 명령어 프리셋으로 [](React), [](Vue) 등 여러 프레임워크를 지원하고 있지만. [](Flutter) Web은 지원되지 않습니다. 다음 Build Command를 통해서 빌드머신에 직접 플러터를 설치하고 실행할 수 있습니다. Build Output Directory는 /build/web로 설정하면 됩니다.

```javascript
wget "https://storage.googleapis.com/flutter_infra_release/releases/stable/linux/flutter_linux_3.0.5-stable.tar.xz" && tar -xf ./flutter_linux_3.0.5-stable.tar.xz && export PATH="$PATH:`pwd`/flutter/bin" && flutter build web --release
```

2022년 8월 2일 기준 최신버전인 3.0.5를 설치하고 빌드를 실행하는 명령어입니다. 이러한 단계들을 한 번에 실행하는 명령어입니다.

1. 구글 공식 릴리즈 서버에서 Flutter 3.0.5를 버전을 다운로드합니다

2. tar 압축을 해제합니다

3. 압축 해제한 플러터 폴더를 PATH 환경변수에 등록합니다

4. 웹 빌드 명령어를 실행합니다

![](../images/8c19ce84-f941-4518-8017-b496ca30618a.png)
