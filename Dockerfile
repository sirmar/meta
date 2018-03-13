FROM golang:1.9-alpine

# Os dependencies
RUN apk add --update --no-cache git

WORKDIR /go/src/meta

# Golang dev dependencies
RUN go get \
    github.com/stretchr/testify \
    github.com/vektra/mockery/.../

# Golang production dependencies
RUN go get \
    github.com/devfacet/gocmd \
    gopkg.in/yaml.v2 \
    github.com/mitchellh/go-homedir \
    github.com/blang/semver

# Copy source
COPY . .

# Build binary
RUN cp -r config ~/.meta
RUN go build -o /go/bin/meta ./cmd

CMD /go/bin/meta
