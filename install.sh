#!/bin/sh

DIR="$(cd "$(dirname "$0")" && pwd)"
ARCH=`uname -s`

docker build --quiet --tag meta:latest $DIR

if [ $ARCH == 'Darwin' ]; then
   docker run --rm \
       -v "$DIR":/go/src/meta \
       -v "$DIR"/installed:/go/bin \
       -e GOOS=darwin \
       -e GOARCH=386 \
       meta:latest \
       go build -o /go/bin/meta
else
   docker run --rm \
       -v "$DIR":/go/src/meta \
       -v "$DIR"/installed:/go/bin \
       meta:latest \
       go build -o /go/bin/meta
fi

if [ ! -f /usr/local/bin/meta ]; then
    ln -sf ~/Code/meta/installed/meta /usr/local/bin/meta
fi
