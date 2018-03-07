#!/bin/sh -e

ROOT="$(cd "$(dirname "$0")"/.. && pwd)"
ARCH=`uname -s`

echo "Building docker image ..."
docker build --quiet --tag meta:latest $ROOT

echo "Building binary  ..."
if [ $ARCH == 'Darwin' ]; then
    docker run --rm \
	   -v /tmp:/go/bin \
	   -e GOOS=darwin \
	   -e GOARCH=386 \
	   meta:latest \
	   go build -o /go/bin/meta ./cmd
else
    docker run --rm \
	   -v /tmp:/go/bin \
	   meta:latest \
	   go build -o /go/bin/meta ./cmd
fi
mv -f /tmp/meta /usr/local/bin/meta

echo "Install configuration files  ..."
rm -rf ~/.meta
cp -r "$ROOT"/config ~/.meta

echo "Run: 'source ~/.meta/bash_completion' to activate completion."
