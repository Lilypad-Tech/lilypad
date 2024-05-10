FROM docker:dind

ADD https://github.com/bacalhau-project/bacalhau/releases/download/v1.3.0/bacalhau_v1.3.0_linux_amd64.tar.gz .
RUN tar xfv bacalhau_v1.3.0_linux_amd64.tar.gz
RUN mv bacalhau /usr/local/bin

COPY ./do .
#COPY ./stack .
#COPY ./docker/bacalhau/Dockerfile .

CMD ["/bin/sh", "./do"]

# TODO
# contents of do:
# # wait for docker to initialize
# while ! docker -v; do sleep 1; done

# export BACALHAU_SERVE_IPFS_PATH=/tmp/lilypad/data/ipfs
# TODO
# check if this is needed
# export BACALHAU_API_HOST=localhost
# export LOG_LEVEL=debug

# bacalhau serve --node-type compute,requester --peer none --private-internal-ipfs=false --job-selection-accept-networked --ipfs-api-listen-addresses=/ip4/0.0.0.0/tcp/0
