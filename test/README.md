# Tests

### Overview
The `archaeologist_test.go` file contains some end-to-end tests.

The testing suite uses bash scripts to locally deploy:
1. The contracts + eth blockchain via https://github.com/sarcophagus-org/smart-contracts
2. The arweave blockchain

See main README.md for information about prerequisites for deploying these.
For the sarcophagus contracts, these tests assume `npm install` & `npm run build` have been run.

### Setup
Update the values in: `/test/test_setup.yml`.

The tests load the `test/test_config.yml`. 
The values in this file are pre-configured to work with the defaults in the Sarcophagus contracts. 
You can update these values if you want to use a different arch private key, etc.

`test/arweave.json` - This arweave wallet can be replaced if you want to use a different local arweave wallet for testing.

