//go:build generate

package main

//go:generate echo "go build -ldflags=\"-X lilypad.VERSION=$(git describe --tags --abbrev=0) -X lilypad.COMMIT_SHA=$(git rev-parse HEAD)\" ." | tee /dev/tty | sh


//go:generate ./lilypad -v


go:generate go build -ldflags="-X lilypad.VERSION=$(git describe --tags --abbrev=0) -X lilypad.COMMIT_SHA=$(git rev-parse HEAD)"
