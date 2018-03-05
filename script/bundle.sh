#!/bin/sh -e

ROOT="$(cd "$(dirname "$0")"/.. && pwd)"

echo "Building docker image ..."
docker build --quiet --tag meta:latest $ROOT

rm -rf /tmp/meta
mkdir -p /tmp/meta
cp -r "$ROOT"/config /tmp/meta/
cp -r "$ROOT"/etc/bash_completion /tmp/meta/

echo "Building Mac bundle ..."
docker run --rm \
       -v /tmp/meta:/go/bin \
       -e GOOS=darwin \
       -e GOARCH=386 \
       meta:latest \
       go build -o /go/bin/meta ./cmd
tar -czf "$ROOT"/bundle/meta.mac.tar.gz -C /tmp meta

echo "Building Linux bundle ..."
docker run --rm \
       -v /tmp/meta:/go/bin \
       meta:latest \
       go build -o /go/bin/meta ./cmd
tar -czf "$ROOT"/bundle/meta.linux.tar.gz -C /tmp meta
rm -rf /tmp/meta
