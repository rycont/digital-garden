---
title: Project IDX
layout: ../layouts/article.astro
date: 2024-05-25T11:18:27.377Z
---

# Project IDX

[[구글]]의 웹 기반 IDE 서비스이다. Code OSS([[VSCode]])를 프론트엔드로 한다. 오래 전부터 Waitlist를 걸어두었는데, [[2024년]] 5월 15일부로 Waitlist 없이 공개되었다.

첫인상: [[Codespaces]]보다 반응이 빠르다, 날렵하다, [[Gemini]] 자동완성은 느리고 멍청하다

원격 개발환경으로 Codespaces를 사용하다가 컴퓨팅 성능 부족과 오류 때문에 VSCode Tunnel([[Self Hosted]])로 이주했다. 만약 써보는 동안 오류와 성능 이슈가 드러나지 않는다면 메인 개발환경으로 사용해볼 것 같다. 며칠간 써봐야겠다.

[[안드로이드]] 에뮬레이터를 제공한다고 해서 추후 앱개발 용도로 테스트해보고자 한다. 특히 Flutter 개발에 최적화된 것 같다.

## 성능


[다른 블로그 글(oracle VM.Standard.A1.Flex CPU 성능 테스트 – AWS 인스턴스와의 비교 포함)](https://blog.layer1.dev/post/it/oracle-vm-standard-a1-flex-cpu-%EC%84%B1%EB%8A%A5-%ED%85%8C%EC%8A%A4%ED%8A%B8-aws-%EC%9D%B8%EC%8A%A4%ED%84%B4%EC%8A%A4%EC%99%80%EC%9D%98-%EB%B9%84%EA%B5%90-%ED%8F%AC%ED%95%A8/)을 참고하여, 다음 명령어로 테스트하였다.

```bash
cat /proc/cpuinfo | grep CPU | wc -l
sysbench cpu --events=10000 --cpu-max-prime=20000 --time=0 --threads=2 run
```

| 환경        | 스레드 | 작업 소요 시간 | 램 크기 | 작업공간(Home dir), /tmp에 할당된 디스크 | 확장 레지스트리             | AI 지원        |
|-------------|-------|---------------|---------|----------------------------------------|---------------------------|----------------|
| Codespaces  | 2     | 6.548s (3.2x) | 8GB     | 32GB, 44GB                             | Visual Studio Marketplace | GitHub Copilot |
| Project IDX | 2     | 21.0521s      | 8GB     | 10GB, 4GB                              | Open VSX                  | Gemini         |

컴퓨팅 환경은 Codespaces가 IDX보다 월등하다. 단위작업 소요 시간은 3배 빠르며, 작업 공간은 3배 이상 크다.

그럼에도 불구하고 체감 반응속도 등에 의해 IDX가 더 쾌적하게 느껴진다. 
