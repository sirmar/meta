#!/bin/sh

DIR="$(cd "$(dirname "$0")" && pwd)"

echo "Building docker image ..."
docker build --quiet --tag meta:latest $DIR

echo "Building Mac binary ..."
docker run --rm \
       -v "$DIR"/bin:/go/bin \
       -e GOOS=darwin \
       -e GOARCH=386 \
       meta:latest \
       go build -o /go/bin/meta-mac

echo "Building Linux binary ..."
docker run --rm \
       -v "$DIR"/bin:/go/bin \
       meta:latest \
       go build -o /go/bin/meta-linux
