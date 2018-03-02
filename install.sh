#!/bin/sh -e

DIR="$(cd "$(dirname "$0")" && pwd)"
ARCH=`uname -s`

echo "Create link to meta binary ..."
if [ $ARCH == 'Darwin' ]; then
    ln -sf $DIR/bin/meta-mac /usr/local/bin/meta
else
    ln -sf $DIR/bin/meta-linux /usr/local/bin/meta
fi

echo "Copy configuration to ~/.meta ..."
rm -rf ~/.meta
cp -r $DIR/config ~/.meta
