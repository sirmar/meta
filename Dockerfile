FROM golang:1.8-alpine

RUN apk update && \
    apk add --no-cache git docker

COPY src /go/src/meta

WORKDIR /go/src/meta

RUN go get
