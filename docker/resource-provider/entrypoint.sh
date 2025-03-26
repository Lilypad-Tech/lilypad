#!/bin/sh

NETWORK=${NETWORK:-"testnet"}

/usr/local/bin/lilypad resource-provider --network ${NETWORK}  "$@"