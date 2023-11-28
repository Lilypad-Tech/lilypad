//go:build ignore
package main

go:generate echo \"go build -ldflags="-X lilypad.VERSION=$(git describe --tags --abbrev=0) -X lilypad.COMMIT_SHA=$(git rev-parse HEAD)"\" | exec $0

go:generate go build -ldflags="-X lilypad.VERSION=$(git describe --tags --abbrev=0) -X lilypad.COMMIT_SHA=$(git rev-parse HEAD)"
