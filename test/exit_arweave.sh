#!/bin/sh
for id in $(docker ps -q)
do
    if [[ $(docker port "${id}") == *$1* ]]; then
        echo "stopping container ${id}"
        docker stop "${id}"
    fi
done