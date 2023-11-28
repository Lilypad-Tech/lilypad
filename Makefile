release:
	#go build -v -ldflags="-X lilypad.VERSION=$(git describe --tags --abbrev=0) -X lilypad.COMMIT_SHA=$(git rev-parse HEAD)" .
	#go build -v -ldflags="-X lilypad.VERSION=x -X lilypad.COMMIT_SHA=x.xx." .
	go build -v -ldflags="-X 'lilypad/cmd/lilypad/lilypad.VERSION=v2.0.0'" .
	./lilypad version