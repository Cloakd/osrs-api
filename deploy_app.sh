#!/bin/bash
# Usage template : ./build_images.sh app-name     image                                tag   valuesFile
# Usage example  : ./build_images.sh item_scraper eu.gcr.io/hague-hosting/app-name     1.0.0 ./chart/single-chart/valuesFiles/file.yaml
# Usage example  : ./build_images.sh item_scraper eu.gcr.io/hague-hosting/item-scraper 1.0.0 ./chart/single-chart/valuesFiles/item-scraper.yaml

## Requirements
# You will require the following tools installed;
# google cloud sdk (https://cloud.google.com/sdk/docs/downloads-interactive)
# helm3 (version 3.1.2 is good; https://github.com/helm/helm/releases/tag/v3.1.2)

## Be sure to also run the following commands once gcloud sdk is installed;
# $ gcloud auth login
# $ gcloud auth configure-docker
# $ gcloud components install kubectl

APP_NAME=$1
DOCKER_IMAGE=$2
DOCKER_TAG=$3
VALUES_FILE=$4

echo "Checking for GKE Cluster"
gcloud container clusters get-credentials hague-hosting --zone europe-west2-b --project hague-hosting > /dev/null
if [[ $? -ne 0 ]]; then
    echo "[ERROR] Issue obtaining the gke cluster credentials."
    exit 1
fi

kubectl config current-context | grep hague-hosting > /dev/null
if [[ $? -ne 0 ]]; then
    echo "[ERROR] Incorrect cluster found... problem with script."
    exit 1
fi

helm3 upgrade --install --set image.repository="${DOCKER_IMAGE}",image.tag="${DOCKER_TAG}" --values "${VALUES_FILE}" "${APP_NAME}" ./chart-template