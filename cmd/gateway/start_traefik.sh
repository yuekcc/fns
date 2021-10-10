#!/bin/env bash

docker run -d --rm \
    -p 80:80 \
    -p 9090:8080 \
    -v $PWD/traefik.yml:/etc/traefik/traefik.yml \
    -v /var/run/docker.sock:/var/run/docker.sock \
    --label "traefik.enable=false" \
    traefik:v2.5
