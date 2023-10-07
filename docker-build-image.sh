#!/bin/bash
set -o allexport; source release.env; set +o allexport

sudo chmod 666 /var/run/docker.sock
ls -l /var/run/docker.sock
sudo systemctl restart docker.service

echo -n "${APPLICATION_NAME} ${TAG} ${NICK_NAME}" > ./version.txt
#mkdir -p release

env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o linux/arm64/${APPLICATION_NAME}
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o linux/amd64/${APPLICATION_NAME}

docker login -u ${DOCKER_USER} -p ${DOCKER_PWD}

docker buildx create --use
docker buildx build -t ${DOCKER_USER}/${IMAGE_BASE_NAME}:${IMAGE_TAG} --platform=linux/arm64,linux/amd64 . --push
