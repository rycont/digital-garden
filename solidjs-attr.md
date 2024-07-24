---
title: SolidJS에서 WebComponent에 동적으로 Attributes 넘기기
date: 2024-05-20T07:13:05.921Z
---

# SolidJS에서 WebComponent에 동적으로 Attributes 넘기기

처음부터 [[SolidJS]]는 [[WebComponent]] 지원을 염두에 두고 설계되었다. 그렇기에 WebComponent와 호환성도 좋고 섞어쓰기 편리하다.

## 문제 상황

SolidJS에서 WebComponent를 호출해서 사용할 때, Attributes에 Computed Value는 전달되지 않고 Literal로 작성한 값만 전달된다. 예를 들어,

```
<some-web-component
    name="Hojin"
    birth={+new Date()}
/>
``` 

라고 작성한다면 attributes엔 name: Hojin만 전달된다.

## 해결 방안

birth 값이 Property로 취급되었기 때문이다. WebComponent는 Attributes만 읽을 수 있기에, 강제로 attributes로 전달하는 디렉티브를 사용해야 한다.

```
<some-web-component
    attr:name="Hojin"
    attr:birth={+new Date()}
/>
```
