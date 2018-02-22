docker build --tag meta .

docker run --rm \
       -v "$PWD"/src:/go/src/meta -v "$PWD"/bin:/go/bin \
       -w /go/src/meta \
       -e GOOS=darwin -e GOARCH=386 \
       meta go build -o /go/bin/meta

if [ ! -f /usr/local/bin/meta ]; then
    ln -sf ~/Code/meta/bin/meta /usr/local/bin/meta
fi
