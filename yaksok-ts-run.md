---
title: 코드 실행하기
date: 2023-12-12T02:15:12.046Z
---
# 코드 실행하기

각 노드가 해야 할 일을 정의해주었으니, 이젠 최상위의 노드들을 실행시켜주기만 하면 됩니다. 

```javascript
import { LineBreak } from './token.js'

export function runAST(ast) {
	// AST를 실행하는 함수입니다

	// 변수를 저장하기 위한 스코프를 선언했습니다
	let scope = new Map()

	// AST의 요소들을 실행하겠습니다
	for (let node of ast) {
		// LineBreak는 무시합니다
		if (node instanceof LineBreak) {
			continue
		}

		node.execute(scope)
	}
}
```

이제 실행에 필요한 모든 준비를 마쳤습니다!
