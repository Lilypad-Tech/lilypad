FROM ethereum/client-go:v1.11.6

WORKDIR /app
COPY ./genesis.json /app/genesis.json
COPY ./keys /app/keys
COPY ./password /app/password

RUN geth --datadir .dev-chain init /app/genesis.json
RUN geth --datadir .dev-chain account import -password /app/password /app/keys/admin
RUN geth --datadir .dev-chain account import -password /app/password /app/keys/alice
RUN geth --datadir .dev-chain account import -password /app/password /app/keys/bob
RUN geth --datadir .dev-chain account import -password /app/password /app/keys/charlie
RUN geth --datadir .dev-chain account import -password /app/password /app/keys/dave
RUN geth --datadir .dev-chain account import -password /app/password /app/keys/eve
RUN geth --datadir .dev-chain account import -password /app/password /app/keys/frank

ENTRYPOINT geth --datadir .dev-chain --dev --http --http.api personal,eth,net,web3,debug --allow-insecure-unlock --http.addr "0.0.0.0"