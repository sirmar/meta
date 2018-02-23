FROM golang:1.8-alpine

RUN apk add --update --no-cache git

WORKDIR /go/src/meta
