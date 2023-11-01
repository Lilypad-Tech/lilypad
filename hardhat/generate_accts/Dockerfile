# Start with the official Golang image
FROM golang:1.21.3

WORKDIR /geth
ARG VERSION=v1.13.4
RUN git clone --quiet --branch ${VERSION} --depth 1 https://github.com/ethereum/go-ethereum .
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build go run build/ci.go install -static ./cmd/geth

RUN /geth/build/bin/geth version
RUN mv /geth/build/bin/geth /usr/local/bin/

# Set the working directory
WORKDIR /app

# Copy the Go file
COPY ./extractKeys.go .

# Initialize a Go module and get the required packages
RUN go mod init extractkeys
RUN go get github.com/ethereum/go-ethereum/accounts/keystore@${VERSION}
RUN go get github.com/ethereum/go-ethereum/crypto@${VERSION}

# Build the Go application
RUN go build -o extractkeys extractKeys.go

# Create a volume to persist Ethereum keystore data
VOLUME ["/root/.ethereum/keystore"]

# Use bash as the default entrypoint
ENTRYPOINT ["/bin/bash"]