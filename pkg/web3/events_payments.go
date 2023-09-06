package web3

import (
	"context"
	"math/big"

	"github.com/bacalhau-project/lilypad/pkg/web3/bindings/payments"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type PaymentEventChannels struct {
	Payment chan *payments.PaymentsPayment
}

func NewPaymentEventChannels() (*PaymentEventChannels, error) {
	return &PaymentEventChannels{
		Payment: make(chan *payments.PaymentsPayment),
	}, nil
}

func (paymentEventChannels *PaymentEventChannels) Listen(ctx context.Context, sdk *ContractSDK) error {
	blockNumber, err := sdk.getBlockNumber()
	if err != nil {
		return err
	}
	sub, err := sdk.Contracts.Payments.WatchPayment(
		&bind.WatchOpts{Start: &blockNumber, Context: ctx},
		paymentEventChannels.Payment,
		[]*big.Int{},
	)
	if err != nil {
		return err
	}
	<-ctx.Done()
	sub.Unsubscribe()
	return nil
}
