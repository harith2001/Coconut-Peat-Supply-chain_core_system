// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tracking

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

// TrackingTypeShipment is an auto generated low-level Go binding around an user-defined struct.
type TrackingTypeShipment struct {
	Sender       common.Address
	Receiver     common.Address
	PickupTime   *big.Int
	DeliveryTime *big.Int
	Distance     *big.Int
	Price        *big.Int
	Status       uint8
	IsPaid       bool
}

// TrackingMetaData contains all meta data concerning the Tracking contract.
var TrackingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pickupTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"distance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"ShipmentCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deliveryTime\",\"type\":\"uint256\"}],\"name\":\"ShipmentDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pickupTime\",\"type\":\"uint256\"}],\"name\":\"ShipmentInTransit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ShipmentPaid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"completeShipment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pickupTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_distance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"createShipment\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllTransactions\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pickupTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deliveryTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"distance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"enumTracking.ShipmentStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isPaid\",\"type\":\"bool\"}],\"internalType\":\"structTracking.TypeShipment[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getShipment\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"enumTracking.ShipmentStatus\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"getShipmentsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shipmentCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"shipments\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pickupTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deliveryTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"distance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"enumTracking.ShipmentStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isPaid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"startShipment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506000600181905550611be6806100286000396000f3fe60806040526004361061007b5760003560e01c80637ff5265d1161004e5780637ff5265d146101555780638365da541461017e578063aace20b7146101c2578063cbc4b181146101ed5761007b565b8063179a878c1461008057806327506f53146100a957806341fdc195146100d45780635667e02f14610111575b600080fd5b34801561008c57600080fd5b506100a760048036038101906100a29190611499565b610209565b005b3480156100b557600080fd5b506100be6106c1565b6040516100cb91906118fc565b60405180910390f35b3480156100e057600080fd5b506100fb60048036038101906100f69190611470565b610897565b60405161010891906119be565b60405180910390f35b34801561011d57600080fd5b50610138600480360381019061013391906114e8565b6108e2565b60405161014c98979695949392919061187e565b60405180910390f35b34801561016157600080fd5b5061017c60048036038101906101779190611499565b6109a1565b005b34801561018a57600080fd5b506101a560048036038101906101a091906114e8565b610ce1565b6040516101b998979695949392919061187e565b60405180910390f35b3480156101ce57600080fd5b506101d7610f3d565b6040516101e491906119be565b60405180910390f35b61020760048036038101906102029190611524565b610f43565b005b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208281548110610281577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906007020190506000600283815481106102cd577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906007020190508373ffffffffffffffffffffffffffffffffffffffff168260010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461036f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103669061197e565b60405180910390fd5b600160028111156103a9577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8260060160009054906101000a900460ff1660028111156103f3577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14610433576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161042a9061195e565b60405180910390fd5b8160060160019054906101000a900460ff1615610485576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047c9061191e565b60405180910390fd5b60028260060160006101000a81548160ff021916908360028111156104d3577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555060028160060160006101000a81548160ff02191690836002811115610526577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b02179055504281600301819055504282600301819055506000826005015490508260010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156105b0573d6000803e3d6000fd5b5060018360060160016101000a81548160ff02191690831515021790555060018260060160016101000a81548160ff0219169083151502179055508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167f3746ccceb54f44c9b0ee0120e596168d07ec9f6ca303cfdd9810d3b2e2686a2e856003015460405161064c91906119be565b60405180910390a38473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167f2b6f2473aebafaf8d0889669490135928ec4454d2e29a103fd39693bbca6aef1836040516106b191906119be565b60405180910390a3505050505050565b60606002805480602002602001604051908101604052809291908181526020016000905b8282101561088e5783829060005260206000209060070201604051806101000160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015481526020016003820154815260200160048201548152602001600582015481526020016006820160009054906101000a900460ff166002811115610828577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811115610860577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81526020016006820160019054906101000a900460ff161515151581525050815260200190600101906106e5565b50505050905090565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490509050919050565b600060205281600052604060002081815481106108fe57600080fd5b9060005260206000209060070201600091509150508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030154908060040154908060050154908060060160009054906101000a900460ff16908060060160019054906101000a900460ff16905088565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208281548110610a19577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002090600702019050600060028381548110610a65577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906007020190508373ffffffffffffffffffffffffffffffffffffffff168260010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610b07576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610afe9061197e565b60405180910390fd5b60006002811115610b41577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8260060160009054906101000a900460ff166002811115610b8b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b14610bcb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bc29061199e565b60405180910390fd5b60018260060160006101000a81548160ff02191690836002811115610c19577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555060018160060160006101000a81548160ff02191690836002811115610c6c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b02179055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167faafab9f94340179ff9cefd29621e1772a4b8efab8023163838f1b0a8f6ce5df58460020154604051610cd291906119be565b60405180910390a35050505050565b60008060008060008060008060008060008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208a81548110610d65577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000209060070201604051806101000160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015481526020016003820154815260200160048201548152602001600582015481526020016006820160009054906101000a900460ff166002811115610e9e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811115610ed6577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81526020016006820160019054906101000a900460ff1615151515815250509050806000015181602001518260400151836060015184608001518560a001518660c001518760e0015198509850985098509850985098509850509295985092959890939650565b60015481565b803414610f85576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f7c9061193e565b60405180910390fd5b60006040518061010001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff168152602001858152602001600081526020018481526020018381526020016000600281111561101e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81526020016000151581525090506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081908060018154018082558091505060019003906000526020600020906007020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020155606082015181600301556080820151816004015560a0820151816005015560c08201518160060160006101000a81548160ff0219169083600281111561119b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555060e08201518160060160016101000a81548160ff0219169083151502179055505050600160008154809291906111d590611ac7565b919050555060026040518061010001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1681526020018681526020016000815260200185815260200184815260200160006002811115611273577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815260200160001515815250908060018154018082558091505060019003906000526020600020906007020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020155606082015181600301556080820151816004015560a0820151816005015560c08201518160060160006101000a81548160ff021916908360028111156113af577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555060e08201518160060160016101000a81548160ff02191690831515021790555050508473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f5001b34ecf2db5b7bbeddcbbe761ec9ecadcc7638018b98ce3e24995b0e946b6868686604051611437939291906119d9565b60405180910390a35050505050565b60008135905061145581611b82565b92915050565b60008135905061146a81611b99565b92915050565b60006020828403121561148257600080fd5b600061149084828501611446565b91505092915050565b6000806000606084860312156114ae57600080fd5b60006114bc86828701611446565b93505060206114cd86828701611446565b92505060406114de8682870161145b565b9150509250925092565b600080604083850312156114fb57600080fd5b600061150985828601611446565b925050602061151a8582860161145b565b9150509250929050565b6000806000806080858703121561153a57600080fd5b600061154887828801611446565b94505060206115598782880161145b565b935050604061156a8782880161145b565b925050606061157b8782880161145b565b91505092959194509250565b600061159383836117be565b6101008301905092915050565b6115a981611a5a565b82525050565b6115b881611a5a565b82525050565b60006115c982611a20565b6115d38185611a38565b93506115de83611a10565b8060005b8381101561160f5781516115f68882611587565b975061160183611a2b565b9250506001810190506115e2565b5085935050505092915050565b61162581611a6c565b82525050565b61163481611a6c565b82525050565b61164381611ab5565b82525050565b61165281611ab5565b82525050565b6000611665601683611a49565b91507f536869706d656e7420616c726561647920706169642e000000000000000000006000830152602082019050919050565b60006116a5602483611a49565b91507f5061796d656e7420616d6f756e74206d757374206d617463682074686520707260008301527f6963652e000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b600061170b601883611a49565b91507f536869706d656e74206e6f7420696e207472616e7369742e00000000000000006000830152602082019050919050565b600061174b601183611a49565b91507f496e76616c69642072656365697665722e0000000000000000000000000000006000830152602082019050919050565b600061178b601c83611a49565b91507f536869706d656e7420616c726561647920696e207472616e7369742e000000006000830152602082019050919050565b610100820160008201516117d560008501826115a0565b5060208201516117e860208501826115a0565b5060408201516117fb6040850182611860565b50606082015161180e6060850182611860565b5060808201516118216080850182611860565b5060a082015161183460a0850182611860565b5060c082015161184760c085018261163a565b5060e082015161185a60e085018261161c565b50505050565b61186981611aab565b82525050565b61187881611aab565b82525050565b600061010082019050611894600083018b6115af565b6118a1602083018a6115af565b6118ae604083018961186f565b6118bb606083018861186f565b6118c8608083018761186f565b6118d560a083018661186f565b6118e260c0830185611649565b6118ef60e083018461162b565b9998505050505050505050565b6000602082019050818103600083015261191681846115be565b905092915050565b6000602082019050818103600083015261193781611658565b9050919050565b6000602082019050818103600083015261195781611698565b9050919050565b60006020820190508181036000830152611977816116fe565b9050919050565b600060208201905081810360008301526119978161173e565b9050919050565b600060208201905081810360008301526119b78161177e565b9050919050565b60006020820190506119d3600083018461186f565b92915050565b60006060820190506119ee600083018661186f565b6119fb602083018561186f565b611a08604083018461186f565b949350505050565b6000819050602082019050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b6000611a6582611a8b565b9050919050565b60008115159050919050565b6000819050611a8682611b6e565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611ac082611a78565b9050919050565b6000611ad282611aab565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611b0557611b04611b10565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60038110611b7f57611b7e611b3f565b5b50565b611b8b81611a5a565b8114611b9657600080fd5b50565b611ba281611aab565b8114611bad57600080fd5b5056fea264697066735822122055829a0ce9a1bc7cb30c8633d07811a8b67fc706edbe85462cfc6881698a073364736f6c63430008000033",
}

