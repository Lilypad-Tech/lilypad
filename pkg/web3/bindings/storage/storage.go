// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package storage

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// SharedStructsAgreement is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsAgreement struct {
	State                    uint8
	Mediator                 common.Address
	ResourceProviderAgreedAt *big.Int
	JobCreatorAgreedAt       *big.Int
	DealAgreedAt             *big.Int
	ResultsSubmittedAt       *big.Int
	ResultsAcceptedAt        *big.Int
	ResultsCheckedAt         *big.Int
	MediationAcceptedAt      *big.Int
	MediationRejectedAt      *big.Int
	TimeoutSubmitResultsAt   *big.Int
	TimeoutJudgeResultsAt    *big.Int
	TimeoutMediateResultsAt  *big.Int
}

// SharedStructsDeal is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsDeal struct {
	DealId                    *big.Int
	ResourceProvider          common.Address
	JobCreator                common.Address
	InstructionPrice          *big.Int
	Timeout                   *big.Int
	TimeoutCollateral         *big.Int
	PaymentCollateral         *big.Int
	ResultsCollateralMultiple *big.Int
	MediationFee              *big.Int
}

// SharedStructsResult is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsResult struct {
	DealId           *big.Int
	ResultsId        *big.Int
	InstructionCount *big.Int
}

