---
title: HTML Input의 change 이벤트는 input이 blur될 때 발생한다
layout: ../layouts/article.astro
date: 2024-06-07T10:41:05.224Z
---

# [[HTML]] Input의 change 이벤트는 input이 blur될 때 발생한다

제목 그대로이다. input의 onChange가 값이 Change될 때 일어난다고 착각하고 있었다. 왜냐하면, [[리액트]]가 그렇게 구현되었기 때문이다.

다음 내용이 2024년 06월 기준 최신 문서인 [react.dev](https://react.dev/reference/react-dom/components/input#props)에 나와있다.

> **onChange**   
> [\<input> - React(react.dev)](https://react.dev/reference/react-dom/components/input#props)
>   
> (중략) ... Fires immediately when the input’s value is changed by the user (for example, it fires on every keystroke) ... (후략)   
>   
> 사용자가 `input`의 값을 바꿀 때 즉시 발생합니다 (예를들어, 키를 누를 때 마다 실행됩니다)

이러한 선택을 한 이유는 legacy 문서에 남아있다.

> **onChange**   
> [DOM Elements – React](https://legacy.reactjs.org/docs/dom-elements.html#onchange)
>   
> The onChange event behaves as you would expect it to: whenever a form field is changed, this event is fired. We intentionally do not use the existing browser behavior because onChange is a misnomer for its behavior and React relies on this event to handle user input in real time.   
>   
> 예상한 대로, onChange 이벤트는 입력이 변경될 때마다 동작합니다. React는 브라우저가 제공하는 기존 이벤트를 의도적으로 사용하지 않았습니다. 왜냐하면 onChange는 실제 이름과 다르게 동작하지만, React는 사용자 입력을 실시간으로 처리하기 위해 이 이벤트에 의존하기 때문입니다.`

그럼에도 불구하고, 리액트의 onChange 이벤트는 [[브라우저]]의 onInput 이벤트와 동일하다. 왜 리액트 팀이 onInput 이벤트를 두고도 onChange 이벤트를 직접 재구현했는지는 알 수 없다.
