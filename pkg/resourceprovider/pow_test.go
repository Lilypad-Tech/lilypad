package resourceprovider

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestMiner(t *testing.T) {
	xxxx, _ := new(big.Int).SetString("58750003716598352816469", 10)
	fmt.Println("number ", xxxx.Text(2))

	maxUint256 := new(big.Int)
	maxUint256 = maxUint256.Exp(big.NewInt(2), big.NewInt(256), nil).Sub(maxUint256, big.NewInt(1))
	difficulty := maxUint256.Div(maxUint256, xxxx)

	addr := common.Address{}
	_, _ = rand.Read(addr[:])

	//maxDifficulty, _ := new(big.Int).SetString("0b1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111", 0)
	chanllenge := crypto.Keccak256([]byte{1, 2, 3})

	array := [32]byte{}
	copy(array[:], chanllenge)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	Mine(ctx, array, addr, difficulty)
}
