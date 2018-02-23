DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

docker build --squash --tag meta:latest $DIR

docker run --rm \
       -w /go/src/meta \
       -v "$DIR"/.cache:/go/src \
       -v "$DIR"/src:/go/src/meta \
       meta:latest \
       go get -v

docker run --rm \
       -w /go/src/meta \
       -v "$DIR"/.cache:/go/src \
       -v "$DIR"/src:/go/src/meta \
       -v "$DIR"/bin:/go/bin \
       -e GOOS=darwin \
       -e GOARCH=386 \
       meta:latest \
       go build -o /go/bin/meta

if [ ! -f /usr/local/bin/meta ]; then
    ln -sf ~/Code/meta/bin/meta /usr/local/bin/meta
fi
