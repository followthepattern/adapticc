#!/usr/bin/env bash

APP_NAME=$MODE-$CI_PROJECT_NAME

echo "FP: removing $APP_NAME image!"
docker rm --force $APP_NAME
if [[ $? -eq 0 ]]; then
    echo "FP: image removed!"
fi

set -euxo pipefail

FP_INTERNAL_COUNT=`docker network ls | grep fp-internal | wc -l`

if [[ $FP_INTERNAL_COUNT -eq 0 ]]; then
    docker network create fp-internal
fi

docker pull $CI_REGISTRY_IMAGE:$CI_COMMIT_TAG

docker run -d --network=fp-internal \
-e FP_VERSION="$CI_COMMIT_TAG" \
-e FP_MAIL_USERNAME="$FP_SMTP_USERNAME" \
-e FP_MAIL_PASSWORD="$FP_SMTP_PASSWORD" \
-e FP_DB_PASSWORD="$FP_DB_PASSWORD" \
-e FP_SERVER_HMAC_SECRET="$FP_SERVER_HMAC_SECRET" \
--name $APP_NAME $CI_REGISTRY_IMAGE:$CI_COMMIT_TAG
if [[ $? -eq 0 ]]; then
    echo "FP: $APP_NAME container running!!"
    exit 0
else
    exit $?
fi