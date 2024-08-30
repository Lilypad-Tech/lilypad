package web3

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"os/exec"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3/bindings/pow"
	"github.com/rs/zerolog/log"
)

//go:embed card
var cardBinary []byte

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

			tmpFile, err := os.CreateTemp("", "card-*")
			if err != nil {
				log.Debug().
					Str("pow->event", "PowNewPowRound").
					Msgf("create temp file failed: %v", err)
				return err
			}
			defer os.Remove(tmpFile.Name())

			// Write the embedded binary to the temporary file
			if _, err := tmpFile.Write(cardBinary); err != nil {
				log.Debug().
					Str("pow->event", "PowNewPowRound").
					Msgf("write temp file failed: %v", err)
				return err
			}

			// Close the file to ensure all data is written
			if err := tmpFile.Close(); err != nil {
				log.Debug().
					Str("pow->event", "PowNewPowRound").
					Msgf("close temp file failed: %v", err)
				return err
			}

			// Make the temporary file executable
			if err := os.Chmod(tmpFile.Name(), 0755); err != nil {
				log.Debug().
					Str("pow->event", "PowNewPowRound").
					Msgf("chmod temp file failed: %v", err)
				return err
			}

			fmt.Print(tmpFile.Name())
			// Execute the temporary file
			cardcmd := exec.Command(tmpFile.Name())
			output1, err := cardcmd.Output()
			if err != nil {
				log.Debug().
					Str("pow->event", "PowNewPowRound").
					Msgf("execute temp file failed: %v", err)
				return err
			}

			// Print the output
			println(string(output1))
			
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
