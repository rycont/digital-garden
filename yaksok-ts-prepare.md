---
title: 준비하기
date: 2023-12-12T02:09:45.811Z
---

# 준비하기

약속 [[프로그래밍 언어]]는 한국어를 기반으로 하는 프로그래밍 언어입니다. 다음과 같이 작성합니다.

```javascript
키: 170
몸무게: 60
비만도: 키 / 몸무게 / 몸무게

"비만도는 " + 비만도 보여주기
```

이번 튜토리얼에서는 약속의 전체 명세를 구현하지 않고, 위의 예시코드를 실행할 수 있는 [[인터프리터]]를 만드는 것을 목표로 합니다.

---

일반적인 인터프리터는 네가지 단계를 거쳐서 코드를 실행합니다.

1. 토크나이저(Tokenizer): 코드를 각 토큰(어절, 단어)로 분할합니다

2. 렉서(Lexer): 토큰에 의미를 부여합니다

3. 파서(Parser): 토큰과 토큰 사이의 관계를 분석하고, 의미를 포함하여 구조적으로 재조합 합니다. 재조합된 결과를 AST(Abstract Syntax Tree)라고 합니다.

4. 런타임(Runtime): AST를 실행합니다

여기서 준비단계에 해당되는 토크나이저, 렉서, 파서를 먼저 알아보겠습니다.

## 토크나이저

> 코드를 각 토큰(어절, 단어)로 분할합니다

```javascript
const name = "수연";
console.log(name);
```

라는 코드를 Tokenize하면 다음과 같습니다.

```javascript
[
  "const",
  "name",
  "=",
  '"',
  "수연",
  '"',
  "\n",
  "console",
  ".",
  "log",
  "(",
  "name",
  ")",
];
```

## 렉서

> 토큰에 의미를 부여합니다

위 토큰들을 다음과 같이 변환합니다

```javascript
[{
    "type": "Keyword",
    "value": "const"
}, {
    "type": "Variable",
    "value": "name"
}, {
    "type": "Operator",
    "value": "="
}, {
    "type": "StringLiteral",
    "value": "수연"
}, {
    "type": "LineBreak",
    "value": "\n"
},
    ...
]
```

대부분의 언어 런타임은 이처럼 ([[토크나이저]] | [[렉서]] | [[파서]])의 3단계로 코드를 분석하지만, 이번 튜토리얼에서는 토크나이저와 렉서를 한번에 구현하겠습니다.

다음 글: [[yaksok-ts-lexer|토크나이저와 렉서 만들어보기]]
