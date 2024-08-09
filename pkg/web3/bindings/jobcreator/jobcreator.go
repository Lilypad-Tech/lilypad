// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package jobcreator

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

// JobcreatorMetaData contains all meta data concerning the Jobcreator contract.
var JobcreatorMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"calling_contract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"module\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"inputs\",\"type\":\"string[]\"}],\"name\":\"JobAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getControllerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequiredDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextJobID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requiredDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"module\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"inputs\",\"type\":\"string[]\"},{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"runJob\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"name\":\"setRequiredDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"setTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataId\",\"type\":\"string\"}],\"name\":\"submitResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260018060146101000a81548160ff021916908315150217905550348015610029575f80fd5b5061004661003b61004b60201b60201c565b61005260201b60201c565b610113565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b611df4806101205f395ff3fe608060405234801561000f575f80fd5b50600436106100f3575f3560e01c8063a470295811610095578063d2a715c011610064578063d2a715c014610223578063f2fde38b14610241578063f3d3d4481461025d578063fb7cfdd714610279576100f3565b8063a4702958146101af578063c4d66de8146101b9578063c57380a2146101d5578063c75555fa146101f3576100f3565b806358e56db4116100d157806358e56db41461014d5780636c0f1f581461016b578063715018a6146101875780638da5cb5b14610191576100f3565b806310fe9ae8146100f757806326a4e8d2146101155780634c526c7614610131575b5f80fd5b6100ff610297565b60405161010c9190610eb1565b60405180910390f35b61012f600480360381019061012a9190610f05565b6102bf565b005b61014b60048036038101906101469190610f63565b6103d9565b005b61015561042e565b6040516101629190610f9d565b60405180910390f35b610185600480360381019061018091906110f2565b610437565b005b61018f61052c565b005b61019961053f565b6040516101a69190610eb1565b60405180910390f35b6101b7610566565b005b6101d360048036038101906101ce9190610f05565b61058a565b005b6101dd6106d2565b6040516101ea9190610eb1565b60405180910390f35b61020d6004803603810190610208919061125c565b6106fa565b60405161021a9190610f9d565b60405180910390f35b61022b610975565b6040516102389190610f9d565b60405180910390f35b61025b60048036038101906102569190610f05565b61097b565b005b61027760048036038101906102729190610f05565b6109fd565b005b610281610b05565b60405161028e9190610f9d565b60405180910390f35b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6102c7610b0b565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610335576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032c9061133e565b60405180910390fd5b8060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6103e1610b89565b505f8111610424576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041b906113a6565b60405180910390fd5b8060048190555050565b5f600454905090565b61043f610b89565b505f60065f8581526020019081526020015f2090505f815f015403610499576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104909061140e565b60405180910390fd5b806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636c0f1f588585856040518463ffffffff1660e01b81526004016104f993929190611496565b5f604051808303815f87803b158015610510575f80fd5b505af1158015610522573d5f803e3d5ffd5b5050505050505050565b610534610b0b565b61053d5f610cb6565b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61056e610b0b565b5f600160146101000a81548160ff021916908315150217905550565b5f600160169054906101000a900460ff161590508080156105bc575060018060159054906101000a900460ff1660ff16105b806105ea57506105cb30610d77565b1580156105e9575060018060159054906101000a900460ff1660ff16145b5b610629576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161062090611549565b60405180910390fd5b60018060156101000a81548160ff021916908360ff16021790555080156106655760018060166101000a81548160ff0219169083151502179055505b61066e826102bf565b5f60058190555080156106ce575f600160166101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516106c591906115b5565b60405180910390a15b5050565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f60045460035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e846107446106d2565b6040518363ffffffff1660e01b81526004016107619291906115ce565b602060405180830381865afa15801561077c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107a09190611609565b10156107e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107d89061167e565b60405180910390fd5b60016005546107f091906116c9565b6005819055506040518060a0016040528060055481526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481525060065f60055481526020019081526020015f205f820151815f01556020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600301908161090791906118ed565b506080820151816004019080519060200190610924929190610da0565b509050507faa171d38d08c39622dbd189ddd0a5c0cf3ea10eeda0c7d4c2c75d13fcc4fe14a60055433848787604051610961959493929190611abf565b60405180910390a160055490509392505050565b60055481565b610983610b0b565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036109f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109e890611b8e565b60405180910390fd5b6109fa81610cb6565b50565b610a05610b0b565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610a73576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a6a90611c1c565b60405180910390fd5b600160149054906101000a900460ff16610ac2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ab990611caa565b60405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60045481565b610b13610d99565b73ffffffffffffffffffffffffffffffffffffffff16610b3161053f565b73ffffffffffffffffffffffffffffffffffffffff1614610b87576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b7e90611d12565b60405180910390fd5b565b5f8073ffffffffffffffffffffffffffffffffffffffff1660015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610c19576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1090611c1c565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610c59610d99565b73ffffffffffffffffffffffffffffffffffffffff1614610caf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ca690611da0565b60405180910390fd5b6001905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f33905090565b828054828255905f5260205f20908101928215610de6579160200282015b82811115610de5578251829081610dd591906118ed565b5091602001919060010190610dbe565b5b509050610df39190610df7565b5090565b5b80821115610e16575f8181610e0d9190610e1a565b50600101610df8565b5090565b508054610e2690611729565b5f825580601f10610e375750610e54565b601f0160209004905f5260205f2090810190610e539190610e57565b5b50565b5b80821115610e6e575f815f905550600101610e58565b5090565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610e9b82610e72565b9050919050565b610eab81610e91565b82525050565b5f602082019050610ec45f830184610ea2565b92915050565b5f604051905090565b5f80fd5b5f80fd5b610ee481610e91565b8114610eee575f80fd5b50565b5f81359050610eff81610edb565b92915050565b5f60208284031215610f1a57610f19610ed3565b5b5f610f2784828501610ef1565b91505092915050565b5f819050919050565b610f4281610f30565b8114610f4c575f80fd5b50565b5f81359050610f5d81610f39565b92915050565b5f60208284031215610f7857610f77610ed3565b5b5f610f8584828501610f4f565b91505092915050565b610f9781610f30565b82525050565b5f602082019050610fb05f830184610f8e565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61100482610fbe565b810181811067ffffffffffffffff8211171561102357611022610fce565b5b80604052505050565b5f611035610eca565b90506110418282610ffb565b919050565b5f67ffffffffffffffff8211156110605761105f610fce565b5b61106982610fbe565b9050602081019050919050565b828183375f83830152505050565b5f61109661109184611046565b61102c565b9050828152602081018484840111156110b2576110b1610fba565b5b6110bd848285611076565b509392505050565b5f82601f8301126110d9576110d8610fb6565b5b81356110e9848260208601611084565b91505092915050565b5f805f6060848603121561110957611108610ed3565b5b5f61111686828701610f4f565b935050602084013567ffffffffffffffff81111561113757611136610ed7565b5b611143868287016110c5565b925050604084013567ffffffffffffffff81111561116457611163610ed7565b5b611170868287016110c5565b9150509250925092565b5f67ffffffffffffffff82111561119457611193610fce565b5b602082029050602081019050919050565b5f80fd5b5f6111bb6111b68461117a565b61102c565b905080838252602082019050602084028301858111156111de576111dd6111a5565b5b835b8181101561122557803567ffffffffffffffff81111561120357611202610fb6565b5b80860161121089826110c5565b855260208501945050506020810190506111e0565b5050509392505050565b5f82601f83011261124357611242610fb6565b5b81356112538482602086016111a9565b91505092915050565b5f805f6060848603121561127357611272610ed3565b5b5f84013567ffffffffffffffff8111156112905761128f610ed7565b5b61129c868287016110c5565b935050602084013567ffffffffffffffff8111156112bd576112bc610ed7565b5b6112c98682870161122f565b92505060406112da86828701610ef1565b9150509250925092565b5f82825260208201905092915050565b7f546f6b656e2061646472657373000000000000000000000000000000000000005f82015250565b5f611328600d836112e4565b9150611333826112f4565b602082019050919050565b5f6020820190508181035f8301526113558161131c565b9050919050565b7f4d696e206465706f7369740000000000000000000000000000000000000000005f82015250565b5f611390600b836112e4565b915061139b8261135c565b602082019050919050565b5f6020820190508181035f8301526113bd81611384565b9050919050565b7f4a6f62206e6f7420666f756e64000000000000000000000000000000000000005f82015250565b5f6113f8600d836112e4565b9150611403826113c4565b602082019050919050565b5f6020820190508181035f830152611425816113ec565b9050919050565b5f81519050919050565b5f5b83811015611453578082015181840152602081019050611438565b5f8484015250505050565b5f6114688261142c565b61147281856112e4565b9350611482818560208601611436565b61148b81610fbe565b840191505092915050565b5f6060820190506114a95f830186610f8e565b81810360208301526114bb818561145e565b905081810360408301526114cf818461145e565b9050949350505050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f611533602e836112e4565b915061153e826114d9565b604082019050919050565b5f6020820190508181035f83015261156081611527565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f61159f61159a61159584611567565b61157c565b611570565b9050919050565b6115af81611585565b82525050565b5f6020820190506115c85f8301846115a6565b92915050565b5f6040820190506115e15f830185610ea2565b6115ee6020830184610ea2565b9392505050565b5f8151905061160381610f39565b92915050565b5f6020828403121561161e5761161d610ed3565b5b5f61162b848285016115f5565b91505092915050565b7f546f6b656e20616c6c6f77616e6365206e6f7420656e6f7567680000000000005f82015250565b5f611668601a836112e4565b915061167382611634565b602082019050919050565b5f6020820190508181035f8301526116958161165c565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6116d382610f30565b91506116de83610f30565b92508282019050808211156116f6576116f561169c565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061174057607f821691505b602082108103611753576117526116fc565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026117b57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261177a565b6117bf868361177a565b95508019841693508086168417925050509392505050565b5f6117f16117ec6117e784610f30565b61157c565b610f30565b9050919050565b5f819050919050565b61180a836117d7565b61181e611816826117f8565b848454611786565b825550505050565b5f90565b611832611826565b61183d818484611801565b505050565b5b81811015611860576118555f8261182a565b600181019050611843565b5050565b601f8211156118a55761187681611759565b61187f8461176b565b8101602085101561188e578190505b6118a261189a8561176b565b830182611842565b50505b505050565b5f82821c905092915050565b5f6118c55f19846008026118aa565b1980831691505092915050565b5f6118dd83836118b6565b9150826002028217905092915050565b6118f68261142c565b67ffffffffffffffff81111561190f5761190e610fce565b5b6119198254611729565b611924828285611864565b5f60209050601f831160018114611955575f8415611943578287015190505b61194d85826118d2565b8655506119b4565b601f19841661196386611759565b5f5b8281101561198a57848901518255600182019150602085019450602081019050611965565b868310156119a757848901516119a3601f8916826118b6565b8355505b6001600288020188555050505b505050505050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f82825260208201905092915050565b5f6119ff8261142c565b611a0981856119e5565b9350611a19818560208601611436565b611a2281610fbe565b840191505092915050565b5f611a3883836119f5565b905092915050565b5f602082019050919050565b5f611a56826119bc565b611a6081856119c6565b935083602082028501611a72856119d6565b805f5b85811015611aad5784840389528151611a8e8582611a2d565b9450611a9983611a40565b925060208a01995050600181019050611a75565b50829750879550505050505092915050565b5f60a082019050611ad25f830188610f8e565b611adf6020830187610ea2565b611aec6040830186610ea2565b8181036060830152611afe818561145e565b90508181036080830152611b128184611a4c565b90509695505050505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f611b786026836112e4565b9150611b8382611b1e565b604082019050919050565b5f6020820190508181035f830152611ba581611b6c565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a20436f6e74726f6c6c65722061645f8201527f6472657373206d75737420626520646566696e65640000000000000000000000602082015250565b5f611c066035836112e4565b9150611c1182611bac565b604082019050919050565b5f6020820190508181035f830152611c3381611bfa565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a2063616e4368616e6765436f6e745f8201527f726f6c6c6572416464726573732069732064697361626c656400000000000000602082015250565b5f611c946039836112e4565b9150611c9f82611c3a565b604082019050919050565b5f6020820190508181035f830152611cc181611c88565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f611cfc6020836112e4565b9150611d0782611cc8565b602082019050919050565b5f6020820190508181035f830152611d2981611cf0565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a204f6e6c792074686520636f6e745f8201527f726f6c6c65722063616e2063616c6c2074686973206d6574686f640000000000602082015250565b5f611d8a603b836112e4565b9150611d9582611d30565b604082019050919050565b5f6020820190508181035f830152611db781611d7e565b905091905056fea2646970667358221220830d679c4d8605ce6c32f50d1132366330757ab9701b1808f3e6dfe153bf885464736f6c63430008150033",
}

