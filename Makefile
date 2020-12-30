REGISTRY      ?= 192.168.6.122:5000
VERSION       ?= 0.0.1

PACKAGE_NAME  = github.com/jitaeyun/image-scanning-webhook

IMAGE_SCANNING_WEBHOOK_NAME  = image-scanning-webhook
IMAGE_SCANNING_WEBHOOK_IMG   = $(REGISTRY)/$(IMAGE_SCANNING_WEBHOOK_NAME):$(VERSION)

BIN = ./build/_output/bin

.PHONY: build build-isw
build: build-isw
build-isw:
	GOOS=linux CGO_ENABLED=0 go build -o $(BIN)/image-scanning-webhook $(PACKAGE_NAME)/pkg

.PHONY: image image-isw
image: image-isw
image-isw:
	docker build -f build/Dockerfile -t $(IMAGE_SCANNING_WEBHOOK_IMG) .

.PHONY: push push-isw
push: push-isw
push-isw: 
	docker push $(IMAGE_SCANNING_WEBHOOK_IMG)
