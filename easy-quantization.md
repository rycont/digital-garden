---
title: 딥러닝 모델 쉽게 양자화하기
layout: ../layouts/article.astro
date: 2023-06-09T00:22:08.910Z
---

# 딥러닝 모델 쉽게 양자화하기

> 함께 보기
>
> - [1. 한국어 GPT 모델 직접 증류해보기](https://tilnote.io/pages/6480a6c7e92fe5ef635f4cc9)
>
> - [\[Intro\] 딥러닝 경량화 모델 방법론 소개](https://computistics.tistory.com/22)

[[양자화]]는 딥러닝 모델을 경량화하는 방법중 하나입니다. [[Huggingface]]의 [Transformers](transformer) 라이브러리가 제공하는 툴킷을 사용하면 Huggingface에 업로드된 [[LLM:모델]]을 간편하게 양자화할 수 있습니다. 이번 글에서는 [카카오브레인의 KoGPT 6B (kakaobrain/kogpt)](https://huggingface.co/kakaobrain/kogpt) 모델을 양자화하고 성능 변화를 알아보겠습니다.

### 환경 설정

transformers, accelerate, bitsandbytes 라이브러리가 필요합니다. 특히 Transformers는 4.29 이후의 버전이 필요하기 때문에, 업데이트를 권장합니다.

```prompt
pip install -q -U bitsandbytes transformers accelerate
```

### 양자화 하기

```python
from transformers import AutoModelForCausalLM

model = AutoModelForCausalLM.from_pretrained(
  'kakaobrain/kogpt',
  revision = 'KoGPT6B-ryan1.5b-float16',
  load_in_8bit = True,
  device_map = 'auto'
)
```

`from_pretrained`**<u>로 모델을 불러올 때</u>** `load_in_8bit` **<u>옵션을 함께 주면 모델을 불러올 때 int8로 양자화할 수 있습니다.</u>**

### 모델 배포하기

```python
tokenizer = AutoTokenizer.from_pretrained(
  'kakaobrain/kogpt',
  revision = 'KoGPT6B-ryan1.5b-float16'
)
```

[[토크나이저]]도 불러오겠습니다. 토크나이저를 양자화하진 않지만, 사용자 편의를 위해 모델과 함께 배포하고자 합니다.

양자화한 모델과 토크나이저를 함께 huggingface에 업로드 해보겠습니다.

```javascript
!pip install huggingface_hub
```

```javascript
from huggingface_hub import notebook_login
notebook_login()
```

Notebook 환경에서는 위 코드를 통해 Huggingface에 로그인할 수 있습니다. 로그인할 때 인증 토큰을 요구하는데, 다음의 URL에서 얻을 수 있습니다.

<https://huggingface.co/settings/tokens>

```javascript
model.push_to_hub("rycont/kakaobrain__kogpt-6b-8bit");
tokenizer.push_to_hub("rycont/kakaobrain__kogpt-6b-8bit");
```

원하는 저장소 이름을 지어서 Huggingface Hub에 업로드하였습니다. fp16으로 배포된 기존 모델의 크기는 12.3GB였지만, int8로 양자화한 모델은 6.7GB로 압축되었습니다. 제가 양자화한 KoGPT Int8 양자화 모델은 다음 링크에 배포되었습니다.

[https://huggingface.co/rycont/kakaobrain\_\_kogpt-6b-8bit](https://huggingface.co/rycont/kakaobrain__kogpt-6b-8bit/tree/main)

###