// JobcreatorABI is the input ABI used to generate the binding from.
// Deprecated: Use JobcreatorMetaData.ABI instead.
var JobcreatorABI = JobcreatorMetaData.ABI

// JobcreatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use JobcreatorMetaData.Bin instead.
var JobcreatorBin = JobcreatorMetaData.Bin

// DeployJobcreator deploys a new Ethereum contract, binding an instance of Jobcreator to it.
func DeployJobcreator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Jobcreator, error) {
	parsed, err := JobcreatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(JobcreatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Jobcreator{JobcreatorCaller: JobcreatorCaller{contract: contract}, JobcreatorTransactor: JobcreatorTransactor{contract: contract}, JobcreatorFilterer: JobcreatorFilterer{contract: contract}}, nil
}

// Jobcreator is an auto generated Go binding around an Ethereum contract.
type Jobcreator struct {
	JobcreatorCaller     // Read-only binding to the contract
	JobcreatorTransactor // Write-only binding to the contract
	JobcreatorFilterer   // Log filterer for contract events
}

// JobcreatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type JobcreatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobcreatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JobcreatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobcreatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JobcreatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobcreatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JobcreatorSession struct {
	Contract     *Jobcreator       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JobcreatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JobcreatorCallerSession struct {
	Contract *JobcreatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// JobcreatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JobcreatorTransactorSession struct {
	Contract     *JobcreatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// JobcreatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type JobcreatorRaw struct {
	Contract *Jobcreator // Generic contract binding to access the raw methods on
}

// JobcreatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JobcreatorCallerRaw struct {
	Contract *JobcreatorCaller // Generic read-only contract binding to access the raw methods on
}

// JobcreatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JobcreatorTransactorRaw struct {
	Contract *JobcreatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJobcreator creates a new instance of Jobcreator, bound to a specific deployed contract.
func NewJobcreator(address common.Address, backend bind.ContractBackend) (*Jobcreator, error) {
	contract, err := bindJobcreator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Jobcreator{JobcreatorCaller: JobcreatorCaller{contract: contract}, JobcreatorTransactor: JobcreatorTransactor{contract: contract}, JobcreatorFilterer: JobcreatorFilterer{contract: contract}}, nil
}

// NewJobcreatorCaller creates a new read-only instance of Jobcreator, bound to a specific deployed contract.
func NewJobcreatorCaller(address common.Address, caller bind.ContractCaller) (*JobcreatorCaller, error) {
	contract, err := bindJobcreator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JobcreatorCaller{contract: contract}, nil
}

// NewJobcreatorTransactor creates a new write-only instance of Jobcreator, bound to a specific deployed contract.
func NewJobcreatorTransactor(address common.Address, transactor bind.ContractTransactor) (*JobcreatorTransactor, error) {
	contract, err := bindJobcreator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JobcreatorTransactor{contract: contract}, nil
}

// NewJobcreatorFilterer creates a new log filterer instance of Jobcreator, bound to a specific deployed contract.
func NewJobcreatorFilterer(address common.Address, filterer bind.ContractFilterer) (*JobcreatorFilterer, error) {
	contract, err := bindJobcreator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JobcreatorFilterer{contract: contract}, nil
}

