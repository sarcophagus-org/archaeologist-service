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
Install arweave-deploy
https://docs.arweave.org/developers/tools/arweave-deploy

Generate arweave key:
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

__TODO: Work on automating the process below using go generate: https://geth.ethereum.org/docs/dapp/native-bindings
.__

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

Use abigen to compile abi -- below examples output to "abi" / "abiToken" directories. Alternatively use abi in build folder.
```
solc --abi @openzeppelin/=/<local absolute path to repo>/node_modules/@openzeppelin/ contracts/Sarcophagus.sol -o abi
solc --abi @openzeppelin/=/<local absolute path to repo>/node_modules/@openzeppelin/ contracts/SarcophagusToken.sol -o abiToken
```

Compile Contracts to Go

```
abigen --abi=./abi/Sarcophagus.abi --pkg=contracts --out=Sarcophagus.go
abigen --abi=./abiToken/SarcophagusToken.abi --pkg=contracts --out=SarcophagusToken.go
```

Copy/Replace generated files into /contracts directory

If you have any issues with the compiled go code, you may need to download an older version of abigen.

https://geth.ethereum.org/downloads/

The latest tested version working with the service is Geth & Tools 1.9.22

### Sending a file
If you want to test sending a file locally (the Sarcophagus payload) after creating a sarcophagus: 

```
curl -v -X POST -F file=@<your file> -F "signedAssetDoubleHash=<your double hash>" http://127.0.0.1:<your port>/file
```