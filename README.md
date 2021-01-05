# Image-scanning-webhook

> Image-scanning-webhook for HyperCloud Service

## prerequisite Install
- ElasticSearch 7.X


## Build Image-scanning-webhook
1. binary build
    - make build
2. docker-image build
    - make image
3. docker-image push
    - make push
4. 비고
    - Makefile의 REGISTRY 및 version은 사용자에 맞게 변경해야 합니다.

## Install Image-scanning-webhook
> 이미지 스캐닝 오퍼레이터에서 처리한 취약점 결과를 ElasticSearch에 저장하기 위한 웹훅 서버 입니다.
1. Image-scanning-webhook를 설치하기 위한 네임스페이스를 생성 합니다.
    - kubectl create namespace {YOUR_NAMESPACE}
    - ex) kubectl create namespace image-scanning
2. 아래의 command로 Image-scanning-webhook을 생성 합니다.
    - kubectl apply -f webhook.yaml -n {YOUR_NAMESPACE} ([파일](./deploy/webhook.yaml))
    - 비고: deployment 내부의 image 경로는 사용자 환경에 맞게 수정 해야 합니다.
    - 비고: deployment 내부의 env중 ELASTIC_SEARCH_URL 은 사용자 환경에 맞게 변경해야 합니다. (예시: http://192.168.X.X:6060)