---
title: 프레임워크 없이 만든 프론트엔드에 작은 부분부터 Svelte 도입하기
layout: ../layouts/article.astro
date: 2024-04-12T05:49:13.793Z
---

# 프레임워크 없이 만든 프론트엔드에 작은 부분부터 Svelte 도입하기

## 들어가며

[[프론트엔드]] 개발을 [[프레임워크]] 없이 구현하다 보면, 복잡한 [[UI]]와 데이터 흐름을 관리하기 어려워지는 경우가 많습니다. 이럴 때 프레임워크를 도입하는 것이 좋지만, 기존 프로젝트 전체를 재작성하기에는 큰 부담이 될 수 있습니다.

이번 글에서는 프레임워크 없이 작성된 프론트엔드에 Svelte를 점진적으로 도입하는 방법에 대해 다루어 보겠습니다. [[Svelte]]는 [[WebComponent]] 호환을 지원하므로, 기존 프로젝트 구조에 큰 변화 없이 컴포넌트를 점진적으로 적용할 수 있습니다.

## 사전 준비

다음과 같은 파일의 버튼을 Svelte 컴포넌트로 분리해보겠습니다.

```html
<!DOCTYPE html>
<html lang="ko">
  <head>
    <meta charset="UTF-8" />
    <style>
      button {
        padding: 0.5rem;
        background: #fefefe;
        border: 1px solid #d0d0d0;
        border-radius: 0.25rem;
      }
    </style>
  </head>
  <body>
    <button class="say-name" data-name="재우">이름 말하기</button>
    <button class="say-name" data-name="수열">이름 말하기</button>
    <script>
      const sayNames = document.getElementsByClassName("say-name");
      for (const button of sayNames) {
        button.addEventListener("click", () => {
          alert(button.getAttribute("data-name"));
        });
      }
    </script>
  </body>
</html>
```

### Vite 설치하기

[[Vite]]는 프론트엔드 개발 및 빌드 도구로, Svelte 컴포넌트를 쉽게 통합할 수 있습니다. Vite는 빠른 빌드 속도와 다양한 플러그인을 제공합니다.

필요한 라이브러리를 설치합니다:

```javascript
pnpm add -D vite @sveltejs/vite-plugin-svelte svelte glob
```

#### 설정

그 다음 package.json에 다음과 같은 스크립트를 추가합니다

```javascript
"dev": "vite",
"build": "tsc && vite build",
```

package.json의 type 속성을 module로 설정합니다

```javascript
"type": "module"
```

vite.config.ts 파일을 만들고 다음과 같이 설정합니다

```javascript
import { svelte, vitePreprocess } from "@sveltejs/vite-plugin-svelte";
import { defineConfig } from "vite";
import { resolve } from "path";
import { glob } from "glob";

const pages = await glob("**/*.html", {
  ignore: ["node_modules/**", "dist/**"],
});

const entryPoint = Object.fromEntries(
  pages.map((path) => [path, resolve(__dirname, path)])
);

export default defineConfig({
  appType: "mpa",
  build: {
    target: "ESNext",
    rollupOptions: {
      input: entryPoint,
    },
  },
  plugins: [
    svelte({
      preprocess: vitePreprocess(),
      compilerOptions: {
        customElement: true,
      },
    }),
  ],
});
```

이제 Vite와 Svelte 설정이 완료되었습니다!

## Svelte 컴포넌트 만들기

기존 프로젝트에서 버튼 관련 코드를 Svelte 컴포넌트로 만들어 보겠습니다.

```javascript
<svelte:options customElement="say-name-button" />

<script>
  export let name
</script>

<button on:click={() => alert(name)}>
  이름 말하기
</button>

<style>
  button {
    padding: 0.5rem;
    background: #fefefe;
    border: 1px solid #d0d0d0;
    border-radius: 0.25rem;
  }
</style>
```

이 Svelte 컴포넌트는 `<svelte:options>` 태그를 통해 \`say-name-button\`이라는 이름의 Custom Element로 정의되며, `name` props를 받아 버튼 클릭 시 해당 이름을 alert로 출력합니다.

## HTML에서 불러오기

이제 Svelte 컴포넌트를 기존 HTML 파일에서 사용할 수 있습니다!

```html
<script type="module" src="./say-name-button.svelte"></script>
<say-name-button name="소희"></say-name-button>
```

이렇게 Svelte 컴포넌트를 점진적으로 도입하면, 기존 프로젝트 구조를 크게 변경하지 않으면서도 프레임워크의 장점을 활용할 수 있습니다.
