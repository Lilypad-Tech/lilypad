package main

import (
	"fmt"

	"github.com/bacalhau-project/lilypad/cmd/lilypad"
)

func main() {
	lilypad.Execute()
}

func init() {
	fmt.Printf("Lilypad: %s\n", lilypad.VERSION)
}
