package web3

import (
	"crypto/ecdsa"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetPublicKey(privateKey *ecdsa.PrivateKey) ecdsa.PublicKey {
	return privateKey.PublicKey
}

func GetAddress(privateKey *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}

func SignMessage(privateKey *ecdsa.PrivateKey, message []byte) ([]byte, error) {
	hash := crypto.Keccak256Hash(message)
	return crypto.Sign(hash.Bytes(), privateKey)
}

func GetAddressFromSignedMessage(message []byte, sig []byte) (common.Address, error) {
	hash := crypto.Keccak256Hash(message)
	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), sig)
	if err != nil {
		return common.Address{}, err
	}
	recoveredPubKey, err := crypto.UnmarshalPubkey(sigPublicKey)
	if err != nil {
		return common.Address{}, err
	}
	return crypto.PubkeyToAddress(*recoveredPubKey), nil
}

func ParsePrivateKey(privateKey string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(strings.Replace(privateKey, "0x", "", 1))
}

func EtherToWei(etherAmount float64) *big.Int {
	ether := new(big.Float).SetFloat64(etherAmount)
	weiMultiplier := new(big.Float).SetFloat64(1e18)
	wei := new(big.Float).Mul(ether, weiMultiplier)

	weiInt := new(big.Int)
	wei.Int(weiInt)

	return weiInt
}

func EtherToWeiUint64(etherAmount float64) uint64 {
	wei := EtherToWei(etherAmount)
	return wei.Uint64()
}

func ConvertStringToBigInt(st string) big.Int {
	bigInt, _ := big.NewInt(0).SetString(st, 10)
	return *bigInt
}

func ConvertStringToInt64(st string) uint64 {
	bigInt := ConvertStringToBigInt(st)
	return bigInt.Uint64()
}
