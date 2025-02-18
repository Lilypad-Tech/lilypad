# GPU Tools

## Development

### Setup

1. For testing, copy the sample environment file:
```bash
cp tests/.env.sample tests/.env
```

2. The tests require a valid JWT key for Kafka authentication. Run: 
```bash
cd cmd/generate_token
go run main.go -subject "rp_gputooltest" -duration 1y -kid dev > rp_gputooltest_token.txt
```

Format:
```
JWT_TESTRP_TOKEN=eyJhbGciOiJIUzI1NiIsImtpZCI6ImtleS1kZXYiLCJ0eXAiOiJKV1QifQ...
```

### Running

#### Using Dev Container

1. Start the service in dev container terminal:
```bash
cargo run
```

2. In a second dev container terminal, run tests:
```bash
cargo test -- --test-threads=1 --nocapture
cargo test --test grpc_test -- --test-threads=1 --nocapture
```

#### Using Docker Compose

1. Start the service:
```bash
docker compose -f docker/docker-compose.base.yml up gpu-tools
```

2. In another terminal, run tests:
```bash
docker exec gpu-tools cargo test --test grpc_test -- --test-threads=1 --nocapture
```

### Regenerate Go files if proto definition has changed

1. First install the protoc compiler and Go plugins (only needed once):
```bash
# Install protoc compiler (Ubuntu/Debian)
sudo apt-get install -y protobuf-compiler
# Or for Mac
brew install protobuf

# Install Go plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Add Go bin directory to PATH (add this to your ~/.bashrc or ~/.zshrc)
export PATH="$PATH:$(go env GOPATH)/bin"
# Then reload your shell or run:
source ~/.bashrc  # or source ~/.zshrc
```

2. From the root `lilypad` directory, generate Go files from proto:
```bash

protoc \
  --proto_path=pkg/gpu-tools/proto \
  --go_out=pkg/resourceprovider/gpu-tools-protobuf \
  --go_opt=paths=source_relative \
  --go-grpc_out=pkg/resourceprovider/gpu-tools-protobuf \
  --go-grpc_opt=paths=source_relative \
  pkg/gpu-tools/proto/gpu_tools.proto
```

The generated files will be created in `pkg/resourceprovider/gpu-tools-protobuf/`.




