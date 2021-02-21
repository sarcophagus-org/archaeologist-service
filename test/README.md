# Tests

### Overview
The `archaeologist_test.go` file contains some end-to-end tests to test the entire workflow of a sarcophagus
being created, updated, unwrapped and rewrapped. 

The testing suite uses bash scripts to deploy and stop the local arweave client - https://github.com/rootmos/loom

### Setup
1. Run the Sarcophagus contracts locally: 
- https://github.com/sarcophagus-org/sarcophagus-contracts
- See README in that repo for instructions.

2. Update the values in: `/test/test_setup.yml` to match your local ethereum client configuration

3. Update The tests config file `test/test_config.yml`.
- This is pre-configured to work with the default contract & token addresses in the sarcophagus contracts. 
- The values in this file are pre-configured to work with the values in the embalmer config file. 
-   You can update these values if you want to use a different arch private key, etc.

4. `test/arweave.json` - This arweave wallet can be replaced if you want to use a different local arweave wallet for testing.

