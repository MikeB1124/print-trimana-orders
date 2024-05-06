#!/bin/bash

echo "Stopping trimana container if it exists"
docker stop trimana

echo "Removing trimana container if it exists"
docker rm trimana

echo "Building new trimana-image"
docker build -t trimana-image ./

echo "Running trimana container with trimana-image"
docker run --name trimana -d trimana-image

echo "Exec into trimana container"
docker exec -it trimana sh