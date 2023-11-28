//go:build generate

package main

//go:generate make release
//go:generate ./lilypad -v

/*
go:generate echo "go build -v -ldflags=\"-X lilypad.VERSION=$(git describe --tags --abbrev=0) -X lilypad.COMMIT_SHA=$(git rev-parse HEAD)\" ." | tee /dev/tty | sh
go:generate go build -ldflags="-X lilypad.VERSION=$(git describe --tags --abbrev=0) -X lilypad.COMMIT_SHA=$(git rev-parse HEAD)"

//go:generate echo 'go build -v -ldflags=\"-X lilypad.VERSION=$(git describe --tags --abbrev=0) -X lilypad.COMMIT_SHA=$(git rev-parse HEAD)\" .' | tee | sh


FIXME:: not working cuz of issue: https://stackoverflow.com/questions/28558110/go-generate-escape-characters
*/