// SharedStructsUser is an auto generated low-level Go binding around an user-defined struct.
type SharedStructsUser struct {
	UserAddress        common.Address
	MetadataCID        *big.Int
	Url                string
	Roles              []uint8
	TrustedMediators   []common.Address
	TrustedDirectories []common.Address
}

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"DealStateChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"acceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"name\":\"addResult\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Result\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSharedStructs.ServiceType\",\"name\":\"serviceType\",\"type\":\"uint8\"}],\"name\":\"addUserToList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"agreeJobCreator\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"agreeResourceProvider\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"}],\"name\":\"checkResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"ensureDeal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getAgreement\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"mediator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resourceProviderAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jobCreatorAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dealAgreedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCheckedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationAcceptedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationRejectedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutSubmitResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutJudgeResultsAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutMediateResultsAt\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Agreement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getDeal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"instructionPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateralMultiple\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Deal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"party\",\"type\":\"address\"}],\"name\":\"getDealsForParty\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getJobCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getResult\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"instructionCount\",\"type\":\"uint256\"}],\"internalType\":\"structSharedStructs.Result\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"getResultsCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getUser\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"metadataCID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.ServiceType[]\",\"name\":\"roles\",\"type\":\"uint8[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedMediators\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedDirectories\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"hasDeal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"},{\"internalType\":\"enumSharedStructs.AgreementState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"isState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"mediationAcceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"mediationRejectResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSharedStructs.ServiceType\",\"name\":\"serviceType\",\"type\":\"uint8\"}],\"name\":\"removeUserFromList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSharedStructs.ServiceType\",\"name\":\"serviceType\",\"type\":\"uint8\"}],\"name\":\"showUsersInList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"timeoutJudgeResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"timeoutMediateResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dealId\",\"type\":\"uint256\"}],\"name\":\"timeoutSubmitResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"metadataCID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.ServiceType[]\",\"name\":\"roles\",\"type\":\"uint8[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedMediators\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedDirectories\",\"type\":\"address[]\"}],\"name\":\"updateUser\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"metadataCID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"enumSharedStructs.ServiceType[]\",\"name\":\"roles\",\"type\":\"uint8[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedMediators\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"trustedDirectories\",\"type\":\"address[]\"}],\"internalType\":\"structSharedStructs.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260018060146101000a81548160ff0219169083151502179055506001600260146101000a81548160ff02191690831515021790555034801562000045575f80fd5b50620000666200005a6200006c60201b60201c565b6200007360201b60201c565b62000134565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b61521980620001425f395ff3fe608060405234801561000f575f80fd5b50600436106101d8575f3560e01c80637fb9c45f11610102578063a4702958116100a0578063d74553011161006f578063d745530114610584578063ef816fd9146105b4578063f2fde38b146105e4578063f3d3d44814610600576101d8565b8063a470295814610512578063a7f96f061461051c578063aeb5ec0114610538578063c1390a8914610554576101d8565b80638462d54e116100dc5780638462d54e146104645780638da5cb5b14610494578063995e4339146104b2578063a15dcc8a146104e2576101d8565b80637fb9c45f1461040e5780638129fc1c1461042a57806382fd5bac14610434576101d8565b8063407c2d621161017a578063546cfd3411610149578063546cfd34146103885780635dd049fd146103b85780636f77926b146103d4578063715018a614610404576101d8565b8063407c2d62146102dc57806342e4074e1461030c5780634f9f6fe61461033c578063542785671461036c576101d8565b806323a9a862116101b657806323a9a8621461025857806332ebea041461027457806335a7e26814610290578063372725fa146102ac576101d8565b80630396e3c1146101dc57806311d5af331461020c578063172257c71461023c575b5f80fd5b6101f660048036038101906101f19190613185565b61061c565b60405161020391906131bf565b60405180910390f35b61022660048036038101906102219190613232565b610658565b6040516102339190613314565b60405180910390f35b61025660048036038101906102519190613185565b6106eb565b005b610272600480360381019061026d9190613185565b610765565b005b61028e60048036038101906102899190613185565b6107df565b005b6102aa60048036038101906102a59190613185565b610859565b005b6102c660048036038101906102c19190613334565b6108d3565b6040516102d391906134bc565b60405180910390f35b6102f660048036038101906102f19190613185565b61104f565b60405161030391906134f0565b60405180910390f35b61032660048036038101906103219190613185565b611066565b6040516103339190613685565b60405180910390f35b61035660048036038101906103519190613185565b611257565b6040516103639190613685565b60405180910390f35b61038660048036038101906103819190613185565b61137f565b005b6103a2600480360381019061039d919061369f565b6113f9565b6040516103af919061372f565b60405180910390f35b6103d260048036038101906103cd9190613185565b611505565b005b6103ee60048036038101906103e99190613232565b61157f565b6040516103fb9190613a03565b60405180910390f35b61040c611863565b005b61042860048036038101906104239190613a23565b611876565b005b61043261199d565b005b61044e60048036038101906104499190613185565b611ad4565b60405161045b91906134bc565b60405180910390f35b61047e60048036038101906104799190613185565b611bef565b60405161048b9190613685565b60405180910390f35b61049c611de0565b6040516104a99190613a70565b60405180910390f35b6104cc60048036038101906104c79190613185565b611e07565b6040516104d9919061372f565b60405180910390f35b6104fc60048036038101906104f79190613aac565b611e4f565b6040516105099190613b43565b60405180910390f35b61051a611f0f565b005b61053660048036038101906105319190613aac565b611f33565b005b610552600480360381019061054d9190613aac565b612237565b005b61056e60048036038101906105699190613e13565b612432565b60405161057b9190613a03565b60405180910390f35b61059e60048036038101906105999190613185565b612589565b6040516105ab91906131bf565b60405180910390f35b6105ce60048036038101906105c99190613f1d565b6125b9565b6040516105db91906134f0565b60405180910390f35b6105fe60048036038101906105f99190613232565b612644565b005b61061a60048036038101906106159190613232565b6126c6565b005b5f60085f8381526020019081526020015f206002015460055f8481526020019081526020015f20600301546106519190613f88565b9050919050565b606060065f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208054806020026020016040519081016040528092919081815260200182805480156106df57602002820191905f5260205f20905b8154815260200190600101908083116106cb575b50505050509050919050565b6106f36127ce565b506106ff8160026125b9565b61073e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161073590614049565b60405180910390fd5b4260075f8381526020019081526020015f20600501819055506107628160036128fb565b50565b61076d6127ce565b506107798160046125b9565b6107b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107af906140b1565b60405180910390fd5b4260075f8381526020019081526020015f20600701819055506107dc8160056128fb565b50565b6107e76127ce565b506107f38160016125b9565b610832576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161082990614119565b60405180910390fd5b4260075f8381526020019081526020015f20600901819055506108568160076128fb565b50565b6108616127ce565b5061086d8160046125b9565b6108ac576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108a3906140b1565b60405180910390fd5b4260075f8381526020019081526020015f20600b01819055506108d08160096128fb565b50565b6108db612e94565b6108e36127ce565b506108ee8a5f6125b9565b61092d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610924906141a7565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff160361099b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099290614235565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1603610a09576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a009061429d565b60405180910390fd5b8773ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff1603610a77576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a6e90614305565b60405180910390fd5b610a808a61104f565b15610d20575f610a8f8b611ad4565b90508973ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1614610b03576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610afa9061436d565b60405180910390fd5b8873ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff1614610b75576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b6c906143d5565b60405180910390fd5b87816060015114610bbb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bb29061443d565b60405180910390fd5b86816080015114610c01576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bf8906144a5565b60405180910390fd5b858160a0015114610c47576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c3e90614533565b60405180910390fd5b848160c0015114610c8d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c84906145c1565b60405180910390fd5b838160e0015114610cd3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cca9061464f565b60405180910390fd5b8281610100015114610d1a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d11906146b7565b60405180910390fd5b50610f34565b6040518061012001604052808b81526020018a73ffffffffffffffffffffffffffffffffffffffff1681526020018973ffffffffffffffffffffffffffffffffffffffff1681526020018881526020018781526020018681526020018581526020018481526020018381525060055f8c81526020019081526020015f205f820151815f01556020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e08201518160070155610100820151816008015590505060065f8a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208a908060018154018082558091505060019003905f5260205f20015f909190919091505560065f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208a908060018154018082558091505060019003905f5260205f20015f90919091909150555b60055f8b81526020019081526020015f20604051806101200160405290815f8201548152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382015481526020016004820154815260200160058201548152602001600682015481526020016007820154815260200160088201548152505090509998505050505050505050565b5f8061105a83611ad4565b5f015114159050919050565b61106e612f03565b6110766127ce565b506110808261104f565b6110bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110b69061471f565b60405180910390fd5b5f60075f8481526020019081526020015f206001015414611115576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161110c906147ad565b60405180910390fd5b4260075f8481526020019081526020015f206001018190555061113782612979565b60075f8381526020019081526020015f20604051806101a00160405290815f82015f9054906101000a900460ff16600981111561117757611176613509565b5b600981111561118957611188613509565b5b81526020015f820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b820154815250509050919050565b61125f612f03565b60075f8381526020019081526020015f20604051806101a00160405290815f82015f9054906101000a900460ff16600981111561129f5761129e613509565b5b60098111156112b1576112b0613509565b5b81526020015f820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b820154815250509050919050565b6113876127ce565b506113938160026125b9565b6113d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113c990614049565b60405180910390fd5b4260075f8381526020019081526020015f20600a01819055506113f68160086128fb565b50565b611401612f86565b6114096127ce565b506114158460016125b9565b611454576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161144b90614119565b60405180910390fd5b4260075f8681526020019081526020015f20600401819055506114788460026128fb565b60405180606001604052808581526020018481526020018381525060085f8681526020019081526020015f205f820151815f0155602082015181600101556040820151816002015590505060085f8581526020019081526020015f206040518060600160405290815f82015481526020016001820154815260200160028201548152505090509392505050565b61150d6127ce565b506115198160046125b9565b611558576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161154f906140b1565b60405180910390fd5b4260075f8381526020019081526020015f206008018190555061157c8160066128fb565b50565b611587612fa4565b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206040518060c00160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820154815260200160028201805461163d906147f8565b80601f0160208091040260200160405190810160405280929190818152602001828054611669906147f8565b80156116b45780601f1061168b576101008083540402835291602001916116b4565b820191905f5260205f20905b81548152906001019060200180831161169757829003601f168201915b505050505081526020016003820180548060200260200160405190810160405280929190818152602001828054801561173d57602002820191905f5260205f20905f905b82829054906101000a900460ff16600481111561171857611717613509565b5b815260200190600101906020825f010492830192600103820291508084116116f85790505b50505050508152602001600482018054806020026020016040519081016040528092919081815260200182805480156117c857602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161177f575b505050505081526020016005820180548060200260200160405190810160405280929190818152602001828054801561185357602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161180a575b5050505050815250509050919050565b61186b6129de565b6118745f612a5c565b565b61187e6127ce565b5061188a8260026125b9565b6118c9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118c090614049565b60405180910390fd5b5f6118d383611ad4565b90506118e3816020015183612b1d565b611922576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161191990614898565b60405180910390fd5b4260075f8581526020019081526020015f20600601819055508160075f8581526020019081526020015f205f0160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506119988360046128fb565b505050565b5f600160169054906101000a900460ff161590508080156119cf575060018060159054906101000a900460ff1660ff16105b806119fd57506119de30612c42565b1580156119fc575060018060159054906101000a900460ff1660ff16145b5b611a3c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a3390614926565b60405180910390fd5b60018060156101000a81548160ff021916908360ff1602179055508015611a785760018060166101000a81548160ff0219169083151502179055505b8015611ad1575f600160166101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051611ac89190614992565b60405180910390a15b50565b611adc612e94565b60055f8381526020019081526020015f20604051806101200160405290815f8201548152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815250509050919050565b611bf7612f03565b611bff6127ce565b50611c098261104f565b611c48576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c3f9061471f565b60405180910390fd5b5f60075f8481526020019081526020015f206002015414611c9e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c95906149f5565b60405180910390fd5b4260075f8481526020019081526020015f2060020181905550611cc082612979565b60075f8381526020019081526020015f20604051806101a00160405290815f82015f9054906101000a900460ff166009811115611d0057611cff613509565b5b6009811115611d1257611d11613509565b5b81526020015f820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815260200160098201548152602001600a8201548152602001600b820154815250509050919050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b611e0f612f86565b60085f8381526020019081526020015f206040518060600160405290815f8201548152602001600182015481526020016002820154815250509050919050565b606060045f836004811115611e6757611e66613509565b5b6004811115611e7957611e78613509565b5b81526020019081526020015f20805480602002602001604051908101604052809291908181526020018280548015611f0357602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611eba575b50505050509050919050565b611f176129de565b5f600160146101000a81548160ff021916908315150217905550565b5f73ffffffffffffffffffffffffffffffffffffffff1660035f3273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611fff576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ff690614a5d565b60405180910390fd5b5f8061200b8332612c64565b915091508061204f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161204690614ac5565b60405180910390fd5b5f8290505b600160045f86600481111561206c5761206b613509565b5b600481111561207e5761207d613509565b5b81526020019081526020015f20805490506120999190614ae3565b8110156121ba5760045f8560048111156120b6576120b5613509565b5b60048111156120c8576120c7613509565b5b81526020019081526020015f206001826120e29190614b16565b815481106120f3576120f2614b49565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660045f86600481111561213157612130613509565b5b600481111561214357612142613509565b5b81526020019081526020015f20828154811061216257612161614b49565b5b905f5260205f20015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080806121b290614b76565b915050612054565b5060045f8460048111156121d1576121d0613509565b5b60048111156121e3576121e2613509565b5b81526020019081526020015f20805480612200576121ff614bbd565b5b600190038181905f5260205f20015f6101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055505050565b5f73ffffffffffffffffffffffffffffffffffffffff1660035f3273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603612303576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122fa90614a5d565b60405180910390fd5b5f61230e8232612c64565b9150508015612352576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161234990614c5a565b60405180910390fd5b61235c8232612d7a565b61239b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161239290614cc2565b60405180910390fd5b60045f8360048111156123b1576123b0613509565b5b60048111156123c3576123c2613509565b5b81526020019081526020015f2032908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b61243a612fa4565b5f6040518060c001604052803273ffffffffffffffffffffffffffffffffffffffff1681526020018881526020018781526020018681526020018581526020018481525090508060035f3273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020190816125219190614e74565b50606082015181600301908051906020019061253e929190612fee565b50608082015181600401908051906020019061255b92919061309f565b5060a082015181600501908051906020019061257892919061309f565b509050508091505095945050505050565b5f6125938261061c565b60055f8481526020019081526020015f20600701546125b29190613f88565b9050919050565b5f6125c38361104f565b6125f5575f60098111156125da576125d9613509565b5b8260098111156125ed576125ec613509565b5b14905061263e565b81600981111561260857612607613509565b5b60075f8581526020019081526020015f205f015f9054906101000a900460ff16600981111561263a57612639613509565b5b1490505b92915050565b61264c6129de565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036126ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016126b190614fb3565b60405180910390fd5b6126c381612a5c565b50565b6126ce6129de565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361273c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161273390615041565b60405180910390fd5b600160149054906101000a900460ff1661278b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612782906150cf565b60405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f8073ffffffffffffffffffffffffffffffffffffffff1660015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361285e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161285590615041565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661289e612e8d565b73ffffffffffffffffffffffffffffffffffffffff16146128f4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128eb9061515d565b60405180910390fd5b6001905090565b8060075f8481526020019081526020015f205f015f6101000a81548160ff021916908360098111156129305761292f613509565b5b021790555080600981111561294857612947613509565b5b827f17d67e9978d93b39d6ad00aa2225ac1d172c5017e643f96f730bf3405e8fc55d60405160405180910390a35050565b5f60075f8381526020019081526020015f2060010154141580156129b157505f60075f8381526020019081526020015f206002015414155b156129db574260075f8381526020019081526020015f20600301819055506129da8160016128fb565b5b50565b6129e6612e8d565b73ffffffffffffffffffffffffffffffffffffffff16612a04611de0565b73ffffffffffffffffffffffffffffffffffffffff1614612a5a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a51906151c5565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f805f90505f5b60035f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060040180549050811015612c37578373ffffffffffffffffffffffffffffffffffffffff1660035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206004018281548110612bd857612bd7614b49565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603612c245760019150612c37565b8080612c2f90614b76565b915050612b24565b508091505092915050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f805f805f5b60045f886004811115612c8057612c7f613509565b5b6004811115612c9257612c91613509565b5b81526020019081526020015f2080549050811015612d6a578573ffffffffffffffffffffffffffffffffffffffff1660045f896004811115612cd757612cd6613509565b5b6004811115612ce957612ce8613509565b5b81526020019081526020015f208281548110612d0857612d07614b49565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603612d575780925060019150612d6a565b8080612d6290614b76565b915050612c6a565b5081819350935050509250929050565b5f805f90505f5b60035f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060030180549050811015612e8257846004811115612ddf57612dde613509565b5b60035f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206003018281548110612e3157612e30614b49565b5b905f5260205f2090602091828204019190069054906101000a900460ff166004811115612e6157612e60613509565b5b03612e6f5760019150612e82565b8080612e7a90614b76565b915050612d81565b508091505092915050565b5f33905090565b6040518061012001604052805f81526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81525090565b604051806101a001604052805f6009811115612f2257612f21613509565b5b81526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81526020015f81525090565b60405180606001604052805f81526020015f81526020015f81525090565b6040518060c001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f8152602001606081526020016060815260200160608152602001606081525090565b828054828255905f5260205f2090601f0160209004810192821561308e579160200282015f5b8382111561306057835183826101000a81548160ff021916908360048111156130405761303f613509565b5b021790555092602001926001016020815f01049283019260010302613014565b801561308c5782816101000a81549060ff02191690556001016020815f01049283019260010302613060565b505b50905061309b9190613126565b5090565b828054828255905f5260205f20908101928215613115579160200282015b82811115613114578251825f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550916020019190600101906130bd565b5b5090506131229190613126565b5090565b5b8082111561313d575f815f905550600101613127565b5090565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b61316481613152565b811461316e575f80fd5b50565b5f8135905061317f8161315b565b92915050565b5f6020828403121561319a5761319961314a565b5b5f6131a784828501613171565b91505092915050565b6131b981613152565b82525050565b5f6020820190506131d25f8301846131b0565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f613201826131d8565b9050919050565b613211816131f7565b811461321b575f80fd5b50565b5f8135905061322c81613208565b92915050565b5f602082840312156132475761324661314a565b5b5f6132548482850161321e565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61328f81613152565b82525050565b5f6132a08383613286565b60208301905092915050565b5f602082019050919050565b5f6132c28261325d565b6132cc8185613267565b93506132d783613277565b805f5b838110156133075781516132ee8882613295565b97506132f9836132ac565b9250506001810190506132da565b5085935050505092915050565b5f6020820190508181035f83015261332c81846132b8565b905092915050565b5f805f805f805f805f6101208a8c0312156133525761335161314a565b5b5f61335f8c828d01613171565b99505060206133708c828d0161321e565b98505060406133818c828d0161321e565b97505060606133928c828d01613171565b96505060806133a38c828d01613171565b95505060a06133b48c828d01613171565b94505060c06133c58c828d01613171565b93505060e06133d68c828d01613171565b9250506101006133e88c828d01613171565b9150509295985092959850929598565b613401816131f7565b82525050565b61012082015f82015161341c5f850182613286565b50602082015161342f60208501826133f8565b50604082015161344260408501826133f8565b5060608201516134556060850182613286565b5060808201516134686080850182613286565b5060a082015161347b60a0850182613286565b5060c082015161348e60c0850182613286565b5060e08201516134a160e0850182613286565b506101008201516134b6610100850182613286565b50505050565b5f610120820190506134d05f830184613407565b92915050565b5f8115159050919050565b6134ea816134d6565b82525050565b5f6020820190506135035f8301846134e1565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600a811061354757613546613509565b5b50565b5f81905061355782613536565b919050565b5f6135668261354a565b9050919050565b6135768161355c565b82525050565b6101a082015f8201516135915f85018261356d565b5060208201516135a460208501826133f8565b5060408201516135b76040850182613286565b5060608201516135ca6060850182613286565b5060808201516135dd6080850182613286565b5060a08201516135f060a0850182613286565b5060c082015161360360c0850182613286565b5060e082015161361660e0850182613286565b5061010082015161362b610100850182613286565b50610120820151613640610120850182613286565b50610140820151613655610140850182613286565b5061016082015161366a610160850182613286565b5061018082015161367f610180850182613286565b50505050565b5f6101a0820190506136995f83018461357c565b92915050565b5f805f606084860312156136b6576136b561314a565b5b5f6136c386828701613171565b93505060206136d486828701613171565b92505060406136e586828701613171565b9150509250925092565b606082015f8201516137035f850182613286565b5060208201516137166020850182613286565b5060408201516137296040850182613286565b50505050565b5f6060820190506137425f8301846136ef565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561377f578082015181840152602081019050613764565b5f8484015250505050565b5f601f19601f8301169050919050565b5f6137a482613748565b6137ae8185613752565b93506137be818560208601613762565b6137c78161378a565b840191505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6005811061380c5761380b613509565b5b50565b5f81905061381c826137fb565b919050565b5f61382b8261380f565b9050919050565b61383b81613821565b82525050565b5f61384c8383613832565b60208301905092915050565b5f602082019050919050565b5f61386e826137d2565b61387881856137dc565b9350613883836137ec565b805f5b838110156138b357815161389a8882613841565b97506138a583613858565b925050600181019050613886565b5085935050505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f6138f483836133f8565b60208301905092915050565b5f602082019050919050565b5f613916826138c0565b61392081856138ca565b935061392b836138da565b805f5b8381101561395b57815161394288826138e9565b975061394d83613900565b92505060018101905061392e565b5085935050505092915050565b5f60c083015f83015161397d5f8601826133f8565b5060208301516139906020860182613286565b50604083015184820360408601526139a8828261379a565b915050606083015184820360608601526139c28282613864565b915050608083015184820360808601526139dc828261390c565b91505060a083015184820360a08601526139f6828261390c565b9150508091505092915050565b5f6020820190508181035f830152613a1b8184613968565b905092915050565b5f8060408385031215613a3957613a3861314a565b5b5f613a4685828601613171565b9250506020613a578582860161321e565b9150509250929050565b613a6a816131f7565b82525050565b5f602082019050613a835f830184613a61565b92915050565b60058110613a95575f80fd5b50565b5f81359050613aa681613a89565b92915050565b5f60208284031215613ac157613ac061314a565b5b5f613ace84828501613a98565b91505092915050565b5f82825260208201905092915050565b5f613af1826138c0565b613afb8185613ad7565b9350613b06836138da565b805f5b83811015613b36578151613b1d88826138e9565b9750613b2883613900565b925050600181019050613b09565b5085935050505092915050565b5f6020820190508181035f830152613b5b8184613ae7565b905092915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b613ba18261378a565b810181811067ffffffffffffffff82111715613bc057613bbf613b6b565b5b80604052505050565b5f613bd2613141565b9050613bde8282613b98565b919050565b5f67ffffffffffffffff821115613bfd57613bfc613b6b565b5b613c068261378a565b9050602081019050919050565b828183375f83830152505050565b5f613c33613c2e84613be3565b613bc9565b905082815260208101848484011115613c4f57613c4e613b67565b5b613c5a848285613c13565b509392505050565b5f82601f830112613c7657613c75613b63565b5b8135613c86848260208601613c21565b91505092915050565b5f67ffffffffffffffff821115613ca957613ca8613b6b565b5b602082029050602081019050919050565b5f80fd5b5f613cd0613ccb84613c8f565b613bc9565b90508083825260208201905060208402830185811115613cf357613cf2613cba565b5b835b81811015613d1c5780613d088882613a98565b845260208401935050602081019050613cf5565b5050509392505050565b5f82601f830112613d3a57613d39613b63565b5b8135613d4a848260208601613cbe565b91505092915050565b5f67ffffffffffffffff821115613d6d57613d6c613b6b565b5b602082029050602081019050919050565b5f613d90613d8b84613d53565b613bc9565b90508083825260208201905060208402830185811115613db357613db2613cba565b5b835b81811015613ddc5780613dc8888261321e565b845260208401935050602081019050613db5565b5050509392505050565b5f82601f830112613dfa57613df9613b63565b5b8135613e0a848260208601613d7e565b91505092915050565b5f805f805f60a08688031215613e2c57613e2b61314a565b5b5f613e3988828901613171565b955050602086013567ffffffffffffffff811115613e5a57613e5961314e565b5b613e6688828901613c62565b945050604086013567ffffffffffffffff811115613e8757613e8661314e565b5b613e9388828901613d26565b935050606086013567ffffffffffffffff811115613eb457613eb361314e565b5b613ec088828901613de6565b925050608086013567ffffffffffffffff811115613ee157613ee061314e565b5b613eed88828901613de6565b9150509295509295909350565b600a8110613f06575f80fd5b50565b5f81359050613f1781613efa565b92915050565b5f8060408385031215613f3357613f3261314a565b5b5f613f4085828601613171565b9250506020613f5185828601613f09565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f613f9282613152565b9150613f9d83613152565b9250828202613fab81613152565b91508282048414831517613fc257613fc1613f5b565b5b5092915050565b5f82825260208201905092915050565b7f4465616c206e6f7420696e20526573756c74735375626d6974746564207374615f8201527f7465000000000000000000000000000000000000000000000000000000000000602082015250565b5f614033602283613fc9565b915061403e82613fd9565b604082019050919050565b5f6020820190508181035f83015261406081614027565b9050919050565b7f4465616c206e6f7420696e20526573756c7473436865636b65642073746174655f82015250565b5f61409b602083613fc9565b91506140a682614067565b602082019050919050565b5f6020820190508181035f8301526140c88161408f565b9050919050565b7f4465616c206e6f7420696e204465616c416772656564207374617465000000005f82015250565b5f614103601c83613fc9565b915061410e826140cf565b602082019050919050565b5f6020820190508181035f830152614130816140f7565b9050919050565b7f4465616c206973206e6f7420696e204465616c4e65676f74696174696e6720735f8201527f7461746500000000000000000000000000000000000000000000000000000000602082015250565b5f614191602483613fc9565b915061419c82614137565b604082019050919050565b5f6020820190508181035f8301526141be81614185565b9050919050565b7f5265736f757263652070726f7669646572206d75737420626520646566696e655f8201527f6400000000000000000000000000000000000000000000000000000000000000602082015250565b5f61421f602183613fc9565b915061422a826141c5565b604082019050919050565b5f6020820190508181035f83015261424c81614213565b9050919050565b7f4a6f622063726561746f72206d75737420626520646566696e656400000000005f82015250565b5f614287601b83613fc9565b915061429282614253565b602082019050919050565b5f6020820190508181035f8301526142b48161427b565b9050919050565b7f525020616e64204a432063616e6e6f74206265207468652073616d65000000005f82015250565b5f6142ef601c83613fc9565b91506142fa826142bb565b602082019050919050565b5f6020820190508181035f83015261431c816142e3565b9050919050565b7f525020646f6573206e6f74206d617463680000000000000000000000000000005f82015250565b5f614357601183613fc9565b915061436282614323565b602082019050919050565b5f6020820190508181035f8301526143848161434b565b9050919050565b7f4a4320646f6573206e6f74206d617463680000000000000000000000000000005f82015250565b5f6143bf601183613fc9565b91506143ca8261438b565b602082019050919050565b5f6020820190508181035f8301526143ec816143b3565b9050919050565b7f496e737472756374696f6e20707269636520646f6573206e6f74206d617463685f82015250565b5f614427602083613fc9565b9150614432826143f3565b602082019050919050565b5f6020820190508181035f8301526144548161441b565b9050919050565b7f54696d656f757420646f6573206e6f74206d61746368000000000000000000005f82015250565b5f61448f601683613fc9565b915061449a8261445b565b602082019050919050565b5f6020820190508181035f8301526144bc81614483565b9050919050565b7f54696d656f757420636f6c6c61746572616c20646f6573206e6f74206d6174635f8201527f6800000000000000000000000000000000000000000000000000000000000000602082015250565b5f61451d602183613fc9565b9150614528826144c3565b604082019050919050565b5f6020820190508181035f83015261454a81614511565b9050919050565b7f5061796d656e7420636f6c6c61746572616c20646f6573206e6f74206d6174635f8201527f6800000000000000000000000000000000000000000000000000000000000000602082015250565b5f6145ab602183613fc9565b91506145b682614551565b604082019050919050565b5f6020820190508181035f8301526145d88161459f565b9050919050565b7f526573756c747320636f6c6c61746572616c20646f6573206e6f74206d6174635f8201527f6800000000000000000000000000000000000000000000000000000000000000602082015250565b5f614639602183613fc9565b9150614644826145df565b604082019050919050565b5f6020820190508181035f8301526146668161462d565b9050919050565b7f4d6564696174696f6e2066656520646f6573206e6f74206d61746368000000005f82015250565b5f6146a1601c83613fc9565b91506146ac8261466d565b602082019050919050565b5f6020820190508181035f8301526146ce81614695565b9050919050565b7f4465616c20646f6573206e6f74206578697374000000000000000000000000005f82015250565b5f614709601383613fc9565b9150614714826146d5565b602082019050919050565b5f6020820190508181035f830152614736816146fd565b9050919050565b7f5265736f757263652070726f76696465722068617320616c72656164792061675f8201527f7265656400000000000000000000000000000000000000000000000000000000602082015250565b5f614797602483613fc9565b91506147a28261473d565b604082019050919050565b5f6020820190508181035f8301526147c48161478b565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061480f57607f821691505b602082108103614822576148216147cb565b5b50919050565b7f5265736f757263652070726f766964657220646f6573206e6f742074727573745f8201527f2074686174206d65646961746f72000000000000000000000000000000000000602082015250565b5f614882602e83613fc9565b915061488d82614828565b604082019050919050565b5f6020820190508181035f8301526148af81614876565b9050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f614910602e83613fc9565b915061491b826148b6565b604082019050919050565b5f6020820190508181035f83015261493d81614904565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f61497c61497761497284614944565b614959565b61494d565b9050919050565b61498c81614962565b82525050565b5f6020820190506149a55f830184614983565b92915050565b7f4a6f622063726561746f722068617320616c72656164792061677265656400005f82015250565b5f6149df601e83613fc9565b91506149ea826149ab565b602082019050919050565b5f6020820190508181035f830152614a0c816149d3565b9050919050565b7f55736572206d75737420657869737400000000000000000000000000000000005f82015250565b5f614a47600f83613fc9565b9150614a5282614a13565b602082019050919050565b5f6020820190508181035f830152614a7481614a3b565b9050919050565b7f55736572206973206e6f742070617274206f662074686174206c6973740000005f82015250565b5f614aaf601d83613fc9565b9150614aba82614a7b565b602082019050919050565b5f6020820190508181035f830152614adc81614aa3565b9050919050565b5f614aed82613152565b9150614af883613152565b9250828203905081811115614b1057614b0f613f5b565b5b92915050565b5f614b2082613152565b9150614b2b83613152565b9250828201905080821115614b4357614b42613f5b565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f614b8082613152565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614bb257614bb1613f5b565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f5573657220697320616c72656164792070617274206f662074686174206c69735f8201527f7400000000000000000000000000000000000000000000000000000000000000602082015250565b5f614c44602183613fc9565b9150614c4f82614bea565b604082019050919050565b5f6020820190508181035f830152614c7181614c38565b9050919050565b7f55736572206d7573742068617665207468617420726f6c6500000000000000005f82015250565b5f614cac601883613fc9565b9150614cb782614c78565b602082019050919050565b5f6020820190508181035f830152614cd981614ca0565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302614d3c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82614d01565b614d468683614d01565b95508019841693508086168417925050509392505050565b5f614d78614d73614d6e84613152565b614959565b613152565b9050919050565b5f819050919050565b614d9183614d5e565b614da5614d9d82614d7f565b848454614d0d565b825550505050565b5f90565b614db9614dad565b614dc4818484614d88565b505050565b5b81811015614de757614ddc5f82614db1565b600181019050614dca565b5050565b601f821115614e2c57614dfd81614ce0565b614e0684614cf2565b81016020851015614e15578190505b614e29614e2185614cf2565b830182614dc9565b50505b505050565b5f82821c905092915050565b5f614e4c5f1984600802614e31565b1980831691505092915050565b5f614e648383614e3d565b9150826002028217905092915050565b614e7d82613748565b67ffffffffffffffff811115614e9657614e95613b6b565b5b614ea082546147f8565b614eab828285614deb565b5f60209050601f831160018114614edc575f8415614eca578287015190505b614ed48582614e59565b865550614f3b565b601f198416614eea86614ce0565b5f5b82811015614f1157848901518255600182019150602085019450602081019050614eec565b86831015614f2e5784890151614f2a601f891682614e3d565b8355505b6001600288020188555050505b505050505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f614f9d602683613fc9565b9150614fa882614f43565b604082019050919050565b5f6020820190508181035f830152614fca81614f91565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a20436f6e74726f6c6c65722061645f8201527f6472657373206d75737420626520646566696e65640000000000000000000000602082015250565b5f61502b603583613fc9565b915061503682614fd1565b604082019050919050565b5f6020820190508181035f8301526150588161501f565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a2063616e4368616e6765436f6e745f8201527f726f6c6c6572416464726573732069732064697361626c656400000000000000602082015250565b5f6150b9603983613fc9565b91506150c48261505f565b604082019050919050565b5f6020820190508181035f8301526150e6816150ad565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a204f6e6c792074686520636f6e745f8201527f726f6c6c65722063616e2063616c6c2074686973206d6574686f640000000000602082015250565b5f615147603b83613fc9565b9150615152826150ed565b604082019050919050565b5f6020820190508181035f8301526151748161513b565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f6151af602083613fc9565b91506151ba8261517b565b602082019050919050565b5f6020820190508181035f8301526151dc816151a3565b905091905056fea26469706673582212204217e2f9f73161953e9039d0aa4632e27201ce8df2a8d23df90ec6ceae41457d64736f6c63430008150033",
}

// StorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageMetaData.ABI instead.
var StorageABI = StorageMetaData.ABI

// StorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageMetaData.Bin instead.
var StorageBin = StorageMetaData.Bin

// DeployStorage deploys a new Ethereum contract, binding an instance of Storage to it.
func DeployStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Storage, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// GetAgreement is a free data retrieval call binding the contract method 0x4f9f6fe6.
//
// Solidity: function getAgreement(uint256 dealId) view returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageCaller) GetAgreement(opts *bind.CallOpts, dealId *big.Int) (SharedStructsAgreement, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getAgreement", dealId)

	if err != nil {
		return *new(SharedStructsAgreement), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsAgreement)).(*SharedStructsAgreement)

	return out0, err

}

// GetAgreement is a free data retrieval call binding the contract method 0x4f9f6fe6.
//
// Solidity: function getAgreement(uint256 dealId) view returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageSession) GetAgreement(dealId *big.Int) (SharedStructsAgreement, error) {
	return _Storage.Contract.GetAgreement(&_Storage.CallOpts, dealId)
}

// GetAgreement is a free data retrieval call binding the contract method 0x4f9f6fe6.
//
// Solidity: function getAgreement(uint256 dealId) view returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageCallerSession) GetAgreement(dealId *big.Int) (SharedStructsAgreement, error) {
	return _Storage.Contract.GetAgreement(&_Storage.CallOpts, dealId)
}

