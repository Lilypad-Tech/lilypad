package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// Check that the password was provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: extractkeys <password>")
		return
	}

	keystorePath := "/root/.ethereum/keystore/"
	password := os.Args[1]

	// Get the list of files in the keystore directory
	files, e := os.ReadDir(keystorePath)
	if e != nil {
		panic(e)
	}
	for _, file := range files {
		// Extract the private key from each file
		keyjson, e := os.ReadFile(keystorePath + file.Name())
		if e != nil {
			panic(e)
		}
		key, e := keystore.DecryptKey(keyjson, password)
		if e != nil {
			panic(e)
		}

		// Convert the private key to its hexadecimal representation
		privateKeyHex := hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))
		fmt.Println(file.Name(), privateKeyHex)
	}
}
