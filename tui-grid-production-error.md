---
title: Tui Grid가 프로덕션 빌드에서만 오류를 낼 때 (Minified React error #130)
date: 2023-06-16T13:27:05.984Z
---

# Tui Grid가 프로덕션 빌드에서만 오류를 낼 때 (Minified React error #130)

![](../images/5c3157ec-4efd-4af5-8667-11c7be14e69a.png)

[[TUI Grid]]의 Default import엔 Grid 리액트 컴포넌트가 담겨있는데, 왠지 프로덕션에서는 `{ default: Grid }`의 형태로 임포트가 된다. 그래서 Object를 요소로 렌더링 할 수 없다는 오류가 발생한다.

다음과 같이 임포트 해주면 정상적으로 사용할 수 있다.

```javascript
import _Grid from '@toast-ui/react-grid'
let Grid = _Grid

if ('default' in Grid) {
    Grid = Grid['default'] as typeof _Grid
}
```

.. I don't like it 😶
