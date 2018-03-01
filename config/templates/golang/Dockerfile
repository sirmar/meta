FROM golang:1.9-alpine

# Os dependencies
RUN apk add --update --no-cache git

WORKDIR /go/src/{{name}}

# Golang dev dependencies
RUN go get \
    github.com/stretchr/testify \
    github.com/vektra/mockery/.../

# Golang production dependencies
#RUN go get

# Copy source
COPY . .

# Build binary
RUN go build
