---
title: 양자화한 모델의 성능을 평가하기
layout: ../layouts/article.astro
date: 2023-11-27T13:52:32.941Z
---

# [](양자화)한 모델의 성능을 평가하기

[](한국어) 능력 평가 지표중 하나인 [KOBEST(Korean Balanced Evaluation of Significant Tasks)](https://arxiv.org/abs/2204.04541)를 이용하여 모델의 성능을 비교해보겠습니다.

간편한 측정을 위해 [EleutherAI의 lm-evaluation-harness](https://github.com/EleutherAI/lm-evaluation-harness)라는 프로그램을 사용하겠습니다. 원본 코드는 양자화한 모델을 아직 완벽하게 지원하지 않습니다. 그래서 원본 코드를 조금 수정하여 테스트를 진행하였습니다. 벤치마크에 영향을 미칠만한 부분(프롬프트, 하이퍼파라미터) 등은 일체 수정하지 않았습니다.

수정한 내용은 다음 커밋에서 볼 수 있습니다.

[fix: keep auto device when 8bit model is loaded · rycont/lm-evaluation-harness@5b0e328](https://github.com/rycont/lm-evaluation-harness/commit/5b0e328218643ba0dbe453053c3379d5726a8205)

**중요: 한국어 평가 데이터셋(KoBEST)는 [](Polyglot)이라는 브랜치에서만 지원합니다. 양자화 지원 코드 수정 또한 해당 브랜치에서 진행했기 때문에, 제 코드를 사용하고자 하신다면 브랜치를 이동해주세요.**

다음과 같은 커맨드로 평가를 수행하였습니다.

```javascript
python main.py \
   --model "hf-causal-experimental" \
   --model_args pretrained='rycont/kakaobrain__kogpt-6b-8bit',load_in_8bit=True \
   --tasks kobest_boolq \
   --num_fewshot=5 \
   --batch_size=3 \
   --device="cuda" \
   --output_path "./output"
```

저는 [](Colab) 환경에서 진행했기 때문에 긴 시간을 IDLE 상태로 둘 수 없어 각 태스크를 하나 하나 실행하였습니다. `--tasks kobest_boolq,kobest_nsmc`와 같이 작성하며 여러 태스크를 한 명령어로 수행할 수 있습니다.

KoGPT 6B의 평가 자료는 다음 링크에서 인용하였습니다. 비교는 5-shot을 기준으로 하였습니다.

[EleutherAI/polyglot-ko-5.8b · Hugging Face](https://huggingface.co/EleutherAI/polyglot-ko-5.8b)

| 5-shot F1 Score | KoGPT fp16 | KoGPT int8 (Ours)           |
| --------------- | ---------- | --------------------------- |
| CoPA            | 0.7287     | 0.7277 (↓0.01%)             |
| HellaSwag       | 0.5833     | **<u>0.4560 (↓21.82%)</u>** |
| BoolQ           | 0.5981     | 0.6015 (↑0.56%)             |
| WiC             | 0.4775     | **<u>0.3706 (↓22.38%)</u>** |

CoPA와 BoolQ는 원본 모델에 비해 오차 범위 내의 결과가 나왔지만, **<u>HellaSwag와 WiC는 성능이 크게 다르게 나왔습니다.</u>** CoPA와 BoolQ는 점수를 거의 유지한 것을 보아 나머지 태스크의 진행 과정에 문제가 있는 것으로 보아, 동일한 코드로 다른 모델을 평가해보았습니다.

대조에 사용한 모델은 skt/kogpt-trinity-1.2b-v0.5 모델입니다.

| 5-shot F1 Score | In report | Self Test                   |
| --------------- | --------- | --------------------------- |
| CoPA            | 0.6477    | 0.6476 (↓0.01%)             |
| HellaSwag       | 0.5272    | **<u>0.3999 (↓24.14%)</u>** |
| BoolQ           | 0.4014    | 0.4013 (↓0.02%)             |
| WiC             | 0.4313    | **<u>0.3953 (↓8.34%)</u>**  |

동일하게 WiC와 HellaSwag에서 큰 차이가 있었습니다. 이러한 현상의 원인은 알 수 없으나, 올바르게 측정되었다고 추정할 수 있는 CoPA와 BoolQ의 점수만을 인용한다면 눈에 띄는 성능의 저하 없이 양자화가 완료되었다고 볼 수 있습니다.

---

Polyglot 1.3B의 Discussion 글에서 KoBEST의 HelloSwag 태스크에서 문제를 겪고 있는 사람을 발견하였습니다.

> As I tried to reproduce the polyglot evaluation, kobest_hellaswag does not match, although other datasets(kobest_copa, wic, boolq) match well.\
> Is there a problem or the kobest_hellaswage data has been changed?
>
> (번역) Polyglot의 평과 결과를 재현해보려고 했지만, 다른 데이터셋(kobest_copq, wic, boolq)는 잘 일치했지만 kobest_hellaswag의 결과는 그렇지 않았습니다. 문제가 있거나 kobest_hellaswag의 데이터가 바뀌었나요?

[EleutherAI/polyglot-ko-1.3b · I can't reproduce the kobest_hellaswag](https://huggingface.co/EleutherAI/polyglot-ko-1.3b/discussions/2)

추후 EleutherAI에 KoBEST 벤치마크 점수를 문의하였습니다: [EleutherAI의 KoBEST 벤치마크에 대해](eleutherai-kobest-error)
