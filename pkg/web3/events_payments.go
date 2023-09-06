package web3

import (
	"context"
	"math/big"

	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/payments"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type PaymentEventChannels struct {
	paymentChan chan *payments.PaymentsPayment
	paymentSubs []func(*payments.PaymentsPayment)
}

func NewPaymentEventChannels() (*PaymentEventChannels, error) {
	return &PaymentEventChannels{
		paymentChan: make(chan *payments.PaymentsPayment),
	}, nil
}

func (p *PaymentEventChannels) Start(ctx context.Context, sdk *ContractSDK) error {
	blockNumber, err := sdk.getBlockNumber()
	if err != nil {
		return err
	}

	paymentSub, err := sdk.Contracts.Payments.WatchPayment(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		p.paymentChan,
		[]*big.Int{},
	)
	if err != nil {
		return err
	}

	go func() {
		<-ctx.Done()
		paymentSub.Unsubscribe()
	}()

	for {
		select {
		case event := <-p.paymentChan:
			for _, handler := range p.paymentSubs {
				handler(event)
			}
		case err := <-paymentSub.Err():
			return err
		}
	}
}

func (p *PaymentEventChannels) SubscribePayment(handler func(*payments.PaymentsPayment)) {
	p.paymentSubs = append(p.paymentSubs, handler)
}
