ARG expose_via=local

FROM golang:1.22.3 AS base

ARG arch=amd64
ARG cloudflare_token_http="not-a-token"
ARG cloudflare_token_ws="not-a-token"
ARG geth_version=v1.13.5

WORKDIR /geth

RUN git clone --quiet --branch ${geth_version} --depth 1 https://github.com/ethereum/go-ethereum .
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build go run build/ci.go install -static ./cmd/geth

RUN /geth/build/bin/geth version
RUN mv /geth/build/bin/geth /usr/local/bin/

ARG admin_address="0x0"

RUN touch fund-admin
RUN echo "#!/bin/bash" >> fund-admin
RUN echo "geth --exec \"eth.sendTransaction({from: eth.coinbase, to: \\\"${admin_address}\\\", value: new web3.BigNumber(eth.getBalance(eth.coinbase)).minus(web3.toWei(1, \\\"ether\\\"))  })\" attach /data/chain/geth.ipc" >> fund-admin
RUN chmod +x fund-admin

RUN touch run-node
RUN echo "#!/bin/bash" >> run-node
RUN echo "geth --datadir /data/chain --dev --ws --ws.api web3,eth,net --ws.addr 0.0.0.0 --ws.port 8546 --ws.origins '*' --http --http.api web3,eth,net --http.addr 0.0.0.0 --http.corsdomain '*' --http.port 8545 --http.vhosts '*' &" >> run-node
RUN chmod +x run-node

RUN touch reset
RUN echo "#!/bin/bash" >> reset
RUN echo "echo '- Kill geth'" >> reset
RUN echo "pkill -INT geth" >> reset
RUN echo "sleep 5" >> reset
RUN echo "echo '- Clear data'" >> reset
RUN echo "rm -rf /data/chain/*" >> reset
RUN echo "echo '- Restart geth'" >> reset
RUN echo "./run-node" >> reset
RUN echo "sleep 5" >> reset
RUN echo "echo '- Fund admin'" >> reset
RUN echo "./fund-admin" >> reset
RUN echo "echo '- Done'" >> reset
RUN chmod +x reset

RUN touch run
RUN echo "#!/bin/bash" >> run

FROM base AS expose-cloudflare
RUN curl -L --output cloudflared.deb https://github.com/cloudflare/cloudflared/releases/latest/download/cloudflared-linux-${arch}.deb
RUN dpkg -i cloudflared.deb
RUN echo "cloudflared tunnel --metrics 0.0.0.0:11111 run --token $cloudflare_token_http --url http://localhost:8545 &" >> run
RUN echo "cloudflared tunnel --metrics 0.0.0.0:11112 run --token $cloudflare_token_ws --url http://localhost:8546 &" >> run

FROM base AS expose-local
EXPOSE 8545
EXPOSE 8546

FROM expose-$expose_via AS final
RUN echo "./run-node" >> run
RUN echo "sleep infinity" >> run
RUN chmod +x run

CMD ["/bin/bash", "./run"]
