---
title: DaramGPT 학습 준비
layout: ../layouts/article.astro
date: 2023-06-07T16:07:12.283Z
---

# DaramGPT 학습 준비

앞에서 알아본 내용을 통해서, [](Knowledge Distillation)을 거치면 모델의 성능은 거의 유지하면서 크기를 줄일 수 있다는 점을 알았습니다. 이전까지의 연구는 1GB 이하의 모델만을 대상으로 Distillation을 하였습니다. 하지만 최근에 공개되는 대부분의 모델들은 그보다 커지고 있고, 일반적인 성능의 기기에서 미세조정과 Inference를 거치기는 더욱 어려워지고 있습니다.

이번 실험에서는 한국어 대규모 생성모델인 KoGPT Trinity 모델을 Distillation해서 압축해보고자 합니다. KoGPT Trinity는 SKT에서 GPT3의 구조를 모방하여 학습한 모델로, 국내에 공개된 GPT 모델들 중에서는 두번째로 큰 크기를 가지고 있습니다. KoGPT Trinity의 이전버전으로는 KoGPT2가 있는데, 다른 모델들이 비해 필요로 하는 성능이 낮아 국내에서 가장 널리 사용되고 있는 것으로 알고 있습니다. 그래서, KoGPT Trinity를 Distillation해서 KoGPT 2보다 작고 강한 모델을 만들어보자! 를 목표로 실험을 진행하였습니다.

#### **환경 준비**

학습에 사용할 컴퓨팅 플랫폼으로는 [](Kaggle) Notebooks를 선정하였습니다. Kaggle Notebooks에서는 일주일에 40시간 내외로 사용할 수 있는 [](NVIDIA) P100 [](GPU) 하나를 무료로 제공합니다. 컴퓨팅 자원이 충분하지 않기 때문에 여러 어려움이 있을 것 같지만, 그래도 가능한 데 까지는 실험을 진행해보고자 합니다.

##### **Student 모델 구성**

Student 모델은 다음과 같이 구성하였습니다

```javascript
n_embd: 768;
n_head: 8;
n_layer: 4;
```

이는 KoGPT2의 구조에서 일부 파라미터 수를 줄인 것입니다. KoGPT2의 원본 모델 구성은 다음과 같습니다.

```javascript
n_embd: 768;
n_head: 12;
n_layer: 12;
```

Embedding의 차원 수(n_embd)를 Teacher 모델과 같은 1920개로 유지하여 Cosine Embedding Loss도 활용하고 싶었지만, n_embd 수를 늘리면 모델의 크기가 1.8GB정도로 늘어나는 모습을 볼 수 있었습니다. 실험의 목적인 “KoGPT2보다 작은 모델”을 달성하기 위해서, 아쉽게도 Embedding Loss는 제외하고 단어 예측 로스만 사용하기로 하였습니다.

##### **코드 구성**

```javascript
kld = torch.nn.KLDivLoss(reduction='batchmean')

student_output = student(**model_input)
student_log_softmax = (student_output.logits / TEMPERATURE).log_softmax(dim = -1)

teacher_output = teacher(**model_input)
teacher_softmax = (teacher_output.logits / TEMPERATURE).softmax(dim = -1)

distill_loss = kld(
  student_log_softmax,
  teacher_softmax
) * TEMPERATURE ** 2

actual_loss = student_output.loss
summary_loss = distill_loss + 0.1 * actual_loss

summary_loss.backward()
optimizer.step()
student.zero_grad()
scheduler.step()
optimizer.step()
```

학습에 사용된 코드입니다. 이전장에서 설명한 내용을 그대로 코드로 구현하였습니다. 최적화함수와 학습 코드 구현은 Huggingface Research의 Distillation.py 코드를 참고하였습니다. 또한 Softmax Temperature를 2로 하는 것, L1 로스와 L2 로스의 비를 10:1로 하는 것도 참고하였습니다.

[transformers/examples/research_projects/distillation at main · huggingface/transformers (github.com)](https://github.com/huggingface/transformers/tree/main/examples/research_projects/distillation)

##### **데이터셋 준비**

학습에 사용한 말뭉치는 다음과 같습니다

- **대규모 웹데이터 기반 한국어 말뭉치**\
  \- 출처: [](AI-Hub)\
  \- 사용한 양: 4.5GB

- **문어 말뭉치**\
  \- 출처: 모두의 말뭉치\
  \- 사용한 양: 3GB

- **한국어 위키백과 덤프**\
  \- 출처: 한국어 위키백과**\
  \-** 사용한 양: 600MB

- **비출판물 말뭉치**\
  \- 출처: 모두의 말뭉치\
  \- 사용한 양: 20MB

메타데이터 등은 제외한 순수 텍스트의 양만 집계하였습니다. 문어[](말뭉치)는 각 텍스트의 한 문단을 통으로 투입하여 학습하였고, 나머지는 문장별로 분리하여 학습하였습니다. 총 8.3GB가량의 한국어 문어체 텍스트가 입력되었습니다. Epoch은 1로 설정하여 반복학습은 진행하지 않았습니다.

각 데이터를 다음과 같은 순서로 투입하였습니다

1. 대규모 웹데이터 기반 한국어 말뭉치 300MB \* 2

2. 비출판물 말뭉치

3. 대규모 웹데이터 기반 한국어 말뭉치 300MB \* 13개

4. 한국어 위키피디아 덤프 600MB

5. 모두의 말뭉치 문어체 300MB \* 10

학습이 끝날 때 까지 꾸준하게 로스가 줄어들었습니다. 후반부에는 로스가 거의 진동하였고, 더 이상 학습이 진행되지 않는 것으로 파악하여 학습을 중단하였습니다. 약 12일, 총 284시간 가량 학습하였습니다.

284시간을 40시간씩 나눠서 사용했으니 Kaggle Notebooks에서 7주동안 학습해야 했으나, 지인의 계정을 동원하여 약 2주일이 조금 넘게 학습하여 끝낼 수 있었습니다.
