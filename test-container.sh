#!/bin/bash
set -o allexport; source release.env; set +o allexport

docker rmi -f ${DOCKER_USER}/${IMAGE_BASE_NAME}:${IMAGE_TAG}

docker run \
    -v $(pwd)/samples:/samples \
    --rm ${DOCKER_USER}/${IMAGE_BASE_NAME}:${IMAGE_TAG}  \
    ./snippets generate \
    --input samples/js.yml \
    --output samples/js.code-snippets 
