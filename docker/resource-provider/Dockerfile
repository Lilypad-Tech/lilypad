ARG COMPUTE_MODE=gpu

FROM golang:latest AS base
WORKDIR /usr/src/app
ARG VERSION
ARG COMMIT_SHA

COPY . .

FROM nvidia/cuda:12.0.1-cudnn8-devel-ubuntu22.04 AS build-gpu
WORKDIR /usr/src/app
ARG VERSION
ARG COMMIT_SHA

COPY --from=base /usr/src/app .
COPY --from=base /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"
RUN nvcc --version && nvcc --ptx -o ./pkg/resourceprovider/cudaminer/keccak.ptx ./pkg/resourceprovider/cudaminer/keccak.cu
RUN go build -v -tags cuda -ldflags="-X 'github.com/Lilypad-Tech/lilypad/v2/pkg/system.Version=${VERSION}' -X 'github.com/Lilypad-Tech/lilypad/v2/pkg/system.CommitSHA=${COMMIT_SHA}'"
ENV DISABLE_POW=false

FROM base AS build-cpu
ARG VERSION
ARG COMMIT_SHA

RUN go build -v -ldflags="-X 'github.com/Lilypad-Tech/lilypad/v2/pkg/system.Version=${VERSION}' -X 'github.com/Lilypad-Tech/lilypad/v2/pkg/system.CommitSHA=${COMMIT_SHA}'"
ENV DISABLE_POW=true

FROM build-$COMPUTE_MODE AS final
# Default environment variables
ARG NETWORK=testnet
ARG DISABLE_TELEMETRY=false
ENV LOG_LEVEL=info
ENV BACALHAU_API_HOST="localhost"
ENV WEB3_PRIVATE_KEY=""

RUN mv lilypad /usr/local/bin
# Install necessary dependencies
RUN apt update && apt install -y wget bash curl && apt clean

# Add both lilypad and bacalhau executables to PATH
ENV PATH="/usr/local/bin:${PATH}"

# Create a startup script to run both services simultaneously
RUN touch run
RUN echo "#!/bin/bash" >> run

# Launch Lilypad
RUN echo "/usr/local/bin/lilypad resource-provider --network ${NETWORK} --disable-pow=${DISABLE_POW} --disable-telemetry=${DISABLE_TELEMETRY} &" >> run
RUN echo "wait -n" >> run
RUN chmod +x run

# Run startup script when container starts
CMD ["/bin/bash", "./run"]
