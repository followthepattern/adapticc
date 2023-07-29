#!/usr/bin/env bash

APP_NAME=$MODE-$CI_PROJECT_NAME

TAG=${CI_COMMIT_TAG:-$CI_COMMIT_SHORT_SHA}

LOG_PREFIX="ADAPTICC"

echo "[$LOG_PREFIX]: removing $APP_NAME image!"
docker rm --force $APP_NAME
if [[ $? -eq 0 ]]; then
    echo "[$LOG_PREFIX]: image removed!"
fi

set -euxo pipefail

ADAPTICC_INTERNAL_COUNT=`docker network ls | grep fp-internal | wc -l`

if [[ $ADAPTICC_INTERNAL_COUNT -eq 0 ]]; then
    docker network create fp-internal
fi

docker pull $CI_REGISTRY_IMAGE:$TAG

docker run -d --network=fp-internal \
-e ADAPTICC_VERSION="$TAG" \
-e ADAPTICC_MAIL_USERNAME="$ADAPTICC_SMTP_USERNAME" \
-e ADAPTICC_MAIL_PASSWORD="$ADAPTICC_SMTP_PASSWORD" \
-e ADAPTICC_DB_PASSWORD="$ADAPTICC_DB_PASSWORD" \
-e ADAPTICC_SERVER_HMAC_SECRET="$ADAPTICC_SERVER_HMAC_SECRET" \
--name $APP_NAME $CI_REGISTRY_IMAGE:$TAG
if [[ $? -eq 0 ]]; then
    echo "[$LOG_PREFIX]: $APP_NAME container running!!"
    exit 0
else
    exit $?
fi