#!/bin/sh -e

cd /tmp
echo "Download bundle ..."
wget -q https://github.com/sirmar/meta/raw/master/bundle/meta.mac.tar.gz

echo "Extract bundle ..."
tar -xzf /tmp/meta.mac.tar.gz

echo "Move binary to /usr/local/bin ..."
mv -f ./meta/meta /usr/local/bin/meta

echo "Move configuration to ~/.meta ..."
mv -f ./meta/config ~/.meta

echo "Move bash completion to /usr/local/etc/bash_completion.d/meta ..."
mv -f ./meta/bash_completion /usr/local/etc/bash_completion.d/meta

echo "Remove tmp files ..."
rm meta.mac.tar.gz
rm -rf /tmp/meta
