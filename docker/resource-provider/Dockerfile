FROM docker:dind
WORKDIR /usr/src/app

# Build args
ARG NETWORK=testnet
ARG DISABLE_POW=false

# Default environment variables
ENV LOG_LEVEL=info
ENV OFFER_GPU=1
ENV BACALHAU_API_HOST="localhost"
ENV WEB3_PRIVATE_KEY=""
ENV DISABLE_TELEMETRY=false

# Install necessary dependencies
RUN apk update
RUN apk add wget
RUN apk add bash

# Install Bacalhau
RUN cd /tmp && \
    wget https://github.com/bacalhau-project/bacalhau/releases/download/v1.3.2/bacalhau_v1.3.2_linux_amd64.tar.gz && \
    tar xfv bacalhau_v1.3.2_linux_amd64.tar.gz && \
    mv bacalhau /usr/local/bin/bacalhau && \
    rm bacalhau_v1.3.2_linux_amd64.tar.gz

# Build and install Lilypad
COPY --from=golang:1.22.4-alpine /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"

COPY . .
RUN go build -v .
RUN mv lilypad /usr/local/bin

# Add both lilypad and bacalhau executables to PATH
ENV PATH="/usr/local/bin:${PATH}"

# Create a startup script to run both services simultaneously
RUN touch run
RUN echo "#!/bin/bash" >> run
# Launch Bacalhau
RUN echo "/usr/local/bin/bacalhau serve --node-type compute,requester --peer none --private-internal-ipfs=false --job-selection-accept-networked &" >> run
# Launch Lilypad
RUN echo "/usr/local/bin/lilypad resource-provider --network ${NETWORK} --disable-pow=${DISABLE_POW} --disable-telemetry=${DISABLE_TELEMETRY} &" >> run
RUN echo "wait -n" >> run
RUN chmod +x run

# Run startup script when container starts
CMD ["/bin/bash", "./run"]
