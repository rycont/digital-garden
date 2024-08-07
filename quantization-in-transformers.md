---
title: 양자화 할 때 Transformer의 내부 동작 살펴보기
date: 2023-11-27T13:52:58.844Z
---

# 양자화 할 때 [[Transformer]]의 내부 동작 살펴보기

1. \`from_pretrained\` 메소드에서 \`load_in_8bit\`를 \`True\`로 설정하여 모델을 로드한다

   1. \[Function\] replace_with_bnb_linear (src/transformers/integrations/bitsandbytes.py)

      - Arguments: (model)

      - Keywords

        1. modules_to_not_convert

        2. current_key_name

        3. quantization_config

      - 이 함수가 하는 일: \`model\` 파라미터로 넘겨받은 모델에서 \`torch.nn.Linear\`를 찾고 \`bnb.nn.Linear8bit\`로 변환한다.

        1. \`model.named_children\` 메소드를 통해 재귀적으로 레이어를 탐색한다

        2. 레이어가 \`Linear\` 혹은 \`Conv1D\`의 instance인지 확인한다

        3. \`bnb.nn.Linear8bitLt\` 레이어로 교체한다. 입력 값 갯수, 출력 값 갯수, 편향 사용 여부를 그대로 복사한다.

           - \`bnb.nn.Linear8bitLt\`: Bits and Bytes 모듈에서 제공하는 8bit linear layer.

           - Bits and Bytes(Bitsandbytes): 8bit 연산에 관련한 CUDA 함수 래핑 라이브러리

      - 이 과정에서 \`lm_head\`는 양자화되지 않는다. 안정성 이슈가 있다고 함.
