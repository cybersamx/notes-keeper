#!/usr/bin/env bash

DOCKER_NAME='postgres-to-do'

# Stop the current postgres container if it's running
docker_id=$(docker ps --filter "name=${DOCKER_NAME}" --filter 'status=running' --format '{{.ID}}')
[[ -z "${docker_id}" ]] || {
    docker stop "${docker_id}"
}

# Start a postgres container
docker pull postgres:12.0-alpine
docker run \
    --name "${DOCKER_NAME}" \
    --rm \
    -d \
    -p 5432:5432 \
    -e POSTGRES_PASSWORD=not-secure-pwd \
    -e POSTGRES_USER=app-user \
    -e POSTGRES_DB=app \
    postgres
docker ps -a
