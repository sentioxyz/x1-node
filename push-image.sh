#!/bin/bash

export DOCKER_DEFAULT_PLATFORM=linux/amd64

make build-docker
docker tag xlayer-node:latest ${SENTIO_DOCKER_REPO}/xlayer-node:v0.3.9
docker push ${SENTIO_DOCKER_REPO}/xlayer-node:v0.3.9