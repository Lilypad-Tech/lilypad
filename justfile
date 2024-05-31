default: go build 

release:
	go build -v -ldflags="\
		-X 'github.com/lilypad-tech/lilypad/v2/cmd/lilypad.VERSION=$$(git describe --tags --abbrev=0)' \
		-X 'github.com/lilypad-tech/lilypad/v2/cmd/lilypad.COMMIT_SHA=$$(git rev-parse HEAD)' \
	" .
	./lilypad version