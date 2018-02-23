#!/bin/bash -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

docker build --quiet --tag meta:latest $DIR

docker run --rm \
       -w /go/src/app \
       -v "$DIR"/.cache:/go \
       -v "$DIR":/go/src/app \
       -e GOOS=darwin \
       -e GOARCH=386 \
       meta:latest \
       go get

docker run --rm \
       -w /go/src/app \
       -v "$DIR"/.cache:/go \
       -v "$DIR":/go/src/app \
       -v "$DIR"/bin:/go/bin \
       -e GOOS=darwin \
       -e GOARCH=386 \
       meta:latest \
       go build -o /go/bin/meta

if [ ! -f /usr/local/bin/meta ]; then
    ln -sf ~/Code/meta/bin/meta /usr/local/bin/meta
fi
