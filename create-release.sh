#!/bin/bash
set -o allexport; source release.env; set +o allexport
echo "TAG: ${TAG}"
echo "IMAGE_TAG: ${IMAGE_TAG}"
echo "IMAGE_BASE_NAME: ${IMAGE_BASE_NAME}"

echo "📦 Create release..."
git add .
git commit -m "📦 create release ${TAG}"
git tag ${TAG}
git push origin main ${TAG}
