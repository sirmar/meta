#!/bin/bash -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

docker build --quiet --tag meta:latest $DIR

docker run --rm \
       -v "$DIR":/go/src/meta \
       -v "$DIR"/installed:/go/bin \
       -e GOOS=darwin \
       -e GOARCH=386 \
       meta:latest \
       go build -o /go/bin/meta

if [ ! -f /usr/local/bin/meta ]; then
    ln -sf ~/Code/meta/installed/meta /usr/local/bin/meta
fi
