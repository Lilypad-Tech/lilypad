#!/bin/sh

EXPOSE_VIA=${EXPOSE_VIA:-"cloudflare"}
NETWORK=${NETWORK:-"testnet"}

if [ "$NETWORK" = "dev" ]; then 
  /usr/local/bin/lilypad solver --network ${NETWORK} "$@"
else 
  if [ "$EXPOSE_VIA" = "local" ]; then
      doppler run -- /usr/local/bin/lilypad solver --network ${NETWORK} "$@"
  else
      doppler run --command "cloudflared tunnel run & /usr/local/bin/lilypad solver --network ${network}"
  fi
fi