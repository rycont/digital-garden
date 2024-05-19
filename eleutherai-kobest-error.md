---
title: EleutherAI의 KoBEST 벤치마크에 대해
layout: ../layouts/article.astro
date: 2023-11-30T11:21:55.829Z
---

# (번외) EleutherAI의 [[KoBEST]] 벤치마크에 대해

[Cannot reproduce the evaluation score of HellaSwag, WiC · Issue #37 · EleutherAI/polyglot](https://github.com/EleutherAI/polyglot/issues/37)

> I evaluated `polyglot-ko-1.3b` model with HellaSwag and WiC from KoBEST, and I got different results with paper and model card from huggingface.

EleutherAI의 점수 측정에 대한 문의를 GitHub Issue로 남기고 Discord에 그 내용을 공유하였습니다.

> **정한 Rycont**\
> [[Polyglot]] 논문에 제시된 HellaSwag, WiC 태스크의 점수가 재현되지 않는 문제를 발견했습니다. Ko-1.3b 모델로 테스트 하였으며, 자세한 내용은 깃허브에 이슈로 등록해두었습니다.

> 저도 다시 돌려봤는데 **@정한(Rycont)** 님이라 같은 결과가 나오네요. Polyglot-v1 때 논문에 참여하셨던 분들이 더 이상 채널에 안 계신 것 같아서 직접 여쭤보긴 조금 힘 들 것 같긴 합니다

> lm-eval-harness의 경우 대대적인 리팩토링 작업이 진행 중이고, 아마 1-2주 내로 main branch로 merge될 것 같은데요. 이에 따라 조만간 polyglot 브랜치에서도 upstream을 맞추고, 이전에 구현했던 task도 리뷰할 계획입니다. 그때 같이 살펴볼게요. 이슈 남겨주셔서 감사합니다
