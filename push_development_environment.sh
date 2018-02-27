#!/bin/sh

docker build --file Dockerfile.devenv --tag sirmar/development-environment .
docker login -u sirmar
docker push sirmar/development-environment
