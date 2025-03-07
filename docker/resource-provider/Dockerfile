ARG COMPUTE_MODE=gpu

FROM golang:latest AS base
WORKDIR /usr/src/app
ARG VERSION
ARG COMMIT_SHA

COPY . .

FROM nvidia/cuda:12.0.1-cudnn8-devel-ubuntu22.04 AS build-gpu
WORKDIR /usr/src/app
COPY --from=base /usr/src/app .
COPY --from=base /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"
RUN nvcc --version && nvcc --ptx -o ./pkg/resourceprovider/cudaminer/keccak.ptx ./pkg/resourceprovider/cudaminer/keccak.cu
RUN go build -v -tags cuda -ldflags="-X 'github.com/lilypad-tech/lilypad/pkg/system.Version=${VERSION}' -X 'github.com/lilypad-tech/lilypad/pkg/system.CommitSHA=${COMMIT_SHA}'"

FROM base AS build-cpu
RUN go build -v -ldflags="-X 'github.com/lilypad-tech/lilypad/pkg/system.Version=${VERSION}' -X 'github.com/lilypad-tech/lilypad/pkg/system.CommitSHA=${COMMIT_SHA}'"

FROM build-$COMPUTE_MODE AS final

RUN mv lilypad /usr/local/bin
# Install necessary dependencies
RUN apt update && apt install -y wget bash curl && apt clean

# Add both lilypad and bacalhau executables to PATH
ENV PATH="/usr/local/bin:${PATH}"
ENV NETWORK=testnet

COPY ./docker/resource-provider/entrypoint.sh /entrypoint.sh

ENTRYPOINT [ "/entrypoint.sh" ]