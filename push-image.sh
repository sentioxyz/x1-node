#!/bin/bash

export DOCKER_DEFAULT_PLATFORM=linux/amd64

make build-docker
docker tag x1-node:latest ${SENTIO_DOCKER_REPO}/x1-node:v0.1.1
docker push ${SENTIO_DOCKER_REPO}/x1-node:v0.1.1
