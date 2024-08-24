ARG expose_via=local
ARG arch=amd64

FROM docker:24.0.5-dind AS base
ARG admin_address="0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"


RUN apk update
RUN apk add bash
RUN apk add nodejs
RUN apk add npm

WORKDIR /l2chain
COPY . .
RUN touch run
RUN echo "#!/bin/bash" >> run
RUN echo "dockerd-entrypoint.sh & " >> run
RUN echo "until docker info >/dev/null 2>&1; do " >> run
RUN echo "   echo 'Waiting for Docker to start...'" >> run
RUN echo "   sleep 1" >> run
RUN echo "done" >> run

RUN echo "docker pull offchainlabs/nitro-node:v2.3.4-b4cc111 " >> run

RUN echo "docker run -d --name devnet -p 0.0.0.0:8547:8547 -p 0.0.0.0:8548:8548 offchainlabs/nitro-node:v2.3.4-b4cc111  "\
  "   --init.dev-init " \
  "   --init.dev-init-address $admin_address "\
  "   --node.dangerous.no-l1-listener "\
  "   --node.parent-chain-reader.enable=false "\
  "   --parent-chain.id=1337 "\
  "   --chain.id=412346 "\
  "   --persistent.chain /tmp/dev-test "\
  "   --node.sequencer "\
  "   --execution.sequencer.enable "\
  "   --node.dangerous.no-sequencer-coordinator "\
  "   --node.staker.enable=false "\
  "   --init.empty=false "\
  "   --http.port 8547  "\
  "   --http.addr 0.0.0.0 "\
  "   --ws.port 8548  "\
  "   --ws.addr 0.0.0.0 "\
  "   --http.corsdomain=* "\
  "   --http.vhosts=* "\
  "   --ws.origins=* "\
  "    " >> run

RUN echo "until wget -qO- http://0.0.0.0:8547 >/dev/null 2>&1; do " >> run
RUN echo "   echo 'Waiting for Nitro to start...'" >> run
RUN echo "   sleep 1" >> run
RUN echo "done" >> run

RUN echo "cd hardhat" >> run
RUN echo "npm install" >> run
RUN echo "npx hardhat deploy --network devnet" >> run
RUN echo "npx hardhat run scripts/fund-services-ether.ts  --network devnet" >> run
RUN echo "npx hardhat run scripts/fund-services-tokens.ts  --network devnet " >> run
RUN echo "npx hardhat run scripts/print-l2-config.ts  --network devnet" >> run
RUN echo "cd .." >> run
RUN echo "./stack go-bindings" >> run



FROM base AS expose-cloudflare
ARG cloudflare_token_http="not-a-token"
ARG cloudflare_token_ws="not-a-token"

RUN wget https://github.com/cloudflare/cloudflared/releases/latest/download/cloudflared-linux-amd64
RUN echo "./cloudflared-linux-amd64 tunnel run --token $cloudflare_token_http  &" >> run
RUN echo "./cloudflared-linux-amd64 tunnel run --token $cloudflare_token_ws &" >> run
RUN chmod +x cloudflared-linux-amd64

FROM base AS expose-local

FROM expose-$expose_via AS final
RUN chmod +x run
RUN echo "wait -n" >> run
ENTRYPOINT ["bash", "./run"]  