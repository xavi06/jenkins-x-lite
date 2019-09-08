FROM golang:1.11.12-stretch AS builder

# add proxy
ENV http_proxy http://172.16.133.102:1087
ENV https_proxy http://172.16.133.102:1087

WORKDIR /go/src/github.com/xavi06/jenkins-x-lite
COPY go.mod /go/src/github.com/xavi06/jenkins-x-lite/
COPY go.sum /go/src/github.com/xavi06/jenkins-x-lite/
ENV GO111MODULE on
RUN go mod download
# copy code
COPY . /go/src/github.com/xavi06/jenkins-x-lite
RUN go build -o jxl

FROM debian:stretch
COPY --from=builder /go/src/github.com/xavi06/jenkins-x-lite/jxl /usr/bin/jxl
CMD ["/usr/bin/jxl"]
