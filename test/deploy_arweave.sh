#!/bin/sh
for id in $(docker ps -q)
do
    if [[ $(docker port "${id}") == *"8000"* ]]; then
        echo "stopping container ${id}"
        docker stop "${id}"
    fi
done
docker run --rm --publish 8000:8000 rootmos/loom
Beneficiary=$1
curl -d "{'beneficiary': ${Beneficiary}, 'quantity': 100000000000000}" http://localhost:8000/loom/faucet