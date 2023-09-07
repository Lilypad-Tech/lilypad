package web3

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestGetPublicKey(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	publicKey := GetPublicKey(privateKey)
	assert.Equal(t, privateKey.PublicKey, publicKey, "Public keys should be equal")
}

func TestSignMessage(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	message := []byte("Hello, world!")
	sig, err := SignMessage(privateKey, message)
	if err != nil {
		t.Fatalf("Failed to sign message: %v", err)
	}

	assert.NotEmpty(t, sig, "Signature should not be empty")
}

func TestGetAddressFromSignedMessage(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	message := []byte("Hello, world!")
	sig, err := SignMessage(privateKey, message)
	if err != nil {
		t.Fatalf("Failed to sign message: %v", err)
	}

	address, err := GetAddressFromSignedMessage(message, sig)
	if err != nil {
		t.Fatalf("Failed to get address from signed message: %v", err)
	}

	expectedAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	assert.Equal(t, expectedAddress, address, "Addresses should be equal")
}
