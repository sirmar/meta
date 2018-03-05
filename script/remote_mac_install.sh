#!/bin/sh -e

cd /tmp
echo "Download bundle ..."
wget -q https://github.com/sirmar/meta/raw/master/bundle/meta.mac.tar.gz

echo "Extract bundle ..."
tar -xzf ./meta.mac.tar.gz

echo "Move binary to /usr/local/bin ..."
mv -f ./meta/meta /usr/local/bin/meta

echo "Move configuration to ~/.meta ..."
mv -f ./meta/config ~/.meta

echo "Remove tmp files ..."
rm meta.mac.tar.gz
rm -rf /tmp/meta

echo "Run: 'source ~/.meta/bash_completion' to activate completion."
