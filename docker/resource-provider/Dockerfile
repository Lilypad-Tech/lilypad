ARG COMPUTE_MODE=gpu

FROM golang:1.22.4 AS base
WORKDIR /usr/src/app
ARG NETWORK=testnet
ARG VERSION
ARG COMMIT_SHA

# Default environment variables
ENV LOG_LEVEL=info
ENV BACALHAU_API_HOST="localhost"
ENV WEB3_PRIVATE_KEY=""
ENV DISABLE_TELEMETRY=false

COPY . .

FROM nvidia/cuda:12.0.1-cudnn8-devel-ubuntu22.04  AS build-gpu
WORKDIR /usr/src/app
COPY --from=base /usr/src/app .
COPY --from=base /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"
RUN nvcc --version && nvcc --ptx -o ./pkg/resourceprovider/cudaminer/keccak.ptx ./pkg/resourceprovider/cudaminer/keccak.cu
RUN go build -v -tags cuda -ldflags="-X 'github.com/lilypad-tech/lilypad/pkg/system.Version=${VERSION}' -X 'github.com/lilypad-tech/lilypad/pkg/system.CommitSHA=${COMMIT_SHA}'"
ENV DISABLE_POW=false

FROM base AS build-cpu
RUN go build -v -ldflags="-X 'github.com/lilypad-tech/lilypad/pkg/system.Version=${VERSION}' -X 'github.com/lilypad-tech/lilypad/pkg/system.CommitSHA=${COMMIT_SHA}'"
ENV DISABLE_POW=true

FROM build-$COMPUTE_MODE AS final
RUN mv lilypad /usr/local/bin

# Install necessary dependencies
RUN apt update && apt install -y wget bash curl && apt clean

# Install Bacalhau
RUN cd /tmp && \
    wget https://github.com/bacalhau-project/bacalhau/releases/download/v1.3.2/bacalhau_v1.3.2_linux_amd64.tar.gz && \
    tar xfv bacalhau_v1.3.2_linux_amd64.tar.gz && \
    mv bacalhau /usr/local/bin/bacalhau && \
    rm bacalhau_v1.3.2_linux_amd64.tar.gz

# Add both lilypad and bacalhau executables to PATH
ENV PATH="/usr/local/bin:${PATH}"

# Create a startup script to run both services simultaneously
RUN touch run
RUN echo "#!/bin/bash" >> run

# Ensure bacalhau is initialized
RUN echo "export BACALHAU_ENVIRONMENT=local" >> run
RUN echo "bacalhau id" >> run

# Launch Lilypad
RUN echo "/usr/local/bin/lilypad resource-provider --network ${NETWORK} --disable-pow=${DISABLE_POW} --disable-telemetry=${DISABLE_TELEMETRY} &" >> run
RUN echo "wait -n" >> run
RUN chmod +x run

# Run startup script when container starts
CMD ["/bin/bash", "./run"]