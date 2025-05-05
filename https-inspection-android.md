---
title: PCAPDroid로 안드로이드 장비에서 HTTPS 요청 훔쳐보기
date: 2025-05-05T04:43:13.459Z
---

# PCAPDroid로 [[안드로이드]] 장비에서 HTTPS 요청 훔쳐보기

[[기차와 버스]]를 개발하던 중, 도저히 [[티머니]] 시외버스 웹사이트를 크롤링하긴 엄두가 나지 않아 티머니고 앱에서 시외버스 관련 API를 추출하기로 했다. 패킷 캡쳐 앱은 PCAPDroid를 사용해보자.

1. PCAPDroid를 설치하고, 앱을 티머니고로 설정하니 암호화된 TLS 트래픽이 보였다.
2. TLS 트래픽을 복호화 하려면 별도 설정이 필요하다고 했다. 설정 화면에서 TLS Decryption 옵션을 켜주었고, CA 인증서를 설치해주었다. 다시 캡쳐해보니 아직도 TLS 트래픽은 암호화되어있었다.
3. 앱의 사이드 메뉴에 "Decryption Rules"라는 메뉴에서 규칙을 설정해줘야 했다. 복호화할 앱으로 티머니고를 골랐다.
4. 그랬더니 이젠 티머니고에서 어떠한 네트워크 연결도 작동하지 않았다. `SSL Pinning`이라는 규칙 때문이였다. SSL Pinning은 통신에 사용할 목적지 서버의 공개키를 미리 앱에 넣어두고, 키가 일치하지 않는 요청은 모두 차단하는 방식이다.
5. 물론 이 또한 PCAPDroid의 계획중에 있었다. 매뉴얼에 따르면 [apk-mitm](https://github.com/niklashigi/apk-mitm)이라는 툴을 사용하면 APK에서 SSL Pinning 관련 기능을 삭제할 수 있다고 한다.
    - > **apk-mitm**   
A CLI application that automatically prepares Android APK files for HTTPS inspection   
https://github.com/niklashigi/apk-mitm
6. APK Pure에서 티머니고의 XAPK 파일을 다운로드 받고 SSL Pinning을 제거해주었다. `-patched` 접미사가 붙은 파일 이름이 생성된다.
```bash
> ~ $ apk-mitm tmoney-go-secure.xapk 

╭ apk-mitm v1.3.0
├ apktool v2.9.3
╰ uber-apk-signer v1.3.0

Using temporary directory:
/tmp/apk-mitm-1b9e9afb769ee34dbed6d4084d3bbe9a

✔ Extracting APKs
✔ Finding base APK path
✔ Patching base APK
✔ Signing APKs
✔ Compressing APKs

Done!  Patched file: ./tmoney-go-secure-patched.xapk
```
7. `XAPK` 파일은 안드로이드 순정 APK 설치기로 설치할 수 없어서, SAI라는 앱으로 설치해주었다. 설치중에 무수한 경고창이 떴지만, 중요하지 않다.
8. 앱이 실행되지 않는다. 패키지 무결성 정책을 사용하는것 같다. 이렇게 된 이상 APK를 수정하는 방법으로는 더 진행할 수 없다.
9. 그렇다면 SSL Pinning이 적용되지 않는 Android 버전을 지원하는 APK 파일을 구해서 설치해봐야겠다. SSL Pinning은 Android 7.0 이상에서 적용된다고 하니, 6.0을 지원하는 APK를 설치해서 패킷을 캡쳐해보기로 하자.
    - 실패 가능성은 다음과 같다:
        1. 구버전이라서 실행이 안된다
        2. 실행은 되는데 동일하게 네트워크 연결이 안된다
        3. 네트워크 연결은 되지만 패킷이 안잡힌다
10. 안된다. 실패 사유는 2번이였다.

오늘의 도전은 이렇게 실패다. 다음에 iOS 기기를 구해서 mitmproxy로 진행 해봐야겠다.
