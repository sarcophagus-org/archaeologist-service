#!/bin/sh
for id in $(docker ps -q)
do
    if [[ $(docker port "${id}") == *"8000"* ]]; then
        echo "stopping container ${id}"
        docker stop "${id}"
    fi
done