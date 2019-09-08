app := jxl
imagename := zj/golang/jxl
imagetag := latest

default: build

build:
	go build -o ~/go/bin/$(app)

docker:
	docker build -t $(imagename):$(imagetag) .

dockerrun:
	docker run -tid --rm --name jxl-container $(imagename):$(imagetag) sh

dockercopy:
	docker cp jxl-container:/usr/bin/$(app) bin/linux/$(app)

dockerstop:
	docker stop jxl-container

linux: docker dockerrun dockercopy dockerstop
macos:
	go build -o bin/macos/$(app)