// GetDeal is a free data retrieval call binding the contract method 0x82fd5bac.
//
// Solidity: function getDeal(uint256 dealId) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageCaller) GetDeal(opts *bind.CallOpts, dealId *big.Int) (SharedStructsDeal, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getDeal", dealId)

	if err != nil {
		return *new(SharedStructsDeal), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsDeal)).(*SharedStructsDeal)

	return out0, err

}

// GetDeal is a free data retrieval call binding the contract method 0x82fd5bac.
//
// Solidity: function getDeal(uint256 dealId) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageSession) GetDeal(dealId *big.Int) (SharedStructsDeal, error) {
	return _Storage.Contract.GetDeal(&_Storage.CallOpts, dealId)
}

// GetDeal is a free data retrieval call binding the contract method 0x82fd5bac.
//
// Solidity: function getDeal(uint256 dealId) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageCallerSession) GetDeal(dealId *big.Int) (SharedStructsDeal, error) {
	return _Storage.Contract.GetDeal(&_Storage.CallOpts, dealId)
}

// GetDealsForParty is a free data retrieval call binding the contract method 0x11d5af33.
//
// Solidity: function getDealsForParty(address party) view returns(uint256[])
func (_Storage *StorageCaller) GetDealsForParty(opts *bind.CallOpts, party common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getDealsForParty", party)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetDealsForParty is a free data retrieval call binding the contract method 0x11d5af33.
//
// Solidity: function getDealsForParty(address party) view returns(uint256[])
func (_Storage *StorageSession) GetDealsForParty(party common.Address) ([]*big.Int, error) {
	return _Storage.Contract.GetDealsForParty(&_Storage.CallOpts, party)
}

// GetDealsForParty is a free data retrieval call binding the contract method 0x11d5af33.
//
// Solidity: function getDealsForParty(address party) view returns(uint256[])
func (_Storage *StorageCallerSession) GetDealsForParty(party common.Address) ([]*big.Int, error) {
	return _Storage.Contract.GetDealsForParty(&_Storage.CallOpts, party)
}

// GetJobCost is a free data retrieval call binding the contract method 0x0396e3c1.
//
// Solidity: function getJobCost(uint256 dealId) view returns(uint256)
func (_Storage *StorageCaller) GetJobCost(opts *bind.CallOpts, dealId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getJobCost", dealId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetJobCost is a free data retrieval call binding the contract method 0x0396e3c1.
//
// Solidity: function getJobCost(uint256 dealId) view returns(uint256)
func (_Storage *StorageSession) GetJobCost(dealId *big.Int) (*big.Int, error) {
	return _Storage.Contract.GetJobCost(&_Storage.CallOpts, dealId)
}

// GetJobCost is a free data retrieval call binding the contract method 0x0396e3c1.
//
// Solidity: function getJobCost(uint256 dealId) view returns(uint256)
func (_Storage *StorageCallerSession) GetJobCost(dealId *big.Int) (*big.Int, error) {
	return _Storage.Contract.GetJobCost(&_Storage.CallOpts, dealId)
}

// GetResult is a free data retrieval call binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 dealId) view returns((uint256,uint256,uint256))
func (_Storage *StorageCaller) GetResult(opts *bind.CallOpts, dealId *big.Int) (SharedStructsResult, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getResult", dealId)

	if err != nil {
		return *new(SharedStructsResult), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsResult)).(*SharedStructsResult)

	return out0, err

}

// GetResult is a free data retrieval call binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 dealId) view returns((uint256,uint256,uint256))
func (_Storage *StorageSession) GetResult(dealId *big.Int) (SharedStructsResult, error) {
	return _Storage.Contract.GetResult(&_Storage.CallOpts, dealId)
}

// GetResult is a free data retrieval call binding the contract method 0x995e4339.
//
// Solidity: function getResult(uint256 dealId) view returns((uint256,uint256,uint256))
func (_Storage *StorageCallerSession) GetResult(dealId *big.Int) (SharedStructsResult, error) {
	return _Storage.Contract.GetResult(&_Storage.CallOpts, dealId)
}

// GetResultsCollateral is a free data retrieval call binding the contract method 0xd7455301.
//
// Solidity: function getResultsCollateral(uint256 dealId) view returns(uint256)
func (_Storage *StorageCaller) GetResultsCollateral(opts *bind.CallOpts, dealId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getResultsCollateral", dealId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResultsCollateral is a free data retrieval call binding the contract method 0xd7455301.
//
// Solidity: function getResultsCollateral(uint256 dealId) view returns(uint256)
func (_Storage *StorageSession) GetResultsCollateral(dealId *big.Int) (*big.Int, error) {
	return _Storage.Contract.GetResultsCollateral(&_Storage.CallOpts, dealId)
}

// GetResultsCollateral is a free data retrieval call binding the contract method 0xd7455301.
//
// Solidity: function getResultsCollateral(uint256 dealId) view returns(uint256)
func (_Storage *StorageCallerSession) GetResultsCollateral(dealId *big.Int) (*big.Int, error) {
	return _Storage.Contract.GetResultsCollateral(&_Storage.CallOpts, dealId)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) view returns((address,uint256,string,uint8[],address[],address[]))
func (_Storage *StorageCaller) GetUser(opts *bind.CallOpts, userAddress common.Address) (SharedStructsUser, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getUser", userAddress)

	if err != nil {
		return *new(SharedStructsUser), err
	}

	out0 := *abi.ConvertType(out[0], new(SharedStructsUser)).(*SharedStructsUser)

	return out0, err

}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) view returns((address,uint256,string,uint8[],address[],address[]))
func (_Storage *StorageSession) GetUser(userAddress common.Address) (SharedStructsUser, error) {
	return _Storage.Contract.GetUser(&_Storage.CallOpts, userAddress)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address userAddress) view returns((address,uint256,string,uint8[],address[],address[]))
func (_Storage *StorageCallerSession) GetUser(userAddress common.Address) (SharedStructsUser, error) {
	return _Storage.Contract.GetUser(&_Storage.CallOpts, userAddress)
}

// HasDeal is a free data retrieval call binding the contract method 0x407c2d62.
//
// Solidity: function hasDeal(uint256 dealId) view returns(bool)
func (_Storage *StorageCaller) HasDeal(opts *bind.CallOpts, dealId *big.Int) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "hasDeal", dealId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasDeal is a free data retrieval call binding the contract method 0x407c2d62.
//
// Solidity: function hasDeal(uint256 dealId) view returns(bool)
func (_Storage *StorageSession) HasDeal(dealId *big.Int) (bool, error) {
	return _Storage.Contract.HasDeal(&_Storage.CallOpts, dealId)
}

// HasDeal is a free data retrieval call binding the contract method 0x407c2d62.
//
// Solidity: function hasDeal(uint256 dealId) view returns(bool)
func (_Storage *StorageCallerSession) HasDeal(dealId *big.Int) (bool, error) {
	return _Storage.Contract.HasDeal(&_Storage.CallOpts, dealId)
}

// IsState is a free data retrieval call binding the contract method 0xef816fd9.
//
// Solidity: function isState(uint256 dealId, uint8 state) view returns(bool)
func (_Storage *StorageCaller) IsState(opts *bind.CallOpts, dealId *big.Int, state uint8) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "isState", dealId, state)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsState is a free data retrieval call binding the contract method 0xef816fd9.
//
// Solidity: function isState(uint256 dealId, uint8 state) view returns(bool)
func (_Storage *StorageSession) IsState(dealId *big.Int, state uint8) (bool, error) {
	return _Storage.Contract.IsState(&_Storage.CallOpts, dealId, state)
}

