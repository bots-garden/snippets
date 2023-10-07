#!/bin/bash
set -o allexport; source release.env; set +o allexport
echo "TAG: ${TAG}"
git tag -d ${TAG}
