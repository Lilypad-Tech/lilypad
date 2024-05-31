release:
	go build -v -ldflags="\
		-X 'github.com/lilypad-tech/lilypad/v2/cmd/lilypad.VERSION=$$(git describe --tags --abbrev=0)' \
		-X 'github.com/lilypad-tech/lilypad/v2/cmd/lilypad.COMMIT_SHA=$$(git rev-parse HEAD)' \
	" -o bin/
	./bin/lilypad version



.PHONY: release install

install:
	make release
	cp ./bin/lilypad.exe $$GOBIN
#Ps1: cmd	cp ./bin/lilypad.exe $env:GOBIN