// IsState is a free data retrieval call binding the contract method 0xef816fd9.
//
// Solidity: function isState(uint256 dealId, uint8 state) view returns(bool)
func (_Storage *StorageCallerSession) IsState(dealId *big.Int, state uint8) (bool, error) {
	return _Storage.Contract.IsState(&_Storage.CallOpts, dealId, state)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Storage *StorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Storage *StorageSession) Owner() (common.Address, error) {
	return _Storage.Contract.Owner(&_Storage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Storage *StorageCallerSession) Owner() (common.Address, error) {
	return _Storage.Contract.Owner(&_Storage.CallOpts)
}

// ShowUsersInList is a free data retrieval call binding the contract method 0xa15dcc8a.
//
// Solidity: function showUsersInList(uint8 serviceType) view returns(address[])
func (_Storage *StorageCaller) ShowUsersInList(opts *bind.CallOpts, serviceType uint8) ([]common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "showUsersInList", serviceType)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ShowUsersInList is a free data retrieval call binding the contract method 0xa15dcc8a.
//
// Solidity: function showUsersInList(uint8 serviceType) view returns(address[])
func (_Storage *StorageSession) ShowUsersInList(serviceType uint8) ([]common.Address, error) {
	return _Storage.Contract.ShowUsersInList(&_Storage.CallOpts, serviceType)
}

// ShowUsersInList is a free data retrieval call binding the contract method 0xa15dcc8a.
//
// Solidity: function showUsersInList(uint8 serviceType) view returns(address[])
func (_Storage *StorageCallerSession) ShowUsersInList(serviceType uint8) ([]common.Address, error) {
	return _Storage.Contract.ShowUsersInList(&_Storage.CallOpts, serviceType)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x172257c7.
//
// Solidity: function acceptResult(uint256 dealId) returns()
func (_Storage *StorageTransactor) AcceptResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "acceptResult", dealId)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x172257c7.
//
// Solidity: function acceptResult(uint256 dealId) returns()
func (_Storage *StorageSession) AcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AcceptResult(&_Storage.TransactOpts, dealId)
}