// bindJobcreator binds a generic wrapper to an already deployed contract.
func bindJobcreator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := JobcreatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Jobcreator *JobcreatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Jobcreator.Contract.JobcreatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Jobcreator *JobcreatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jobcreator.Contract.JobcreatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Jobcreator *JobcreatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Jobcreator.Contract.JobcreatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Jobcreator *JobcreatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Jobcreator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Jobcreator *JobcreatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jobcreator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Jobcreator *JobcreatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Jobcreator.Contract.contract.Transact(opts, method, params...)
}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Jobcreator *JobcreatorCaller) GetControllerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Jobcreator.contract.Call(opts, &out, "getControllerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Jobcreator *JobcreatorSession) GetControllerAddress() (common.Address, error) {
	return _Jobcreator.Contract.GetControllerAddress(&_Jobcreator.CallOpts)
}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Jobcreator *JobcreatorCallerSession) GetControllerAddress() (common.Address, error) {
	return _Jobcreator.Contract.GetControllerAddress(&_Jobcreator.CallOpts)
}

// GetRequiredDeposit is a free data retrieval call binding the contract method 0x58e56db4.
//
// Solidity: function getRequiredDeposit() view returns(uint256)
func (_Jobcreator *JobcreatorCaller) GetRequiredDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Jobcreator.contract.Call(opts, &out, "getRequiredDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequiredDeposit is a free data retrieval call binding the contract method 0x58e56db4.
//
// Solidity: function getRequiredDeposit() view returns(uint256)
func (_Jobcreator *JobcreatorSession) GetRequiredDeposit() (*big.Int, error) {
	return _Jobcreator.Contract.GetRequiredDeposit(&_Jobcreator.CallOpts)
}

// GetRequiredDeposit is a free data retrieval call binding the contract method 0x58e56db4.
//
// Solidity: function getRequiredDeposit() view returns(uint256)
func (_Jobcreator *JobcreatorCallerSession) GetRequiredDeposit() (*big.Int, error) {
	return _Jobcreator.Contract.GetRequiredDeposit(&_Jobcreator.CallOpts)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Jobcreator *JobcreatorCaller) GetTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Jobcreator.contract.Call(opts, &out, "getTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Jobcreator *JobcreatorSession) GetTokenAddress() (common.Address, error) {
	return _Jobcreator.Contract.GetTokenAddress(&_Jobcreator.CallOpts)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Jobcreator *JobcreatorCallerSession) GetTokenAddress() (common.Address, error) {
	return _Jobcreator.Contract.GetTokenAddress(&_Jobcreator.CallOpts)
}

// NextJobID is a free data retrieval call binding the contract method 0xd2a715c0.
//
// Solidity: function nextJobID() view returns(uint256)
func (_Jobcreator *JobcreatorCaller) NextJobID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Jobcreator.contract.Call(opts, &out, "nextJobID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextJobID is a free data retrieval call binding the contract method 0xd2a715c0.
//
// Solidity: function nextJobID() view returns(uint256)
func (_Jobcreator *JobcreatorSession) NextJobID() (*big.Int, error) {
	return _Jobcreator.Contract.NextJobID(&_Jobcreator.CallOpts)
}

// NextJobID is a free data retrieval call binding the contract method 0xd2a715c0.
//
// Solidity: function nextJobID() view returns(uint256)
func (_Jobcreator *JobcreatorCallerSession) NextJobID() (*big.Int, error) {
	return _Jobcreator.Contract.NextJobID(&_Jobcreator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Jobcreator *JobcreatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Jobcreator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Jobcreator *JobcreatorSession) Owner() (common.Address, error) {
	return _Jobcreator.Contract.Owner(&_Jobcreator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Jobcreator *JobcreatorCallerSession) Owner() (common.Address, error) {
	return _Jobcreator.Contract.Owner(&_Jobcreator.CallOpts)
}

// RequiredDeposit is a free data retrieval call binding the contract method 0xfb7cfdd7.
//
// Solidity: function requiredDeposit() view returns(uint256)
func (_Jobcreator *JobcreatorCaller) RequiredDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Jobcreator.contract.Call(opts, &out, "requiredDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredDeposit is a free data retrieval call binding the contract method 0xfb7cfdd7.
//
// Solidity: function requiredDeposit() view returns(uint256)
func (_Jobcreator *JobcreatorSession) RequiredDeposit() (*big.Int, error) {
	return _Jobcreator.Contract.RequiredDeposit(&_Jobcreator.CallOpts)
}

// RequiredDeposit is a free data retrieval call binding the contract method 0xfb7cfdd7.
//
// Solidity: function requiredDeposit() view returns(uint256)
func (_Jobcreator *JobcreatorCallerSession) RequiredDeposit() (*big.Int, error) {
	return _Jobcreator.Contract.RequiredDeposit(&_Jobcreator.CallOpts)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Jobcreator *JobcreatorTransactor) DisableChangeControllerAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "disableChangeControllerAddress")
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Jobcreator *JobcreatorSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _Jobcreator.Contract.DisableChangeControllerAddress(&_Jobcreator.TransactOpts)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Jobcreator *JobcreatorTransactorSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _Jobcreator.Contract.DisableChangeControllerAddress(&_Jobcreator.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tokenAddress) returns()
func (_Jobcreator *JobcreatorTransactor) Initialize(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "initialize", _tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tokenAddress) returns()
func (_Jobcreator *JobcreatorSession) Initialize(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.Initialize(&_Jobcreator.TransactOpts, _tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tokenAddress) returns()
func (_Jobcreator *JobcreatorTransactorSession) Initialize(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.Initialize(&_Jobcreator.TransactOpts, _tokenAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Jobcreator *JobcreatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Jobcreator *JobcreatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Jobcreator.Contract.RenounceOwnership(&_Jobcreator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Jobcreator *JobcreatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Jobcreator.Contract.RenounceOwnership(&_Jobcreator.TransactOpts)
}

// RunJob is a paid mutator transaction binding the contract method 0xc75555fa.
//
// Solidity: function runJob(string module, string[] inputs, address payee) returns(uint256)
func (_Jobcreator *JobcreatorTransactor) RunJob(opts *bind.TransactOpts, module string, inputs []string, payee common.Address) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "runJob", module, inputs, payee)
}

// RunJob is a paid mutator transaction binding the contract method 0xc75555fa.
//
// Solidity: function runJob(string module, string[] inputs, address payee) returns(uint256)
func (_Jobcreator *JobcreatorSession) RunJob(module string, inputs []string, payee common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.RunJob(&_Jobcreator.TransactOpts, module, inputs, payee)
}

// RunJob is a paid mutator transaction binding the contract method 0xc75555fa.
//
// Solidity: function runJob(string module, string[] inputs, address payee) returns(uint256)
func (_Jobcreator *JobcreatorTransactorSession) RunJob(module string, inputs []string, payee common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.RunJob(&_Jobcreator.TransactOpts, module, inputs, payee)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Jobcreator *JobcreatorTransactor) SetControllerAddress(opts *bind.TransactOpts, _controllerAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "setControllerAddress", _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Jobcreator *JobcreatorSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.SetControllerAddress(&_Jobcreator.TransactOpts, _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Jobcreator *JobcreatorTransactorSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.SetControllerAddress(&_Jobcreator.TransactOpts, _controllerAddress)
}

// SetRequiredDeposit is a paid mutator transaction binding the contract method 0x4c526c76.
//
// Solidity: function setRequiredDeposit(uint256 cost) returns()
func (_Jobcreator *JobcreatorTransactor) SetRequiredDeposit(opts *bind.TransactOpts, cost *big.Int) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "setRequiredDeposit", cost)
}

// SetRequiredDeposit is a paid mutator transaction binding the contract method 0x4c526c76.
//
// Solidity: function setRequiredDeposit(uint256 cost) returns()
func (_Jobcreator *JobcreatorSession) SetRequiredDeposit(cost *big.Int) (*types.Transaction, error) {
	return _Jobcreator.Contract.SetRequiredDeposit(&_Jobcreator.TransactOpts, cost)
}

// SetRequiredDeposit is a paid mutator transaction binding the contract method 0x4c526c76.
//
// Solidity: function setRequiredDeposit(uint256 cost) returns()
func (_Jobcreator *JobcreatorTransactorSession) SetRequiredDeposit(cost *big.Int) (*types.Transaction, error) {
	return _Jobcreator.Contract.SetRequiredDeposit(&_Jobcreator.TransactOpts, cost)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_Jobcreator *JobcreatorTransactor) SetTokenAddress(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "setTokenAddress", _tokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_Jobcreator *JobcreatorSession) SetTokenAddress(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.SetTokenAddress(&_Jobcreator.TransactOpts, _tokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_Jobcreator *JobcreatorTransactorSession) SetTokenAddress(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.SetTokenAddress(&_Jobcreator.TransactOpts, _tokenAddress)
}

// SubmitResults is a paid mutator transaction binding the contract method 0x6c0f1f58.
//
// Solidity: function submitResults(uint256 id, string dealId, string dataId) returns()
func (_Jobcreator *JobcreatorTransactor) SubmitResults(opts *bind.TransactOpts, id *big.Int, dealId string, dataId string) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "submitResults", id, dealId, dataId)
}

// SubmitResults is a paid mutator transaction binding the contract method 0x6c0f1f58.
//
// Solidity: function submitResults(uint256 id, string dealId, string dataId) returns()
func (_Jobcreator *JobcreatorSession) SubmitResults(id *big.Int, dealId string, dataId string) (*types.Transaction, error) {
	return _Jobcreator.Contract.SubmitResults(&_Jobcreator.TransactOpts, id, dealId, dataId)
}

// SubmitResults is a paid mutator transaction binding the contract method 0x6c0f1f58.
//
// Solidity: function submitResults(uint256 id, string dealId, string dataId) returns()
func (_Jobcreator *JobcreatorTransactorSession) SubmitResults(id *big.Int, dealId string, dataId string) (*types.Transaction, error) {
	return _Jobcreator.Contract.SubmitResults(&_Jobcreator.TransactOpts, id, dealId, dataId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Jobcreator *JobcreatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Jobcreator *JobcreatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.TransferOwnership(&_Jobcreator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Jobcreator *JobcreatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.TransferOwnership(&_Jobcreator.TransactOpts, newOwner)
}

// JobcreatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Jobcreator contract.
type JobcreatorInitializedIterator struct {
	Event *JobcreatorInitialized // Event containing the contract specifics and raw log

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
func (it *JobcreatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobcreatorInitialized)
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
		it.Event = new(JobcreatorInitialized)
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
func (it *JobcreatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobcreatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobcreatorInitialized represents a Initialized event raised by the Jobcreator contract.
type JobcreatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Jobcreator *JobcreatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*JobcreatorInitializedIterator, error) {

	logs, sub, err := _Jobcreator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &JobcreatorInitializedIterator{contract: _Jobcreator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Jobcreator *JobcreatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *JobcreatorInitialized) (event.Subscription, error) {

	logs, sub, err := _Jobcreator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobcreatorInitialized)
				if err := _Jobcreator.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Jobcreator *JobcreatorFilterer) ParseInitialized(log types.Log) (*JobcreatorInitialized, error) {
	event := new(JobcreatorInitialized)
	if err := _Jobcreator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobcreatorJobAddedIterator is returned from FilterJobAdded and is used to iterate over the raw logs and unpacked data for JobAdded events raised by the Jobcreator contract.
type JobcreatorJobAddedIterator struct {
	Event *JobcreatorJobAdded // Event containing the contract specifics and raw log

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
func (it *JobcreatorJobAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobcreatorJobAdded)
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
		it.Event = new(JobcreatorJobAdded)
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
func (it *JobcreatorJobAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobcreatorJobAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobcreatorJobAdded represents a JobAdded event raised by the Jobcreator contract.
type JobcreatorJobAdded struct {
	Id              *big.Int
	CallingContract common.Address
	Payee           common.Address
	Module          string
	Inputs          []string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterJobAdded is a free log retrieval operation binding the contract event 0xaa171d38d08c39622dbd189ddd0a5c0cf3ea10eeda0c7d4c2c75d13fcc4fe14a.
//
// Solidity: event JobAdded(uint256 id, address calling_contract, address payee, string module, string[] inputs)
func (_Jobcreator *JobcreatorFilterer) FilterJobAdded(opts *bind.FilterOpts) (*JobcreatorJobAddedIterator, error) {

	logs, sub, err := _Jobcreator.contract.FilterLogs(opts, "JobAdded")
	if err != nil {
		return nil, err
	}
	return &JobcreatorJobAddedIterator{contract: _Jobcreator.contract, event: "JobAdded", logs: logs, sub: sub}, nil
}

// WatchJobAdded is a free log subscription operation binding the contract event 0xaa171d38d08c39622dbd189ddd0a5c0cf3ea10eeda0c7d4c2c75d13fcc4fe14a.
//
// Solidity: event JobAdded(uint256 id, address calling_contract, address payee, string module, string[] inputs)
func (_Jobcreator *JobcreatorFilterer) WatchJobAdded(opts *bind.WatchOpts, sink chan<- *JobcreatorJobAdded) (event.Subscription, error) {

	logs, sub, err := _Jobcreator.contract.WatchLogs(opts, "JobAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobcreatorJobAdded)
				if err := _Jobcreator.contract.UnpackLog(event, "JobAdded", log); err != nil {
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

// ParseJobAdded is a log parse operation binding the contract event 0xaa171d38d08c39622dbd189ddd0a5c0cf3ea10eeda0c7d4c2c75d13fcc4fe14a.
//
// Solidity: event JobAdded(uint256 id, address calling_contract, address payee, string module, string[] inputs)
func (_Jobcreator *JobcreatorFilterer) ParseJobAdded(log types.Log) (*JobcreatorJobAdded, error) {
	event := new(JobcreatorJobAdded)
	if err := _Jobcreator.contract.UnpackLog(event, "JobAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobcreatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Jobcreator contract.
type JobcreatorOwnershipTransferredIterator struct {
	Event *JobcreatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *JobcreatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobcreatorOwnershipTransferred)
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
		it.Event = new(JobcreatorOwnershipTransferred)
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
func (it *JobcreatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobcreatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobcreatorOwnershipTransferred represents a OwnershipTransferred event raised by the Jobcreator contract.
type JobcreatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Jobcreator *JobcreatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*JobcreatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Jobcreator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &JobcreatorOwnershipTransferredIterator{contract: _Jobcreator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Jobcreator *JobcreatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *JobcreatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Jobcreator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobcreatorOwnershipTransferred)
				if err := _Jobcreator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Jobcreator *JobcreatorFilterer) ParseOwnershipTransferred(log types.Log) (*JobcreatorOwnershipTransferred, error) {
	event := new(JobcreatorOwnershipTransferred)
	if err := _Jobcreator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
