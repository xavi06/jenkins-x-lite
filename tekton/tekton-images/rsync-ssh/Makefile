IMAGE := zj/tekton/rsync-ssh
TAG := 1.0.0
REGISTRY := jxregistry.docker.fccs.cn:28888
#REGISTRY := registry.cn-hangzhou.aliyuncs.com
DOCKERFILE := Dockerfile

build:
	docker build -f $(DOCKERFILE) -t $(IMAGE):$(TAG) .

tag:
	docker tag $(IMAGE):$(TAG) $(REGISTRY)/$(IMAGE):$(TAG)

push:
	docker push $(REGISTRY)/$(IMAGE):$(TAG)

release: build tag push