// AcceptResult is a paid mutator transaction binding the contract method 0x172257c7.
//
// Solidity: function acceptResult(uint256 dealId) returns()
func (_Storage *StorageTransactorSession) AcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AcceptResult(&_Storage.TransactOpts, dealId)
}

// AddResult is a paid mutator transaction binding the contract method 0x546cfd34.
//
// Solidity: function addResult(uint256 dealId, uint256 resultsId, uint256 instructionCount) returns((uint256,uint256,uint256))
func (_Storage *StorageTransactor) AddResult(opts *bind.TransactOpts, dealId *big.Int, resultsId *big.Int, instructionCount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addResult", dealId, resultsId, instructionCount)
}

// AddResult is a paid mutator transaction binding the contract method 0x546cfd34.
//
// Solidity: function addResult(uint256 dealId, uint256 resultsId, uint256 instructionCount) returns((uint256,uint256,uint256))
func (_Storage *StorageSession) AddResult(dealId *big.Int, resultsId *big.Int, instructionCount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddResult(&_Storage.TransactOpts, dealId, resultsId, instructionCount)
}

// AddResult is a paid mutator transaction binding the contract method 0x546cfd34.
//
// Solidity: function addResult(uint256 dealId, uint256 resultsId, uint256 instructionCount) returns((uint256,uint256,uint256))
func (_Storage *StorageTransactorSession) AddResult(dealId *big.Int, resultsId *big.Int, instructionCount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddResult(&_Storage.TransactOpts, dealId, resultsId, instructionCount)
}

// AddUserToList is a paid mutator transaction binding the contract method 0xaeb5ec01.
//
// Solidity: function addUserToList(uint8 serviceType) returns()
func (_Storage *StorageTransactor) AddUserToList(opts *bind.TransactOpts, serviceType uint8) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addUserToList", serviceType)
}

// AddUserToList is a paid mutator transaction binding the contract method 0xaeb5ec01.
//
// Solidity: function addUserToList(uint8 serviceType) returns()
func (_Storage *StorageSession) AddUserToList(serviceType uint8) (*types.Transaction, error) {
	return _Storage.Contract.AddUserToList(&_Storage.TransactOpts, serviceType)
}

