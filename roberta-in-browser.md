---
title: 브라우저에서 RoBERTa 실행해보기
layout: ../layouts/article.astro
date: 2024-04-02T07:19:54.046Z
---

# 브라우저에서 RoBERTa 실행해보기

[](RoBERTa)를 사용하면 텍스트 분류, 검색, 특징 탐지과 같은 기능을 구현할 수 있습니다. 기술의 발전으로 인해 가벼운 [](언어모델)들은 [](브라우저)에서 직접 실행할 수 있게 됐습니다.

> 이렇게 인공지능 모델을 사용자 디바이스에서 직접 실행하는 기술을 **온디바이스 AI(On-device AI)**라고 합니다.
>
> [일상을 바꿀 내 손 안의 AI, 온디바이스 AI – 제일기획 매거진](https://magazine.cheil.com/54980)

### Transformers.js

자바스크립트 런타임에서 [](Transformer) 언어모델을 실행할 수 있게 하는 라이브러리입니다. 특히 브라우저에서도 실행이 가능합니다. 이 라이브러리를 사용해서 구현할 예정입니다. Transformers.js 라이브러리는 [](ONNX) 포맷으로 저장된 모델을 로드할 수 있습니다.

하지만 대부분의 언어 모델은 pth(PyTorch), h5(TensorFlow), safetensors(SafeTensors)로 배포됩니다. 그렇기에 ONNX로 변환을 해야 Transformers.js에서 사용할 수 있습니다.

### Transformers 모델을 브라우저가 이해할 수 있도록 PyTorch ONNX로 변환하기

시범으로 [leewaay/klue-roberta-base-klueNLI-klueSTS-MSL128](https://huggingface.co/leewaay/klue-roberta-base-klueNLI-klueSTS-MSL128) 모델을 변환해보겠습니다. [](Huggingface)의 Optimum이라는 라이브러리를 사용하면 손쉽게 변환할 수 있습니다.

먼저 필요한 라이브러리를 설치해주겠습니다.

```prompt
pip install optimum onnxruntime sentence_transformers onnx
```

언어모델의 가중치를 ONNX로 변환해보겠습니다

```python
from optimum.onnxruntime import ORTModelForSequenceClassification, ORTModelForFeatureExtraction
from transformers import AutoTokenizer

model_checkpoint = "leewaay/klue-roberta-base-klueNLI-klueSTS-MSL128"
save_directory = "tmp/onnx/"

ort_model = ORTModelForFeatureExtraction.from_pretrained(model_checkpoint, export=True)
tokenizer = AutoTokenizer.from_pretrained(model_checkpoint)

ort_model.save_pretrained(save_directory)
tokenizer.save_pretrained(save_directory)
```

1분 내외로 `./tmp/onnx` 디렉토리에 ONNX 모델 파일이 저장됩니다. 다음과 같은 결과물을 확인할 수 있습니다.

![](../images/826cd39a-0b98-4310-b924-3e4553bcb29f.png)

- tmp/onnx

  - config.json

  - model.onnx

  - special_tokens_map

  - tokenizer_config.json

  - tokenizer.json

  - vocab.txt

### (선택) 양자화로 모델 크기 줄이기

400MB는 앱의 일부로 배포하기에는 부담스러운 크기입니다. 모델의 크기도 줄이고 추론 속도도 높히려면 [](양자화)를 수행할 수 있습니다. 다음 코드로 모델을 양자화할 수 있습니다.

> **같이 보기:** [**딥러닝 모델 쉽게 양자화하기**](easy-quantization)
>
> 양자화는 딥러닝 모델을 경량화하는 방법중 하나입니다. Huggingface의 transformers 라이브러리가 제공하는 툴킷을 사용하면 Huggingface에 업로드된 모델을 간편하게 양자화할 수 있습니다...

```python
from optimum.onnxruntime import ORTQuantizer, AutoQuantizationConfig

quantizer = ORTQuantizer.from_pretrained(ort_model)
dqconfig = AutoQuantizationConfig.avx512_vnni(
  is_static = False,
  per_channel = False
)

model_quantized_path = quantizer.quantize(
    save_dir = save_directory,
    quantization_config = dqconfig,
)
```

위 코드를 실행하면 \`save_directory\`에 \`model_quantized.onnx\`라는 파일이 추가로 생성됩니다. <u>양자화 이후에 크기가 110MB 가량으로 줄었습니다.</u>

### Transformers.js에서 ONNX 모델을 실행하기

transformers.js를 설치해주겠습니다.

```prompt
pnpm install @xenova/transformers
```

다음과 같이 모델을 로드할 수 있습니다.

```javascript
import { AutoTokenizer, AutoModel } from "@xenova/transformers";

const modelName = "model-directory";

const tokenizer = await AutoTokenizer.from_pretrained(modelName);
const model = await AutoModel.from_pretrained(modelName, {
  quantized: false, // 양자화를 수행했다면 생략할 수 있음
});
```

#### 문장 임베딩 계산하기

Sentence Transformers 모델은 mean pooling으로 문장 임베딩을 계산할 수 있다.

```javascript
import { mean_pooling } from "@xenova/transformers";

async function getSentenceEmbedding(text) {
  const inputs = await tokenizer(text);
  const result = await model(inputs);

  const attentionMask = inputs.attention_mask;
  const pooled = mean_pooling(result.last_hidden_state, attentionMask);

  return pooled.data;
}
```

#### 문장 간 유사도 계산하기

```javascript
import { cos_sim } from "@xenova/transformers";

const s1 = await getSentenceEmbedding("난 대학시절 묵찌빠를 전공했단 사실");
const s2 = await getSentenceEmbedding("난 묵찌빠로 유학까지 다녀왔단 사실");
const s3 = await getSentenceEmbedding("네 놈을 이겨 가문의 이름 높이리");

console.log(cos_sim(s1, s2)); // 0.6303152271221659
console.log(cos_sim(s1, s3)); // 0.13926317908806657
```
