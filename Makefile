release:
	#go build -v -ldflags="-X lilypad.VERSION=$(git describe --tags --abbrev=0) -X lilypad.COMMIT_SHA=$(git rev-parse HEAD)" .
	#go build -v -ldflags="-X lilypad.VERSION=x -X lilypad.COMMIT_SHA=x.xx." .
	#go build -v -ldflags="-X lilypad.VERSION=x" .
	go build -v -ldflags="-X main.VERSION=x" .
	./lilypad version