// AddUserToList is a paid mutator transaction binding the contract method 0xaeb5ec01.
//
// Solidity: function addUserToList(uint8 serviceType) returns()
func (_Storage *StorageTransactorSession) AddUserToList(serviceType uint8) (*types.Transaction, error) {
	return _Storage.Contract.AddUserToList(&_Storage.TransactOpts, serviceType)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x8462d54e.
//
// Solidity: function agreeJobCreator(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageTransactor) AgreeJobCreator(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "agreeJobCreator", dealId)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x8462d54e.
//
// Solidity: function agreeJobCreator(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageSession) AgreeJobCreator(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AgreeJobCreator(&_Storage.TransactOpts, dealId)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0x8462d54e.
//
// Solidity: function agreeJobCreator(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageTransactorSession) AgreeJobCreator(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AgreeJobCreator(&_Storage.TransactOpts, dealId)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x42e4074e.
//
// Solidity: function agreeResourceProvider(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageTransactor) AgreeResourceProvider(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "agreeResourceProvider", dealId)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x42e4074e.
//
// Solidity: function agreeResourceProvider(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageSession) AgreeResourceProvider(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AgreeResourceProvider(&_Storage.TransactOpts, dealId)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x42e4074e.
//
// Solidity: function agreeResourceProvider(uint256 dealId) returns((uint8,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageTransactorSession) AgreeResourceProvider(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AgreeResourceProvider(&_Storage.TransactOpts, dealId)
}

// CheckResult is a paid mutator transaction binding the contract method 0x7fb9c45f.
//
// Solidity: function checkResult(uint256 dealId, address mediator) returns()
func (_Storage *StorageTransactor) CheckResult(opts *bind.TransactOpts, dealId *big.Int, mediator common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "checkResult", dealId, mediator)
}

// CheckResult is a paid mutator transaction binding the contract method 0x7fb9c45f.
//
// Solidity: function checkResult(uint256 dealId, address mediator) returns()
func (_Storage *StorageSession) CheckResult(dealId *big.Int, mediator common.Address) (*types.Transaction, error) {
	return _Storage.Contract.CheckResult(&_Storage.TransactOpts, dealId, mediator)
}

// CheckResult is a paid mutator transaction binding the contract method 0x7fb9c45f.
//
// Solidity: function checkResult(uint256 dealId, address mediator) returns()
func (_Storage *StorageTransactorSession) CheckResult(dealId *big.Int, mediator common.Address) (*types.Transaction, error) {
	return _Storage.Contract.CheckResult(&_Storage.TransactOpts, dealId, mediator)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Storage *StorageTransactor) DisableChangeControllerAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "disableChangeControllerAddress")
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Storage *StorageSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _Storage.Contract.DisableChangeControllerAddress(&_Storage.TransactOpts)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Storage *StorageTransactorSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _Storage.Contract.DisableChangeControllerAddress(&_Storage.TransactOpts)
}

// EnsureDeal is a paid mutator transaction binding the contract method 0x372725fa.
//
// Solidity: function ensureDeal(uint256 dealId, address resourceProvider, address jobCreator, uint256 instructionPrice, uint256 timeout, uint256 timeoutCollateral, uint256 paymentCollateral, uint256 resultsCollateralMultiple, uint256 mediationFee) returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageTransactor) EnsureDeal(opts *bind.TransactOpts, dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, instructionPrice *big.Int, timeout *big.Int, timeoutCollateral *big.Int, paymentCollateral *big.Int, resultsCollateralMultiple *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "ensureDeal", dealId, resourceProvider, jobCreator, instructionPrice, timeout, timeoutCollateral, paymentCollateral, resultsCollateralMultiple, mediationFee)
}

// EnsureDeal is a paid mutator transaction binding the contract method 0x372725fa.
//
// Solidity: function ensureDeal(uint256 dealId, address resourceProvider, address jobCreator, uint256 instructionPrice, uint256 timeout, uint256 timeoutCollateral, uint256 paymentCollateral, uint256 resultsCollateralMultiple, uint256 mediationFee) returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageSession) EnsureDeal(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, instructionPrice *big.Int, timeout *big.Int, timeoutCollateral *big.Int, paymentCollateral *big.Int, resultsCollateralMultiple *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.EnsureDeal(&_Storage.TransactOpts, dealId, resourceProvider, jobCreator, instructionPrice, timeout, timeoutCollateral, paymentCollateral, resultsCollateralMultiple, mediationFee)
}

