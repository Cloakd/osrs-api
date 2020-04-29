#!/bin/bash
# Usage template : ./build_images.sh app-name     docker-tag
# Usage example  : ./build_images.sh item_scraper v1.0.0
# Usage example  : ./build_images.sh item_scraper v1.0.1

## Requirements
# You will require the following tools installed;
# google cloud sdk (https://cloud.google.com/sdk/docs/downloads-interactive)
# helm3 (version 3.1.2 is good; https://github.com/helm/helm/releases/tag/v3.1.2)

## Be sure to also run the following commands once gcloud sdk is installed;
# $ gcloud auth login
# $ gcloud auth configure-docker
# $ gcloud components install kubectl

# env vars
DOCKER_REG_URL="eu.gcr.io"
DOCKER_REG_REPO="hague-hosting"
DOCKER_REG_APP_NAME=$1
DOCKER_REG_TAG=$2
FULL_DOCKER_URL="${DOCKER_REG_URL}/${DOCKER_REG_REPO}/${DOCKER_REG_APP_NAME}:${DOCKER_REG_TAG}"

echo "Building docker images with"
echo "docker build -q -t "${FULL_DOCKER_URL}" ."
docker build -q -t "${FULL_DOCKER_URL}" . > /dev/null 2>&1
BUILD_RC=$?

if [[ "${BUILD_RC}" -eq 0 ]]; then
    echo "Pushing docker images"
    docker push "${FULL_DOCKER_URL}"
else
    echo "Docker build failed, nothing to push"
fi