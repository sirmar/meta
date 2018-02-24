FROM golang:1.9-alpine

# Os dependencies
RUN apk add --update --no-cache \
    git

WORKDIR /go/src/meta

# Golang dev dependencies
RUN go get \
    github.com/stretchr/testify \
    github.com/vektra/mockery/.../

# Golang production dependencies
RUN go get \
    github.com/akamensky/argparse \
    gopkg.in/yaml.v2