// EnsureDeal is a paid mutator transaction binding the contract method 0x372725fa.
//
// Solidity: function ensureDeal(uint256 dealId, address resourceProvider, address jobCreator, uint256 instructionPrice, uint256 timeout, uint256 timeoutCollateral, uint256 paymentCollateral, uint256 resultsCollateralMultiple, uint256 mediationFee) returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageTransactorSession) EnsureDeal(dealId *big.Int, resourceProvider common.Address, jobCreator common.Address, instructionPrice *big.Int, timeout *big.Int, timeoutCollateral *big.Int, paymentCollateral *big.Int, resultsCollateralMultiple *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.EnsureDeal(&_Storage.TransactOpts, dealId, resourceProvider, jobCreator, instructionPrice, timeout, timeoutCollateral, paymentCollateral, resultsCollateralMultiple, mediationFee)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Storage *StorageTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Storage *StorageSession) Initialize() (*types.Transaction, error) {
	return _Storage.Contract.Initialize(&_Storage.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Storage *StorageTransactorSession) Initialize() (*types.Transaction, error) {
	return _Storage.Contract.Initialize(&_Storage.TransactOpts)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x23a9a862.
//
// Solidity: function mediationAcceptResult(uint256 dealId) returns()
func (_Storage *StorageTransactor) MediationAcceptResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "mediationAcceptResult", dealId)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x23a9a862.
//
// Solidity: function mediationAcceptResult(uint256 dealId) returns()
func (_Storage *StorageSession) MediationAcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.MediationAcceptResult(&_Storage.TransactOpts, dealId)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x23a9a862.
//
// Solidity: function mediationAcceptResult(uint256 dealId) returns()
func (_Storage *StorageTransactorSession) MediationAcceptResult(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.MediationAcceptResult(&_Storage.TransactOpts, dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x5dd049fd.
//
// Solidity: function mediationRejectResult(uint256 dealId) returns()
func (_Storage *StorageTransactor) MediationRejectResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "mediationRejectResult", dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x5dd049fd.
//
// Solidity: function mediationRejectResult(uint256 dealId) returns()
func (_Storage *StorageSession) MediationRejectResult(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.MediationRejectResult(&_Storage.TransactOpts, dealId)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0x5dd049fd.
//
// Solidity: function mediationRejectResult(uint256 dealId) returns()
func (_Storage *StorageTransactorSession) MediationRejectResult(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.MediationRejectResult(&_Storage.TransactOpts, dealId)
}

// RemoveUserFromList is a paid mutator transaction binding the contract method 0xa7f96f06.
//
// Solidity: function removeUserFromList(uint8 serviceType) returns()
func (_Storage *StorageTransactor) RemoveUserFromList(opts *bind.TransactOpts, serviceType uint8) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "removeUserFromList", serviceType)
}

// RemoveUserFromList is a paid mutator transaction binding the contract method 0xa7f96f06.
//
// Solidity: function removeUserFromList(uint8 serviceType) returns()
func (_Storage *StorageSession) RemoveUserFromList(serviceType uint8) (*types.Transaction, error) {
	return _Storage.Contract.RemoveUserFromList(&_Storage.TransactOpts, serviceType)
}

// RemoveUserFromList is a paid mutator transaction binding the contract method 0xa7f96f06.
//
// Solidity: function removeUserFromList(uint8 serviceType) returns()
func (_Storage *StorageTransactorSession) RemoveUserFromList(serviceType uint8) (*types.Transaction, error) {
	return _Storage.Contract.RemoveUserFromList(&_Storage.TransactOpts, serviceType)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Storage *StorageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Storage *StorageSession) RenounceOwnership() (*types.Transaction, error) {
	return _Storage.Contract.RenounceOwnership(&_Storage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Storage *StorageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Storage.Contract.RenounceOwnership(&_Storage.TransactOpts)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Storage *StorageTransactor) SetControllerAddress(opts *bind.TransactOpts, _controllerAddress common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setControllerAddress", _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Storage *StorageSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetControllerAddress(&_Storage.TransactOpts, _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Storage *StorageTransactorSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetControllerAddress(&_Storage.TransactOpts, _controllerAddress)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0x54278567.
//
// Solidity: function timeoutJudgeResult(uint256 dealId) returns()
func (_Storage *StorageTransactor) TimeoutJudgeResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "timeoutJudgeResult", dealId)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0x54278567.
//
// Solidity: function timeoutJudgeResult(uint256 dealId) returns()
func (_Storage *StorageSession) TimeoutJudgeResult(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutJudgeResult(&_Storage.TransactOpts, dealId)
}

// TimeoutJudgeResult is a paid mutator transaction binding the contract method 0x54278567.
//
// Solidity: function timeoutJudgeResult(uint256 dealId) returns()
func (_Storage *StorageTransactorSession) TimeoutJudgeResult(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutJudgeResult(&_Storage.TransactOpts, dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0x35a7e268.
//
// Solidity: function timeoutMediateResult(uint256 dealId) returns()
func (_Storage *StorageTransactor) TimeoutMediateResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "timeoutMediateResult", dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0x35a7e268.
//
// Solidity: function timeoutMediateResult(uint256 dealId) returns()
func (_Storage *StorageSession) TimeoutMediateResult(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutMediateResult(&_Storage.TransactOpts, dealId)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0x35a7e268.
//
// Solidity: function timeoutMediateResult(uint256 dealId) returns()
func (_Storage *StorageTransactorSession) TimeoutMediateResult(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutMediateResult(&_Storage.TransactOpts, dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x32ebea04.
//
// Solidity: function timeoutSubmitResult(uint256 dealId) returns()
func (_Storage *StorageTransactor) TimeoutSubmitResult(opts *bind.TransactOpts, dealId *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "timeoutSubmitResult", dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x32ebea04.
//
// Solidity: function timeoutSubmitResult(uint256 dealId) returns()
func (_Storage *StorageSession) TimeoutSubmitResult(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutSubmitResult(&_Storage.TransactOpts, dealId)
}

// TimeoutSubmitResult is a paid mutator transaction binding the contract method 0x32ebea04.
//
// Solidity: function timeoutSubmitResult(uint256 dealId) returns()
func (_Storage *StorageTransactorSession) TimeoutSubmitResult(dealId *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.TimeoutSubmitResult(&_Storage.TransactOpts, dealId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Storage *StorageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Storage *StorageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Storage.Contract.TransferOwnership(&_Storage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Storage *StorageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Storage.Contract.TransferOwnership(&_Storage.TransactOpts, newOwner)
}

// UpdateUser is a paid mutator transaction binding the contract method 0xc1390a89.
//
// Solidity: function updateUser(uint256 metadataCID, string url, uint8[] roles, address[] trustedMediators, address[] trustedDirectories) returns((address,uint256,string,uint8[],address[],address[]))
func (_Storage *StorageTransactor) UpdateUser(opts *bind.TransactOpts, metadataCID *big.Int, url string, roles []uint8, trustedMediators []common.Address, trustedDirectories []common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "updateUser", metadataCID, url, roles, trustedMediators, trustedDirectories)
}

// UpdateUser is a paid mutator transaction binding the contract method 0xc1390a89.
//
// Solidity: function updateUser(uint256 metadataCID, string url, uint8[] roles, address[] trustedMediators, address[] trustedDirectories) returns((address,uint256,string,uint8[],address[],address[]))
func (_Storage *StorageSession) UpdateUser(metadataCID *big.Int, url string, roles []uint8, trustedMediators []common.Address, trustedDirectories []common.Address) (*types.Transaction, error) {
	return _Storage.Contract.UpdateUser(&_Storage.TransactOpts, metadataCID, url, roles, trustedMediators, trustedDirectories)
}

// UpdateUser is a paid mutator transaction binding the contract method 0xc1390a89.
//
// Solidity: function updateUser(uint256 metadataCID, string url, uint8[] roles, address[] trustedMediators, address[] trustedDirectories) returns((address,uint256,string,uint8[],address[],address[]))
func (_Storage *StorageTransactorSession) UpdateUser(metadataCID *big.Int, url string, roles []uint8, trustedMediators []common.Address, trustedDirectories []common.Address) (*types.Transaction, error) {
	return _Storage.Contract.UpdateUser(&_Storage.TransactOpts, metadataCID, url, roles, trustedMediators, trustedDirectories)
}

// StorageDealStateChangeIterator is returned from FilterDealStateChange and is used to iterate over the raw logs and unpacked data for DealStateChange events raised by the Storage contract.
type StorageDealStateChangeIterator struct {
	Event *StorageDealStateChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageDealStateChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageDealStateChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageDealStateChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageDealStateChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageDealStateChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageDealStateChange represents a DealStateChange event raised by the Storage contract.
type StorageDealStateChange struct {
	DealId *big.Int
	State  uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDealStateChange is a free log retrieval operation binding the contract event 0x17d67e9978d93b39d6ad00aa2225ac1d172c5017e643f96f730bf3405e8fc55d.
//
// Solidity: event DealStateChange(uint256 indexed dealId, uint8 indexed state)
func (_Storage *StorageFilterer) FilterDealStateChange(opts *bind.FilterOpts, dealId []*big.Int, state []uint8) (*StorageDealStateChangeIterator, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}
	var stateRule []interface{}
	for _, stateItem := range state {
		stateRule = append(stateRule, stateItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "DealStateChange", dealIdRule, stateRule)
	if err != nil {
		return nil, err
	}
	return &StorageDealStateChangeIterator{contract: _Storage.contract, event: "DealStateChange", logs: logs, sub: sub}, nil
}

// WatchDealStateChange is a free log subscription operation binding the contract event 0x17d67e9978d93b39d6ad00aa2225ac1d172c5017e643f96f730bf3405e8fc55d.
//
// Solidity: event DealStateChange(uint256 indexed dealId, uint8 indexed state)
func (_Storage *StorageFilterer) WatchDealStateChange(opts *bind.WatchOpts, sink chan<- *StorageDealStateChange, dealId []*big.Int, state []uint8) (event.Subscription, error) {

	var dealIdRule []interface{}
	for _, dealIdItem := range dealId {
		dealIdRule = append(dealIdRule, dealIdItem)
	}
	var stateRule []interface{}
	for _, stateItem := range state {
		stateRule = append(stateRule, stateItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "DealStateChange", dealIdRule, stateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageDealStateChange)
				if err := _Storage.contract.UnpackLog(event, "DealStateChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDealStateChange is a log parse operation binding the contract event 0x17d67e9978d93b39d6ad00aa2225ac1d172c5017e643f96f730bf3405e8fc55d.
//
// Solidity: event DealStateChange(uint256 indexed dealId, uint8 indexed state)
func (_Storage *StorageFilterer) ParseDealStateChange(log types.Log) (*StorageDealStateChange, error) {
	event := new(StorageDealStateChange)
	if err := _Storage.contract.UnpackLog(event, "DealStateChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Storage contract.
type StorageInitializedIterator struct {
	Event *StorageInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageInitialized represents a Initialized event raised by the Storage contract.
type StorageInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Storage *StorageFilterer) FilterInitialized(opts *bind.FilterOpts) (*StorageInitializedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StorageInitializedIterator{contract: _Storage.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Storage *StorageFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StorageInitialized) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageInitialized)
				if err := _Storage.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Storage *StorageFilterer) ParseInitialized(log types.Log) (*StorageInitialized, error) {
	event := new(StorageInitialized)
	if err := _Storage.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Storage contract.
type StorageOwnershipTransferredIterator struct {
	Event *StorageOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageOwnershipTransferred represents a OwnershipTransferred event raised by the Storage contract.
type StorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Storage *StorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StorageOwnershipTransferredIterator{contract: _Storage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Storage *StorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageOwnershipTransferred)
				if err := _Storage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Storage *StorageFilterer) ParseOwnershipTransferred(log types.Log) (*StorageOwnershipTransferred, error) {
	event := new(StorageOwnershipTransferred)
	if err := _Storage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
