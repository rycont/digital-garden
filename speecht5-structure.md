---
title: 2023-11-14
layout: ../layouts/article.astro
date: 2023-11-14T09:51:08.414Z
---

# 2023-11-14

- [[T5|SpeechT5]] 분석

  - DecoderWithTextPrenet

    - Decoder의 앞에 Prenet을 붙힘.

    - Hidden State가 아닌, deocder에 들어오는 input ids를 처리해주기 위함

    - 그래서 모델의 전체적인 구조로 보면

      - encoder prenet - encoder - decoder - decoder postnet

      - 형태가 됨

  - 자 만약에 여기에 Image 도메인을 추가하고 싶다면? 총 세개의 네트워크를 추가로 구현해야 함

    - **Encoder Image Prenet**\
      이미지를 Feature Vectors로 변환

    - **Decoder Image Prenet**\
      이미지를 Feature Vectors로 변환

    - **Decoder Image Postnet**\
      Feature Vectors를 이미지로 재구성

    - Image를 hidden vector로 변환하고 다시 reconstruct할 수 있는 모델? VitMAE