// TrackingABI is the input ABI used to generate the binding from.
// Deprecated: Use TrackingMetaData.ABI instead.
var TrackingABI = TrackingMetaData.ABI

// TrackingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TrackingMetaData.Bin instead.
var TrackingBin = TrackingMetaData.Bin

// DeployTracking deploys a new Ethereum contract, binding an instance of Tracking to it.
func DeployTracking(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tracking, error) {
	parsed, err := TrackingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TrackingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tracking{TrackingCaller: TrackingCaller{contract: contract}, TrackingTransactor: TrackingTransactor{contract: contract}, TrackingFilterer: TrackingFilterer{contract: contract}}, nil
}

// Tracking is an auto generated Go binding around an Ethereum contract.
type Tracking struct {
	TrackingCaller     // Read-only binding to the contract
	TrackingTransactor // Write-only binding to the contract
	TrackingFilterer   // Log filterer for contract events
}

// TrackingCaller is an auto generated read-only Go binding around an Ethereum contract.
type TrackingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrackingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TrackingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrackingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TrackingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrackingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TrackingSession struct {
	Contract     *Tracking         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TrackingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TrackingCallerSession struct {
	Contract *TrackingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TrackingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TrackingTransactorSession struct {
	Contract     *TrackingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TrackingRaw is an auto generated low-level Go binding around an Ethereum contract.
type TrackingRaw struct {
	Contract *Tracking // Generic contract binding to access the raw methods on
}

// TrackingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TrackingCallerRaw struct {
	Contract *TrackingCaller // Generic read-only contract binding to access the raw methods on
}

// TrackingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TrackingTransactorRaw struct {
	Contract *TrackingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTracking creates a new instance of Tracking, bound to a specific deployed contract.
func NewTracking(address common.Address, backend bind.ContractBackend) (*Tracking, error) {
	contract, err := bindTracking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tracking{TrackingCaller: TrackingCaller{contract: contract}, TrackingTransactor: TrackingTransactor{contract: contract}, TrackingFilterer: TrackingFilterer{contract: contract}}, nil
}

// NewTrackingCaller creates a new read-only instance of Tracking, bound to a specific deployed contract.
func NewTrackingCaller(address common.Address, caller bind.ContractCaller) (*TrackingCaller, error) {
	contract, err := bindTracking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TrackingCaller{contract: contract}, nil
}

// NewTrackingTransactor creates a new write-only instance of Tracking, bound to a specific deployed contract.
func NewTrackingTransactor(address common.Address, transactor bind.ContractTransactor) (*TrackingTransactor, error) {
	contract, err := bindTracking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TrackingTransactor{contract: contract}, nil
}

// NewTrackingFilterer creates a new log filterer instance of Tracking, bound to a specific deployed contract.
func NewTrackingFilterer(address common.Address, filterer bind.ContractFilterer) (*TrackingFilterer, error) {
	contract, err := bindTracking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TrackingFilterer{contract: contract}, nil
}

// bindTracking binds a generic wrapper to an already deployed contract.
func bindTracking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TrackingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tracking *TrackingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tracking.Contract.TrackingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tracking *TrackingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tracking.Contract.TrackingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tracking *TrackingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tracking.Contract.TrackingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tracking *TrackingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tracking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tracking *TrackingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tracking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tracking *TrackingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tracking.Contract.contract.Transact(opts, method, params...)
}

// GetAllTransactions is a free data retrieval call binding the contract method 0x27506f53.
//
// Solidity: function getAllTransactions() view returns((address,address,uint256,uint256,uint256,uint256,uint8,bool)[])
func (_Tracking *TrackingCaller) GetAllTransactions(opts *bind.CallOpts) ([]TrackingTypeShipment, error) {
	var out []interface{}
	err := _Tracking.contract.Call(opts, &out, "getAllTransactions")

	if err != nil {
		return *new([]TrackingTypeShipment), err
	}

	out0 := *abi.ConvertType(out[0], new([]TrackingTypeShipment)).(*[]TrackingTypeShipment)

	return out0, err

}

// GetAllTransactions is a free data retrieval call binding the contract method 0x27506f53.
//
// Solidity: function getAllTransactions() view returns((address,address,uint256,uint256,uint256,uint256,uint8,bool)[])
func (_Tracking *TrackingSession) GetAllTransactions() ([]TrackingTypeShipment, error) {
	return _Tracking.Contract.GetAllTransactions(&_Tracking.CallOpts)
}

// GetAllTransactions is a free data retrieval call binding the contract method 0x27506f53.
//
// Solidity: function getAllTransactions() view returns((address,address,uint256,uint256,uint256,uint256,uint8,bool)[])
func (_Tracking *TrackingCallerSession) GetAllTransactions() ([]TrackingTypeShipment, error) {
	return _Tracking.Contract.GetAllTransactions(&_Tracking.CallOpts)
}

// GetShipment is a free data retrieval call binding the contract method 0x8365da54.
//
// Solidity: function getShipment(address _sender, uint256 _index) view returns(address, address, uint256, uint256, uint256, uint256, uint8, bool)
func (_Tracking *TrackingCaller) GetShipment(opts *bind.CallOpts, _sender common.Address, _index *big.Int) (common.Address, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, uint8, bool, error) {
	var out []interface{}
	err := _Tracking.contract.Call(opts, &out, "getShipment", _sender, _index)

	if err != nil {
		return *new(common.Address), *new(common.Address), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(uint8), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(uint8)).(*uint8)
	out7 := *abi.ConvertType(out[7], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, out5, out6, out7, err

}

// GetShipment is a free data retrieval call binding the contract method 0x8365da54.
//
// Solidity: function getShipment(address _sender, uint256 _index) view returns(address, address, uint256, uint256, uint256, uint256, uint8, bool)
func (_Tracking *TrackingSession) GetShipment(_sender common.Address, _index *big.Int) (common.Address, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, uint8, bool, error) {
	return _Tracking.Contract.GetShipment(&_Tracking.CallOpts, _sender, _index)
}

// GetShipment is a free data retrieval call binding the contract method 0x8365da54.
//
// Solidity: function getShipment(address _sender, uint256 _index) view returns(address, address, uint256, uint256, uint256, uint256, uint8, bool)
func (_Tracking *TrackingCallerSession) GetShipment(_sender common.Address, _index *big.Int) (common.Address, common.Address, *big.Int, *big.Int, *big.Int, *big.Int, uint8, bool, error) {
	return _Tracking.Contract.GetShipment(&_Tracking.CallOpts, _sender, _index)
}

// GetShipmentsCount is a free data retrieval call binding the contract method 0x41fdc195.
//
// Solidity: function getShipmentsCount(address _sender) view returns(uint256)
func (_Tracking *TrackingCaller) GetShipmentsCount(opts *bind.CallOpts, _sender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tracking.contract.Call(opts, &out, "getShipmentsCount", _sender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetShipmentsCount is a free data retrieval call binding the contract method 0x41fdc195.
//
// Solidity: function getShipmentsCount(address _sender) view returns(uint256)
func (_Tracking *TrackingSession) GetShipmentsCount(_sender common.Address) (*big.Int, error) {
	return _Tracking.Contract.GetShipmentsCount(&_Tracking.CallOpts, _sender)
}

// GetShipmentsCount is a free data retrieval call binding the contract method 0x41fdc195.
//
// Solidity: function getShipmentsCount(address _sender) view returns(uint256)
func (_Tracking *TrackingCallerSession) GetShipmentsCount(_sender common.Address) (*big.Int, error) {
	return _Tracking.Contract.GetShipmentsCount(&_Tracking.CallOpts, _sender)
}

// ShipmentCount is a free data retrieval call binding the contract method 0xaace20b7.
//
// Solidity: function shipmentCount() view returns(uint256)
func (_Tracking *TrackingCaller) ShipmentCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tracking.contract.Call(opts, &out, "shipmentCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ShipmentCount is a free data retrieval call binding the contract method 0xaace20b7.
//
// Solidity: function shipmentCount() view returns(uint256)
func (_Tracking *TrackingSession) ShipmentCount() (*big.Int, error) {
	return _Tracking.Contract.ShipmentCount(&_Tracking.CallOpts)
}

// ShipmentCount is a free data retrieval call binding the contract method 0xaace20b7.
//
// Solidity: function shipmentCount() view returns(uint256)
func (_Tracking *TrackingCallerSession) ShipmentCount() (*big.Int, error) {
	return _Tracking.Contract.ShipmentCount(&_Tracking.CallOpts)
}

// Shipments is a free data retrieval call binding the contract method 0x5667e02f.
//
// Solidity: function shipments(address , uint256 ) view returns(address sender, address receiver, uint256 pickupTime, uint256 deliveryTime, uint256 distance, uint256 price, uint8 status, bool isPaid)
func (_Tracking *TrackingCaller) Shipments(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Sender       common.Address
	Receiver     common.Address
	PickupTime   *big.Int
	DeliveryTime *big.Int
	Distance     *big.Int
	Price        *big.Int
	Status       uint8
	IsPaid       bool
}, error) {
	var out []interface{}
	err := _Tracking.contract.Call(opts, &out, "shipments", arg0, arg1)

	outstruct := new(struct {
		Sender       common.Address
		Receiver     common.Address
		PickupTime   *big.Int
		DeliveryTime *big.Int
		Distance     *big.Int
		Price        *big.Int
		Status       uint8
		IsPaid       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Sender = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Receiver = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.PickupTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DeliveryTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Distance = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.IsPaid = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// Shipments is a free data retrieval call binding the contract method 0x5667e02f.
//
// Solidity: function shipments(address , uint256 ) view returns(address sender, address receiver, uint256 pickupTime, uint256 deliveryTime, uint256 distance, uint256 price, uint8 status, bool isPaid)
func (_Tracking *TrackingSession) Shipments(arg0 common.Address, arg1 *big.Int) (struct {
	Sender       common.Address
	Receiver     common.Address
	PickupTime   *big.Int
	DeliveryTime *big.Int
	Distance     *big.Int
	Price        *big.Int
	Status       uint8
	IsPaid       bool
}, error) {
	return _Tracking.Contract.Shipments(&_Tracking.CallOpts, arg0, arg1)
}

// Shipments is a free data retrieval call binding the contract method 0x5667e02f.
//
// Solidity: function shipments(address , uint256 ) view returns(address sender, address receiver, uint256 pickupTime, uint256 deliveryTime, uint256 distance, uint256 price, uint8 status, bool isPaid)
func (_Tracking *TrackingCallerSession) Shipments(arg0 common.Address, arg1 *big.Int) (struct {
	Sender       common.Address
	Receiver     common.Address
	PickupTime   *big.Int
	DeliveryTime *big.Int
	Distance     *big.Int
	Price        *big.Int
	Status       uint8
	IsPaid       bool
}, error) {
	return _Tracking.Contract.Shipments(&_Tracking.CallOpts, arg0, arg1)
}

// CompleteShipment is a paid mutator transaction binding the contract method 0x179a878c.
//
// Solidity: function completeShipment(address _sender, address _receiver, uint256 _index) returns()
func (_Tracking *TrackingTransactor) CompleteShipment(opts *bind.TransactOpts, _sender common.Address, _receiver common.Address, _index *big.Int) (*types.Transaction, error) {
	return _Tracking.contract.Transact(opts, "completeShipment", _sender, _receiver, _index)
}

// CompleteShipment is a paid mutator transaction binding the contract method 0x179a878c.
//
// Solidity: function completeShipment(address _sender, address _receiver, uint256 _index) returns()
func (_Tracking *TrackingSession) CompleteShipment(_sender common.Address, _receiver common.Address, _index *big.Int) (*types.Transaction, error) {
	return _Tracking.Contract.CompleteShipment(&_Tracking.TransactOpts, _sender, _receiver, _index)
}

// CompleteShipment is a paid mutator transaction binding the contract method 0x179a878c.
//
// Solidity: function completeShipment(address _sender, address _receiver, uint256 _index) returns()
func (_Tracking *TrackingTransactorSession) CompleteShipment(_sender common.Address, _receiver common.Address, _index *big.Int) (*types.Transaction, error) {
	return _Tracking.Contract.CompleteShipment(&_Tracking.TransactOpts, _sender, _receiver, _index)
}

// CreateShipment is a paid mutator transaction binding the contract method 0xcbc4b181.
//
// Solidity: function createShipment(address _receiver, uint256 _pickupTime, uint256 _distance, uint256 _price) payable returns()
func (_Tracking *TrackingTransactor) CreateShipment(opts *bind.TransactOpts, _receiver common.Address, _pickupTime *big.Int, _distance *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Tracking.contract.Transact(opts, "createShipment", _receiver, _pickupTime, _distance, _price)
}

// CreateShipment is a paid mutator transaction binding the contract method 0xcbc4b181.
//
// Solidity: function createShipment(address _receiver, uint256 _pickupTime, uint256 _distance, uint256 _price) payable returns()
func (_Tracking *TrackingSession) CreateShipment(_receiver common.Address, _pickupTime *big.Int, _distance *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Tracking.Contract.CreateShipment(&_Tracking.TransactOpts, _receiver, _pickupTime, _distance, _price)
}

// CreateShipment is a paid mutator transaction binding the contract method 0xcbc4b181.
//
// Solidity: function createShipment(address _receiver, uint256 _pickupTime, uint256 _distance, uint256 _price) payable returns()
func (_Tracking *TrackingTransactorSession) CreateShipment(_receiver common.Address, _pickupTime *big.Int, _distance *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Tracking.Contract.CreateShipment(&_Tracking.TransactOpts, _receiver, _pickupTime, _distance, _price)
}

// StartShipment is a paid mutator transaction binding the contract method 0x7ff5265d.
//
// Solidity: function startShipment(address _sender, address _receiver, uint256 _index) returns()
func (_Tracking *TrackingTransactor) StartShipment(opts *bind.TransactOpts, _sender common.Address, _receiver common.Address, _index *big.Int) (*types.Transaction, error) {
	return _Tracking.contract.Transact(opts, "startShipment", _sender, _receiver, _index)
}

// StartShipment is a paid mutator transaction binding the contract method 0x7ff5265d.
//
// Solidity: function startShipment(address _sender, address _receiver, uint256 _index) returns()
func (_Tracking *TrackingSession) StartShipment(_sender common.Address, _receiver common.Address, _index *big.Int) (*types.Transaction, error) {
	return _Tracking.Contract.StartShipment(&_Tracking.TransactOpts, _sender, _receiver, _index)
}

// StartShipment is a paid mutator transaction binding the contract method 0x7ff5265d.
//
// Solidity: function startShipment(address _sender, address _receiver, uint256 _index) returns()
func (_Tracking *TrackingTransactorSession) StartShipment(_sender common.Address, _receiver common.Address, _index *big.Int) (*types.Transaction, error) {
	return _Tracking.Contract.StartShipment(&_Tracking.TransactOpts, _sender, _receiver, _index)
}

// TrackingShipmentCreatedIterator is returned from FilterShipmentCreated and is used to iterate over the raw logs and unpacked data for ShipmentCreated events raised by the Tracking contract.
type TrackingShipmentCreatedIterator struct {
	Event *TrackingShipmentCreated // Event containing the contract specifics and raw log

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
func (it *TrackingShipmentCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrackingShipmentCreated)
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
		it.Event = new(TrackingShipmentCreated)
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
func (it *TrackingShipmentCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrackingShipmentCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrackingShipmentCreated represents a ShipmentCreated event raised by the Tracking contract.
type TrackingShipmentCreated struct {
	Sender     common.Address
	Receiver   common.Address
	PickupTime *big.Int
	Distance   *big.Int
	Price      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterShipmentCreated is a free log retrieval operation binding the contract event 0x5001b34ecf2db5b7bbeddcbbe761ec9ecadcc7638018b98ce3e24995b0e946b6.
//
// Solidity: event ShipmentCreated(address indexed sender, address indexed receiver, uint256 pickupTime, uint256 distance, uint256 price)
func (_Tracking *TrackingFilterer) FilterShipmentCreated(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*TrackingShipmentCreatedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Tracking.contract.FilterLogs(opts, "ShipmentCreated", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &TrackingShipmentCreatedIterator{contract: _Tracking.contract, event: "ShipmentCreated", logs: logs, sub: sub}, nil
}

// WatchShipmentCreated is a free log subscription operation binding the contract event 0x5001b34ecf2db5b7bbeddcbbe761ec9ecadcc7638018b98ce3e24995b0e946b6.
//
// Solidity: event ShipmentCreated(address indexed sender, address indexed receiver, uint256 pickupTime, uint256 distance, uint256 price)
func (_Tracking *TrackingFilterer) WatchShipmentCreated(opts *bind.WatchOpts, sink chan<- *TrackingShipmentCreated, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Tracking.contract.WatchLogs(opts, "ShipmentCreated", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrackingShipmentCreated)
				if err := _Tracking.contract.UnpackLog(event, "ShipmentCreated", log); err != nil {
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

// ParseShipmentCreated is a log parse operation binding the contract event 0x5001b34ecf2db5b7bbeddcbbe761ec9ecadcc7638018b98ce3e24995b0e946b6.
//
// Solidity: event ShipmentCreated(address indexed sender, address indexed receiver, uint256 pickupTime, uint256 distance, uint256 price)
func (_Tracking *TrackingFilterer) ParseShipmentCreated(log types.Log) (*TrackingShipmentCreated, error) {
	event := new(TrackingShipmentCreated)
	if err := _Tracking.contract.UnpackLog(event, "ShipmentCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrackingShipmentDeliveredIterator is returned from FilterShipmentDelivered and is used to iterate over the raw logs and unpacked data for ShipmentDelivered events raised by the Tracking contract.
type TrackingShipmentDeliveredIterator struct {
	Event *TrackingShipmentDelivered // Event containing the contract specifics and raw log

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
func (it *TrackingShipmentDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrackingShipmentDelivered)
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
		it.Event = new(TrackingShipmentDelivered)
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
func (it *TrackingShipmentDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrackingShipmentDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrackingShipmentDelivered represents a ShipmentDelivered event raised by the Tracking contract.
type TrackingShipmentDelivered struct {
	Sender       common.Address
	Receiver     common.Address
	DeliveryTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterShipmentDelivered is a free log retrieval operation binding the contract event 0x3746ccceb54f44c9b0ee0120e596168d07ec9f6ca303cfdd9810d3b2e2686a2e.
//
// Solidity: event ShipmentDelivered(address indexed sender, address indexed receiver, uint256 deliveryTime)
func (_Tracking *TrackingFilterer) FilterShipmentDelivered(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*TrackingShipmentDeliveredIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Tracking.contract.FilterLogs(opts, "ShipmentDelivered", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &TrackingShipmentDeliveredIterator{contract: _Tracking.contract, event: "ShipmentDelivered", logs: logs, sub: sub}, nil
}

// WatchShipmentDelivered is a free log subscription operation binding the contract event 0x3746ccceb54f44c9b0ee0120e596168d07ec9f6ca303cfdd9810d3b2e2686a2e.
//
// Solidity: event ShipmentDelivered(address indexed sender, address indexed receiver, uint256 deliveryTime)
func (_Tracking *TrackingFilterer) WatchShipmentDelivered(opts *bind.WatchOpts, sink chan<- *TrackingShipmentDelivered, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Tracking.contract.WatchLogs(opts, "ShipmentDelivered", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrackingShipmentDelivered)
				if err := _Tracking.contract.UnpackLog(event, "ShipmentDelivered", log); err != nil {
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

// ParseShipmentDelivered is a log parse operation binding the contract event 0x3746ccceb54f44c9b0ee0120e596168d07ec9f6ca303cfdd9810d3b2e2686a2e.
//
// Solidity: event ShipmentDelivered(address indexed sender, address indexed receiver, uint256 deliveryTime)
func (_Tracking *TrackingFilterer) ParseShipmentDelivered(log types.Log) (*TrackingShipmentDelivered, error) {
	event := new(TrackingShipmentDelivered)
	if err := _Tracking.contract.UnpackLog(event, "ShipmentDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrackingShipmentInTransitIterator is returned from FilterShipmentInTransit and is used to iterate over the raw logs and unpacked data for ShipmentInTransit events raised by the Tracking contract.
type TrackingShipmentInTransitIterator struct {
	Event *TrackingShipmentInTransit // Event containing the contract specifics and raw log

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
func (it *TrackingShipmentInTransitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrackingShipmentInTransit)
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
		it.Event = new(TrackingShipmentInTransit)
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
func (it *TrackingShipmentInTransitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrackingShipmentInTransitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrackingShipmentInTransit represents a ShipmentInTransit event raised by the Tracking contract.
type TrackingShipmentInTransit struct {
	Sender     common.Address
	Receiver   common.Address
	PickupTime *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterShipmentInTransit is a free log retrieval operation binding the contract event 0xaafab9f94340179ff9cefd29621e1772a4b8efab8023163838f1b0a8f6ce5df5.
//
// Solidity: event ShipmentInTransit(address indexed sender, address indexed receiver, uint256 pickupTime)
func (_Tracking *TrackingFilterer) FilterShipmentInTransit(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*TrackingShipmentInTransitIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Tracking.contract.FilterLogs(opts, "ShipmentInTransit", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &TrackingShipmentInTransitIterator{contract: _Tracking.contract, event: "ShipmentInTransit", logs: logs, sub: sub}, nil
}

// WatchShipmentInTransit is a free log subscription operation binding the contract event 0xaafab9f94340179ff9cefd29621e1772a4b8efab8023163838f1b0a8f6ce5df5.
//
// Solidity: event ShipmentInTransit(address indexed sender, address indexed receiver, uint256 pickupTime)
func (_Tracking *TrackingFilterer) WatchShipmentInTransit(opts *bind.WatchOpts, sink chan<- *TrackingShipmentInTransit, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Tracking.contract.WatchLogs(opts, "ShipmentInTransit", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrackingShipmentInTransit)
				if err := _Tracking.contract.UnpackLog(event, "ShipmentInTransit", log); err != nil {
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

// ParseShipmentInTransit is a log parse operation binding the contract event 0xaafab9f94340179ff9cefd29621e1772a4b8efab8023163838f1b0a8f6ce5df5.
//
// Solidity: event ShipmentInTransit(address indexed sender, address indexed receiver, uint256 pickupTime)
func (_Tracking *TrackingFilterer) ParseShipmentInTransit(log types.Log) (*TrackingShipmentInTransit, error) {
	event := new(TrackingShipmentInTransit)
	if err := _Tracking.contract.UnpackLog(event, "ShipmentInTransit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrackingShipmentPaidIterator is returned from FilterShipmentPaid and is used to iterate over the raw logs and unpacked data for ShipmentPaid events raised by the Tracking contract.
type TrackingShipmentPaidIterator struct {
	Event *TrackingShipmentPaid // Event containing the contract specifics and raw log

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
func (it *TrackingShipmentPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrackingShipmentPaid)
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
		it.Event = new(TrackingShipmentPaid)
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
func (it *TrackingShipmentPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrackingShipmentPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrackingShipmentPaid represents a ShipmentPaid event raised by the Tracking contract.
type TrackingShipmentPaid struct {
	Sender   common.Address
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterShipmentPaid is a free log retrieval operation binding the contract event 0x2b6f2473aebafaf8d0889669490135928ec4454d2e29a103fd39693bbca6aef1.
//
// Solidity: event ShipmentPaid(address indexed sender, address indexed receiver, uint256 amount)
func (_Tracking *TrackingFilterer) FilterShipmentPaid(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*TrackingShipmentPaidIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Tracking.contract.FilterLogs(opts, "ShipmentPaid", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &TrackingShipmentPaidIterator{contract: _Tracking.contract, event: "ShipmentPaid", logs: logs, sub: sub}, nil
}

// WatchShipmentPaid is a free log subscription operation binding the contract event 0x2b6f2473aebafaf8d0889669490135928ec4454d2e29a103fd39693bbca6aef1.
//
// Solidity: event ShipmentPaid(address indexed sender, address indexed receiver, uint256 amount)
func (_Tracking *TrackingFilterer) WatchShipmentPaid(opts *bind.WatchOpts, sink chan<- *TrackingShipmentPaid, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Tracking.contract.WatchLogs(opts, "ShipmentPaid", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrackingShipmentPaid)
				if err := _Tracking.contract.UnpackLog(event, "ShipmentPaid", log); err != nil {
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

// ParseShipmentPaid is a log parse operation binding the contract event 0x2b6f2473aebafaf8d0889669490135928ec4454d2e29a103fd39693bbca6aef1.
//
// Solidity: event ShipmentPaid(address indexed sender, address indexed receiver, uint256 amount)
func (_Tracking *TrackingFilterer) ParseShipmentPaid(log types.Log) (*TrackingShipmentPaid, error) {
	event := new(TrackingShipmentPaid)
	if err := _Tracking.contract.UnpackLog(event, "ShipmentPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
