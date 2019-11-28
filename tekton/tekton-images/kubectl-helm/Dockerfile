FROM debian:stretch-slim

USER root

RUN apt-get update  
RUN apt-get install -y ca-certificates

RUN mkdir -p /root/.kube

COPY helm /usr/local/bin/
COPY kubectl /usr/local/bin/

COPY config /root/.kube/config

RUN helm init --client-only && \
    helm repo add chartmuseum http://10.30.64.99:8088
