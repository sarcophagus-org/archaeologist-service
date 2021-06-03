# Archaeologist Service

[![Discord](https://img.shields.io/discord/753398645507883099?color=768AD4&label=discord)](https://discord.com/channels/753398645507883099/)
[![Twitter](https://img.shields.io/twitter/follow/sarcophagusio?style=social)](https://twitter.com/sarcophagusio)

Sarcophagus is a decentralized dead man's switch built on Ethereum and Arweave.

## Overview

This repository contains the Go service that enables a user to participate in the Sarcophagus system as an Archaeologist. Archaeologists are responsible for executing Sarcophagus jobs created via the [https://github.com/sarcophagus-org/sarcophagus-contracts](Sarcophagus contracts).

## Prerequisites

Before you can start accepting jobs, you must have the following:

###### ARWEAVE tokens
- You will need Arweave tokens to accept new jobs. 
- You can get a wallet [here].(https://www.arweave.org/wallet)
- Your wallet key file will need to be accessible to the service (further instructions provided below).

###### SARCO Tokens
- You must have a SARCO token balance to accept new jobs. 
- You can get SARCO on [Uniswap](https://app.uniswap.org/#/swap?inputCurrency=0x7697b462a7c4ff5f8b55bdbc2f4076c2af9cf51a)
- The address that holds these tokens is the one derived from the `eth_private_key` value in the config file.

###### ETH
- The same address that holds your SARCO tokens must have an Eth balance.
- This ETH will be used to create transactions necessary for the Archaeologist service to function (Registering Archaeologists, Unwrapping Sarcophogi, etc)

## Quick Start

### Build

First, clone the repo:

```sh
$ git clone <https://github.com/sarcophagus-org/archaeologist-service>
```

Then, initialize the configuration file:

```sh
$ cp config.example.yml config.yml
```

Finally, build the service:

```sh
$ go build
```

## Configuration

Update the config.yml file with the appropriate values. 
- Use the comments in the config.example.yml file as a guide. 
- **The config.yml must be writable.**
- All comments in config.yml will be removed after the first time the service is run.
- Put your arweave wallet file in a path on the filesystem, and update the config.yml `areave_key_file` value to point to the location of your arweave wallet. (i.e.`/usr/local/arweave.json`)
- You will need to use a domain for your endpoint to accommodate SSL. This domain name will be used as the `endpoint` config file value (i.e. `https://arch1.myarch.com`)
- Update your domain's DNS to point at the IP address of the server running the Archaeologist service.
- Expose the IP address on port 443 and map to the "file_port" config value (default is 8080). One option for this is to use an [nginx reverse proxy](https://docs.nginx.com/nginx/admin-guide/web-server/reverse-proxy/). 

## Example

An nginx proxy using letsencrypt for the SSL cert and `arch1.myarch.io` domain:

- Open Port 443 on your server
- [Install nginx](https://www.nginx.com/resources/wiki/start/topics/tutorials/install/)
- [Use certbot to install SSL](https://certbot.eff.org/) 
- Create the reverse proxy file `/etc/nginx/sites-available/reverse-proxy.conf`

```nginx
server {
        # SSL Setup
        listen 443 ssl;
        server_name arch1.myarch.com;
        ssl_certificate /etc/letsencrypt/live/myarch.com/fullchain.pem;
        ssl_certificate_key /etc/letsencrypt/live/myarch.com/privkey.pem;
        ssl_stapling on;
        ssl_stapling_verify on;

        # End SSL Setup

        access_log /var/log/nginx/reverse-access.log;
        error_log /var/log/nginx/reverse-error.log;

        location / {
            proxy_pass http://127.0.0.1:8080;
        }
}
```

- Create a sym link to sites-enabled

```sh
$ sudo ln -s /etc/nginx/sites-available/reverse-proxy.conf /etc/nginx/sites-enabled/reverse-proxy.conf
```

- Restart nginx

```sh
$ sudo service nginx restart
```

## Run

To run the service:

```sh
$ ./archaeologist-service
```

## Install Service (optional)

**Alternatively you can install the service globally with:**

```sh
$ go install
```

Copy the config file to $GOPATH/bin (assuming you have `$GOPATH/bin` in your $PATH)

```sh
$ cp config.example.yml config.yml
$ cp config.yml $GOPATH/bin
```

Update the config.yml `areave_key_file` value to point to the location of your arweave wallet.

If installing on Ubuntu, before building/installing, you may need to run:

```sh
$ apt-get install build-essential
```

## Local Development

### Deploy Sarcophagus Contract
Clone https://github.com/sarcophagus-org/sarcophagus-contracts

Follow Readme directions to deploy the contract locally.

### Fill out values in config file
See `config/config.example.yml` for descriptions and examples for each config key/value.

##### Run Arweave Blockchain

```sh
$ docker pull rootmos/loom
$ docker run --rm --publish 8000:8000 rootmos/loom
```

### Create Arweave Key
Install arweave-deploy
https://docs.arweave.org/developers/tools/arweave-deploy

Generate arweave key:

```sh
$ arweave key-create your-arweave-key.json
```

1. Copy the generated key to /config directory.
2. Update the `config/config.yml` file `AREWAVE_KEY_FILE` to be the name of this file.

### Add tokens from faucet to your wallet address

```sh
$ curl -d '{"beneficiary": "<arweave wallet address>", "quantity": 1000000000000}' http://localhost:8000/loom/faucet
```

See https://github.com/rootmos/loom for more information

### Redeploy Contract

If the Sarcophagus contract gets updated, you need to re-compile the contract for the service to use.

Install ethereum and solidity. Below are instructions for homebrew. See this link for other methods:
https://solidity.readthedocs.io/en/v0.5.3/installing-solidity.html

```sh
$ brew update
$ brew upgrade
$ brew tap ethereum/ethereum
$ brew install solidity
```

Install abigen

```sh
$ export GOPATH=$HOME/go
$ export PATH=$PATH:$GOPATH/bin
$ make
$ make devtools
```

Use abigen to compile abi -- below examples output to "abi" & "abiToken" directories. Alternatively use abi in Sarcophagus Contracts build folder.
NOTE: The solidity compiler version must match the version specified in the contract (currently ^0.6.0)

```sh
$ solc --abi @openzeppelin/=$(pwd)/node_modules/@openzeppelin/ contracts/Sarcophagus.sol -o abi
$ solc --abi @openzeppelin/=$(pwd)/node_modules/@openzeppelin/ contracts/mocks/SarcoTokenMock.sol -o abiToken --allow-paths $(pwd)
$ solc --abi @openzeppelin/=$(pwd)/node_modules/@openzeppelin/ contracts/libraries/Events.sol -o abiEvents
```

Compile Contracts to Go

```sh
$ abigen --abi=./abi/Sarcophagus.abi --pkg=sarcophagus --out=Sarcophagus.go
$ abigen --abi=./abiToken/SarcoTokenMock.abi --pkg=token --out=SarcophagusToken.go
$ abigen --abi=./abiEvents/Events.abi --pkg=events --out=Events.go
```

1. Copy/Replace generated files into /contracts directory
2. Rename the package name on these files to "contracts"

If you have any issues with the compiled go code, you may need to download an older version of abigen.

https://geth.ethereum.org/downloads/

The latest tested version working with the service is Geth & Tools 1.9.25

## Testing

An embalmer package is provided for local testing. 
Embalmer config values can be updated in `embalmer/embalmer_config.yml`
These config values are configured to work with the main service's test config at `test/test_config.yml`

Examples are below.

```sh
# Start the arch service. Config values must be set correctly and free bond must be added to accept new jobs.
$ go run main.go

# The embalmer must have a sufficient Sarco token balance.

# The seed flag is used to generate file bytes, which will be used as the asset file and generate the asset double hash for a Sarcophagus.
# This seed can be changed to create/modify different sarcophaguses. 

# Create a Sarcophagus
# Exluding the -type flag will set the type as "create" by defaut.
$ go run cmd/embalmer.go -seed=200

# Update a Sarcophagus
$ go run cmd/embalmer.go -seed=200 -type=update

# Rewrap a Sarcophagus
$ go run cmd/embalmer.go -seed=200 -type=rewrap

# Clean a Sarcophagus
$ go run cmd/embalmer.go -seed=200 -type=clean

# Bury a Sarcophagus
$ go run cmd/embalmer.go -seed=200 -type=bury

# Cancel a Sarcophagus
$ go run cmd/embalmer.go -seed=200 -type=cancel
```

There is a test suite provided that uses the embalmer package, see the README.md in the test directory for more directions.

## Sending a file

If you want to test sending a file locally (the Sarcophagus payload) after creating a sarcophagus: 

```sh
$ curl -v -X POST -F file=@<your file> -F http://127.0.0.1:<your port>/file
```

## Community
[![Discord](https://img.shields.io/discord/753398645507883099?color=768AD4&label=discord)](https://discord.com/channels/753398645507883099/)
[![Twitter](https://img.shields.io/twitter/follow/sarcophagusio?style=social)](https://twitter.com/sarcophagusio)

We can also be found on [Telegram](https://t.me/sarcophagusio).

Made with :skull: and proudly decentralized.
