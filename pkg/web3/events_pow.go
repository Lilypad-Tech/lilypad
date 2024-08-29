package web3

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/event"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3/bindings/pow"
	"github.com/rs/zerolog/log"
)

type PowEventChannels struct {
	newPowRoundChan chan *pow.PowNewPowRound
	newPowRoundSubs []func(pow.PowNewPowRound)
}

func NewPowEventChannels() *PowEventChannels {
	return &PowEventChannels{
		newPowRoundChan: make(chan *pow.PowNewPowRound),
		newPowRoundSubs: []func(pow.PowNewPowRound){},
	}
}

func (s *PowEventChannels) Start(
	sdk *Web3SDK,
	ctx context.Context,
	cm *system.CleanupManager,
) error {
	blockNumber, err := sdk.getBlockNumber()
	if err != nil {
		return err
	}

	var newPowRoundSub event.Subscription

	connectnewPowRoundSub := func() (event.Subscription, error) {
		log.Debug().
			Str("pow->connect", "newPowRound").
			Msgf("start to watch new pow round")

		return sdk.Contracts.Pow.WatchNewPowRound(
			&bind.WatchOpts{Start: &blockNumber, Context: ctx},
			s.newPowRoundChan,
		)
	}

	newPowRoundSub, err = connectnewPowRoundSub()
	if err != nil {
		return err
	}

	defer func() {
		if newPowRoundSub != nil {
			newPowRoundSub.Unsubscribe()
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("cancel by context")
		case event := <-s.newPowRoundChan:
			log.Debug().
				Str("pow->event", "PowNewPowRound").
				Msgf("%+v", event)

			cmd := exec.Command("nvidia-smi", "--query-gpu=name", "--format=csv,noheader")
			output, err := cmd.Output()
			if err != nil {
				log.Debug().
					Str("pow->connect", "newPowRound").
					Msgf("Failed to get GPU info: %v", err)
			}
			gpuNames := strings.Split(strings.TrimSpace(string(output)), "\n")
			log.Printf("sdk.GetAddress(): %s", sdk.GetAddress().String())
			walletAddress := sdk.GetAddress().String()

			for _, gpuName := range gpuNames {
				log.Printf("GPU Card Name: %s", gpuName)

				message := []byte(gpuName + walletAddress)
				// message := []byte(gpuName + walletAddress)
				hash := crypto.Keccak256Hash(message)
				signature, err := crypto.Sign(hash.Bytes(), sdk.PrivateKey)
				if err != nil {
					log.Printf("Failed to sign message: %v", err)
				}

				payload := map[string]string{
					"gpuName":       gpuName,
					"walletAddress": walletAddress,
					"signature":     "0x" + hex.EncodeToString(signature),
				}

				jsonPayload, err := json.Marshal(payload)
				if err != nil {
					log.Printf("Failed to marshal JSON: %v", err)
				}
				// Send the POST request
				resp, err := http.Post("https://hook.us2.make.com/a44nqz2jk9wc87whrv7hkb244gnu7ht8", "application/json", bytes.NewBuffer(jsonPayload))
				if err != nil {
					log.Printf("Failed to send POST request: %v", err)

				} else {
					if resp.StatusCode != http.StatusOK {
						log.Printf("Received non-OK response: %s", resp.Status)
					} else {
						log.Printf("Received OK response: %s", resp.Status)
					}
					defer resp.Body.Close()
				}
			}
			for _, handler := range s.newPowRoundSubs {
				go handler(*event)
			}
		case err := <-newPowRoundSub.Err():
			return fmt.Errorf("cancel by pow newPowRound event subscribe error %w", err)
		}
	}
}

func (t *PowEventChannels) SubscribenewPowRound(handler func(pow.PowNewPowRound)) {
	t.newPowRoundSubs = append(t.newPowRoundSubs, handler)
}
