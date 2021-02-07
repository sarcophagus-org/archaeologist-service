# Archaeologist Service
---
This respository contains the archaeologist service. 

This service is responsible for executing Sarcophagus jobs created via the Sarcophagus contract:
https://github.com/sarcophagus-org/smart-contracts

### Quick Start
You must have an arweave wallet (with Arweave tokens) to accept new jobs. [You can get a wallet here.](https://www.arweave.org/wallet)

```
git clone https://github.com/sarcophagus-org/archaeologist-service
cp config.example.yml config.yml
go build
```

Update the config.yml file with the appropriate values. 
- Use the comments in the config.example.yml file as a guide. 
- **The config.yml must be writable.**
- All comments in config.yml will be removed after the first time the service is run.
- Put your arweave wallet file in a path on the filesystem, and update the config.yml `areave_key_file` value to point to the location of your arweave wallet. (e.g "/usr/local/arweave.json")

To run the service:
```
./archaeologist-service
```

**Alternatively you can install the service globally with:**

```
go install
```

Copy the config file to $GOPATH/bin (assuming you have `$GOPATH/bin` in your $PATH)
```
cp config.example.yml config.yml
cp config.yml $GOPATH/bin
```

Update the config.yml `areave_key_file` value to point to the location of your arweave wallet.

If installing on Ubuntu, before building/installing, you may need to run:
```
apt-get install build-essential
```

### Local Development 
##### Deploy Sarcophagus Contract
Clone https://github.com/sarcophagus-org/smart-contracts

Follow Readme directions to deploy the contract locally.

##### Fill out values in config file
See config/config.example.yml for descriptions and examples for each config key/value.

##### Run Arweave Blockchain
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
curl -d '{"beneficiary": "<arweave wallet address>", "quantity": 1000000000000}' http://localhost:8000/loom/faucet
```

See https://github.com/rootmos/loom for more information

##### Redeploy Contract
If the Sarcophagus contract gets updated, you need to re-compile the contract for the service to use.

Install ethereum and solidity. Below are instructions for homebrew. See this link for other methods:
https://solidity.readthedocs.io/en/v0.5.3/installing-solidity.html

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

Use abigen to compile abi -- below examples output to "abi" & "abiToken" directories. Alternatively use abi in Sarcophagus Contracts build folder.
NOTE: The solidity compiler version must match the version specified in the contract (currently ^0.6.0)
```
solc --abi @openzeppelin/=$(pwd)/node_modules/@openzeppelin/ contracts/Sarcophagus.sol -o abi
solc --abi @openzeppelin/=$(pwd)/node_modules/@openzeppelin/ contracts/SarcophagusToken.sol -o abiToken
solc --abi @openzeppelin/=$(pwd)/node_modules/@openzeppelin/ contracts/libraries/Events.sol -o abiEvents
```

Compile Contracts to Go
```
abigen --abi=./abi/Sarcophagus.abi --pkg=sarcophagus --out=Sarcophagus.go
abigen --abi=./abiToken/SarcophagusToken.abi --pkg=token --out=SarcophagusToken.go
abigen --abi=./abiEvents/Events.abi --pkg=events --out=Events.go
```

1. Copy/Replace generated files into /contracts directory
2. Rename the package name on these files to "contracts"

If you have any issues with the compiled go code, you may need to download an older version of abigen.

https://geth.ethereum.org/downloads/

The latest tested version working with the service is Geth & Tools 1.9.25

##### Testing
An embalmer package is provided for local testing. 
Embalmer config values can be updated in `embalmer/embalmer_config.yml`
These config values are configured to work with the main service's test config at `test/test_config.yml`

Examples are below.
```
# Start the arch service. Config values must be set correctly and free bond must be added to accept new jobs.
go run main.go

# The embalmer must have a sufficient Sarco token balance.

# The seed flag is used to generate file bytes, which will be used as the asset file and generate the asset double hash for a Sarcophagus.
# This seed can be changed to create/modify different sarcophaguses. 

# Create a Sarcophagus
# Exluding the -type flag will set the type as "create" by defaut.
go run cmd/embalmer.go -seed=200

# Update a Sarcophagus
go run cmd/embalmer.go -seed=200 -type=update

# Rewrap a Sarcophagus
go run cmd/embalmer.go -seed=200 -type=rewrap

# Clean a Sarcophagus
go run cmd/embalmer.go -seed=200 -type=clean

# Bury a Sarcophagus
go run cmd/embalmer.go -seed=200 -type=bury

# Cancel a Sarcophagus
go run cmd/embalmer.go -seed=200 -type=cancel
```

There is a test suite provided that uses the embalmer package, see the README.md in the test directory for more directions.

##### Sending a file
If you want to test sending a file locally (the Sarcophagus payload) after creating a sarcophagus: 
```
curl -v -X POST -F file=@<your file> -F http://127.0.0.1:<your port>/file
```