package web3

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"
	"github.com/lilypad-tech/lilypad/pkg/system"
	"github.com/lilypad-tech/lilypad/pkg/web3/bindings/payments"
	"github.com/rs/zerolog/log"
)

type PaymentEventChannels struct {
	paymentChan chan *payments.PaymentsPayment
	paymentSubs []func(payments.PaymentsPayment)
}

func NewPaymentEventChannels() *PaymentEventChannels {
	return &PaymentEventChannels{
		paymentChan: make(chan *payments.PaymentsPayment),
		paymentSubs: []func(payments.PaymentsPayment){},
	}
}

func (p *PaymentEventChannels) Start(
	sdk *Web3SDK,
	ctx context.Context,
	cm *system.CleanupManager,
) error {
	blockNumber, err := sdk.getBlockNumber()
	if err != nil {
		return err
	}

	var paymentSub event.Subscription

	connectPaymentSub := func() (event.Subscription, error) {
		log.Debug().
			Str("payments->connect", "Payment").
			Msgf("")
		return sdk.Contracts.Payments.WatchPayment(
			&bind.WatchOpts{Start: &blockNumber, Context: ctx},
			p.paymentChan,
		)
	}

	paymentSub, err = connectPaymentSub()
	if err != nil {
		return err
	}
	cm.RegisterCallback(unsubscribeSub(paymentSub))

	for {
		select {
		case <-ctx.Done():
			return nil
		case event := <-p.paymentChan:
			log.Debug().
				Str("payments->event", "Payment").
				Msgf("%+v", event)
			for _, handler := range p.paymentSubs {
				go handler(*event)
			}
		case err := <-paymentSub.Err():
			if err != nil {
				return fmt.Errorf("cancel by mediation MediationRequested event subscribe error %w", err)
			}
			return nil
		}
	}
}

func (p *PaymentEventChannels) SubscribePayment(handler func(payments.PaymentsPayment)) {
	p.paymentSubs = append(p.paymentSubs, handler)
}
