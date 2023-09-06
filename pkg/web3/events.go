package web3

import "context"

type EventChannels struct {
	Token       *TokenEventChannels
	Payment     *PaymentEventChannels
	Storage     *StorageEventChannels
	Controller  *ControllerEventChannels
	collections []EventChannelCollection
}

func NewEventChannels(sdk *ContractSDK) (*EventChannels, error) {
	tokenChannels, err := NewTokenEventChannels()
	if err != nil {
		return nil, err
	}
	paymentChannels, err := NewPaymentEventChannels()
	if err != nil {
		return nil, err
	}
	storageChannels, err := NewStorageEventChannels()
	if err != nil {
		return nil, err
	}
	controllerChannels, err := NewControllerEventChannels()
	if err != nil {
		return nil, err
	}
	collections := []EventChannelCollection{
		tokenChannels,
		paymentChannels,
		storageChannels,
		controllerChannels,
	}
	return &EventChannels{
		Token:       tokenChannels,
		Payment:     paymentChannels,
		Storage:     storageChannels,
		Controller:  controllerChannels,
		collections: collections,
	}, nil
}

func (eventChannels *EventChannels) Listen(ctx context.Context, sdk *ContractSDK) error {
	for _, collection := range eventChannels.collections {
		err := collection.Listen(ctx, sdk)
		if err != nil {
			return err
		}
	}
	return nil
}
