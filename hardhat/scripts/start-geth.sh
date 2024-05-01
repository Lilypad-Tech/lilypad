#!/bin/bash

# Start geth
rm -rf .dev-chain
geth --datadir .dev-chain init infra/geth/genesis.json
geth --datadir .dev-chain account import -password infra/geth/password infra/geth/keys/admin
geth --datadir .dev-chain account import -password infra/geth/password infra/geth/keys/alice
geth --datadir .dev-chain account import -password infra/geth/password infra/geth/keys/bob
geth --datadir .dev-chain account import -password infra/geth/password infra/geth/keys/charlie
geth --datadir .dev-chain account import -password infra/geth/password infra/geth/keys/dave
geth --datadir .dev-chain account import -password infra/geth/password infra/geth/keys/eve
geth --datadir .dev-chain account import -password infra/geth/password infra/geth/keys/frank
geth --datadir .dev-chain --dev --http --http.api personal,eth,net,web3,debug --http.corsdomain "*" --http.addr "localhost" --http.port 8545 --allow-insecure-unlock
