#!/bin/bash

# debug
#set -xe

export DOCKER_REGISTRY=${DOCKER_REGISTRY:-localhost}
export DOCKER_REGISTRY_IMAGE_PATH=${DOCKER_REGISTRY_IMAGE_PATH:-cryptastic/chia-price-exporter}
export PLATFORM=${PLATFORM:-linux/amd64,linux/arm64}
export CHIA_PRICE_EXPORTER_VERSION=${CHIA_PRICE_EXPORTER_VERSION:-dev}

export DOCKER_REGISTRY_IMAGE=$DOCKER_REGISTRY/$DOCKER_REGISTRY_IMAGE_PATH

docker buildx create --use --name docker-multiarch --platform "$PLATFORM"
docker buildx build \
  --platform "$PLATFORM" \
  --tag "$DOCKER_REGISTRY_IMAGE:latest" \
  --tag "$DOCKER_REGISTRY_IMAGE:$CHIA_PRICE_EXPORTER_VERSION" \
  --cache-from "$DOCKER_REGISTRY_IMAGE:latest" \
  --pull \
  --push \
  ./
