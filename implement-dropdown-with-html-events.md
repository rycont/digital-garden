---
date: 2024-08-16T13:25:07.348Z
title: "HTML에서 Dropdown을 직접 구현할 때 Focus / Blur Events를 활용하기"
---

# [[HTML]]에서 Dropdown을 직접 구현할 때 Focus / Blur Events를 활용하기

열려 있는 [[Dropdown]]의 밖을 클릭할 때 닫히도록 하는 것은 흔한 요구 사항이다. 그러나 이를 직접 [[구현]]하는 것은 성가신 일이다.

## 직접 구현하기

```tsx
const EVENT_RUN_ONCE = { once: true }

const [selectorOpen, setSelectorOpen] = createSignal(false)

function openSelector() {
    setSelectorOpen(true)

    document.addEventListener(
        'click',
        onDocumentClick,
        EVENT_RUN_ONCE
    )
}

function onDocumentClick(e: MouseEvent) {
    const target = e.target as HTMLElement
    const isClickedInside = target.closest(`.${wrapperStyle}`)

    if (isClickedInside) {
        return
    }

    setSelectorOpen(false)
}

return (
    <div
        class={wrapperStyle}
        onClick={openSelector}
    >
        {/* 현재 선택된 아이템 보여주기 */}
        <Show when={selectorOpen()}>
            {/* 아이템 선택창 보여주기 */}
        </Show>
    </div>
)
```

Dropdown을 열 때 마다 [[document]] click 이벤트를 감지하도록 한다. 이 방법도 좋지만, Focus 이벤트를 사용해서 코드를 더 간단하게 만들어보자.

## Blur Event를 활용하기

```typescript
let ref!: HTMLDivElement;

const [ selectorOpen, setSelectorOpen ] = createSignal(false);

function openSelector() {
    setSelectorOpen(true);
}

function closeSelector() {
    setSelectorOpen(false);
}

return (
    <div
        class={wrapperStyle}
        onBlur={closeSelector}
        onClick={openSelector}
        tabIndex={0}
        ref={ref}
    >
        {/* 현재 선택된 아이템 보여주기 */}
        <Show when={selectorOpen()}>
            {/* 아이템 선택창 보여주기 */}
        </Show>
    </div>
);
```

div는 일반적으로 포커스를 가질 수 없지만, `tabindex` 속성을 준다면 포커스를 가질 수 있게(Focusable) 만들 수 있다. focus 이벤트는 다음의 경우에 발생한다: 

- [[키보드]] [[Tab]] 키로 엘리먼트 포커스를 줄 때
- [[마우스]]로 클릭할 때

또한 Focusable한 엘리먼트는 다음의 경우에 포커스 상태에서 벗어나며, 동시에 Blur 이벤트를 만든다

- Tab(Shift-Tab)을 통해 엘리먼트가 포커스 상태에서 벗어남
- 엘리먼트의 밖을 클릭할 때

focus와 blur 이벤트를 활용하면 document click event를 사용하지 않고도 간단하게 Dropdown을 구현할 수 있다. HTML에서 제공하는 기본 이벤트를 활용함으로써, 직접 구현하지 않고도 복잡한 상호작용을 만들었다.
