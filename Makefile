release:
	go build -v -ldflags="\
		-X 'github.com/bacalhau-project/lilypad/cmd/lilypad.VERSION=$$(git describe --tags --abbrev=0)' \
		-X 'github.com/bacalhau-project/lilypad/cmd/lilypad.COMMIT_SHA=$$(git rev-parse HEAD)' \
	" .
	./lilypad version

install:
	go install -v -ldflags="\
		-X 'github.com/bacalhau-project/lilypad/cmd/lilypad.VERSION=$$(git describe --tags --abbrev=0)' \
		-X 'github.com/bacalhau-project/lilypad/cmd/lilypad.COMMIT_SHA=$$(git rev-parse HEAD)' \
	" .
	lilypad version

.PHONY: release install

