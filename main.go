package main

import (
	"fmt"

	"github.com/bacalhau-project/lilypad/cmd/lilypad"
)

func main() {
	lilypad.Execute()
}

var VERSION string

func init() {
	fmt.Println("Version main:", VERSION)
}
