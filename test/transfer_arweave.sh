#!/bin/sh
curl -d '{"beneficiary": "'"$1"'", "quantity": 100000000000000}' http://localhost:8000/loom/faucet