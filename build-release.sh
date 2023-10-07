#!/bin/bash
set -o allexport; source release.env; set +o allexport

echo -n "${APPLICATION_NAME} ${TAG} ${NICK_NAME}" > ./version.txt

mkdir -p release

env CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o ${APPLICATION_NAME}-${TAG}-darwin-arm64
env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ${APPLICATION_NAME}-${TAG}-darwin-amd64
env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o ${APPLICATION_NAME}-${TAG}-linux-arm64
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${APPLICATION_NAME}-${TAG}-linux-amd64
mv ${APPLICATION_NAME}-${TAG}-* ./release
