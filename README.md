# Archaeologist Service
---
This respository contains the archaeologist service which 

## Local Development 
### Deploy Sarcophagus Contract
Clone https://github.com/decent-labs/airfoil-sarcophagus-v2-contracts

Follow Readme directions to deploy the contract locally.

#### Fill out values in config file
TODO: Add config values and descriptions

### Run Arweave Blockchain
```
docker pull rootmos/loom
docker run --rm --publish 8000:8000 rootmos/loom
```

##### Create Arweave Key
```
arweave key-create your-arweave-key.json
```

1. Copy the generated key to /config directory.
2. Update the config/config.yml file "AREWAVE_KEY_FILE" to be the name of this file.

##### Add tokens from faucet to your wallet address

```
curl -d '{"beneficiary": "kSyg2ajZbqiAE25AnkQcxHoGOnM5jUgT4t9TyZZQI3I", "quantity": 1000000000000}' http://localhost:8000/loom/faucet
```

See https://github.com/rootmos/loom for more information

### Redeploy Contract
If the Sarcophagus contract gets updated, you need to re-compile the contract for the service to use.

Install ethereum and solidity. Below are instructions for homebrew. See this link for other methods:
https://solidity.readthedocs.io/en/v0.5.3/installing-solidity.html

NOTE: The solidity compiler version must match the version specified in the contract (currently ^0.6.0)

```
brew update
brew upgrade
brew tap ethereum/ethereum
brew install solidity
```

Install abigen
```
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
make
make devtools
```

Use abigen to compile abi -- below example outputs to "abi" directory
```
solc --abi @openzeppelin/=/<local absolute path to repo>/node_modules/@openzeppelin/ contracts/Sarcophagus.sol -o abi
```

