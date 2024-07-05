package propertymanager

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

// IPropertyInvestmentValue is an auto generated low-level Go binding around an user-defined struct.
type IPropertyInvestmentValue struct {
	PropertyPrice         *big.Int
	TokenPrice            *big.Int
	TokenAmount           *big.Int
	SoldQuantity          *big.Int
	ProjectedAnnualReturn *big.Int
	ProjectedRentalYield  *big.Int
	RentStartDate         uint64
	SaleStartTime         uint64
	SaleEndTime           uint64
	MinPurchase           *big.Int
	MaxPurchase           *big.Int
	MinIncrement          *big.Int
}

// IPropertyPropertyInfo is an auto generated low-level Go binding around an user-defined struct.
type IPropertyPropertyInfo struct {
	PropertyNumber string
	Creation       uint64
	PropertyStatus uint8
}

// IPropertyPropertyLocationData is an auto generated low-level Go binding around an user-defined struct.
type IPropertyPropertyLocationData struct {
	Country    string
	County     string
	City       string
	State      string
	PostalCode string
	Address1   string
	Address2   string
}

// PropertyManagerMetaData contains all meta data concerning the PropertyManager contract.
var PropertyManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_accessControlAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"propertyId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"propertyNumber\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"propertyAddress\",\"type\":\"address\"}],\"name\":\"PropertyCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"PropertyId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"propertyNumber\",\"type\":\"string\"}],\"name\":\"PropertyCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"soltId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"soldQuantity\",\"type\":\"uint256\"}],\"name\":\"PropertySoldQuantityUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"propertyNumber\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumIProperty.PropertyStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"PropertyStatusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"propertyNumber\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumIProperty.TenancyStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"PropertyTenancyStatusUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PropertyCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accessControl\",\"outputs\":[{\"internalType\":\"contractIAccessControl\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"propertyNumber\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"propertyNumber\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"creation\",\"type\":\"uint64\"},{\"internalType\":\"enumIProperty.PropertyStatus\",\"name\":\"propertyStatus\",\"type\":\"uint8\"}],\"internalType\":\"structIProperty.PropertyInfo\",\"name\":\"info\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"country\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"county\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"city\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"state\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"postalCode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"address1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"address2\",\"type\":\"string\"}],\"internalType\":\"structIProperty.PropertyLocationData\",\"name\":\"location\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"propertyPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"soldQuantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"projectedAnnualReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"projectedRentalYield\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"rentStartDate\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"saleStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"saleEndTime\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"minPurchase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPurchase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minIncrement\",\"type\":\"uint256\"}],\"internalType\":\"structIProperty.InvestmentValue\",\"name\":\"investmentValue\",\"type\":\"tuple\"}],\"name\":\"createProperty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"propertyNumber\",\"type\":\"string\"}],\"name\":\"getPropertyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"propertyNumber\",\"type\":\"string\"}],\"name\":\"getPropertyInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"propertyNumber\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"creation\",\"type\":\"uint64\"},{\"internalType\":\"enumIProperty.PropertyStatus\",\"name\":\"propertyStatus\",\"type\":\"uint8\"}],\"internalType\":\"structIProperty.PropertyInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"soltId\",\"type\":\"uint256\"}],\"name\":\"getPropertyInfoBySolt\",\"outputs\":[{\"internalType\":\"contractIProperty\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"propertyNumber\",\"type\":\"string\"}],\"name\":\"getPropertyInvestmentValue\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"propertyPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"soldQuantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"projectedAnnualReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"projectedRentalYield\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"rentStartDate\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"saleStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"saleEndTime\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"minPurchase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPurchase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minIncrement\",\"type\":\"uint256\"}],\"internalType\":\"structIProperty.InvestmentValue\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"propertys\",\"outputs\":[{\"internalType\":\"contractIProperty\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"soltOfPropertys\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"soltId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"soldQuantity\",\"type\":\"uint256\"}],\"name\":\"updatePropertySoldQuantity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"propertyNumber\",\"type\":\"string\"},{\"internalType\":\"enumIProperty.PropertyStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"updatePropertyStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"propertyNumber\",\"type\":\"string\"},{\"internalType\":\"enumIProperty.TenancyStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"updatePropertyTenancyStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PropertyManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use PropertyManagerMetaData.ABI instead.
var PropertyManagerABI = PropertyManagerMetaData.ABI

// PropertyManager is an auto generated Go binding around an Ethereum contract.
type PropertyManager struct {
	PropertyManagerCaller     // Read-only binding to the contract
	PropertyManagerTransactor // Write-only binding to the contract
	PropertyManagerFilterer   // Log filterer for contract events
}

// PropertyManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PropertyManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PropertyManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PropertyManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PropertyManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PropertyManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PropertyManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PropertyManagerSession struct {
	Contract     *PropertyManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PropertyManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PropertyManagerCallerSession struct {
	Contract *PropertyManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PropertyManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PropertyManagerTransactorSession struct {
	Contract     *PropertyManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PropertyManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PropertyManagerRaw struct {
	Contract *PropertyManager // Generic contract binding to access the raw methods on
}

// PropertyManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PropertyManagerCallerRaw struct {
	Contract *PropertyManagerCaller // Generic read-only contract binding to access the raw methods on
}

// PropertyManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PropertyManagerTransactorRaw struct {
	Contract *PropertyManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPropertyManager creates a new instance of PropertyManager, bound to a specific deployed contract.
func NewPropertyManager(address common.Address, backend bind.ContractBackend) (*PropertyManager, error) {
	contract, err := bindPropertyManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PropertyManager{PropertyManagerCaller: PropertyManagerCaller{contract: contract}, PropertyManagerTransactor: PropertyManagerTransactor{contract: contract}, PropertyManagerFilterer: PropertyManagerFilterer{contract: contract}}, nil
}

// NewPropertyManagerCaller creates a new read-only instance of PropertyManager, bound to a specific deployed contract.
func NewPropertyManagerCaller(address common.Address, caller bind.ContractCaller) (*PropertyManagerCaller, error) {
	contract, err := bindPropertyManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PropertyManagerCaller{contract: contract}, nil
}

// NewPropertyManagerTransactor creates a new write-only instance of PropertyManager, bound to a specific deployed contract.
func NewPropertyManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*PropertyManagerTransactor, error) {
	contract, err := bindPropertyManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PropertyManagerTransactor{contract: contract}, nil
}

// NewPropertyManagerFilterer creates a new log filterer instance of PropertyManager, bound to a specific deployed contract.
func NewPropertyManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*PropertyManagerFilterer, error) {
	contract, err := bindPropertyManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PropertyManagerFilterer{contract: contract}, nil
}

// bindPropertyManager binds a generic wrapper to an already deployed contract.
func bindPropertyManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PropertyManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PropertyManager *PropertyManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PropertyManager.Contract.PropertyManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PropertyManager *PropertyManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PropertyManager.Contract.PropertyManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PropertyManager *PropertyManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PropertyManager.Contract.PropertyManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PropertyManager *PropertyManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PropertyManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PropertyManager *PropertyManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PropertyManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PropertyManager *PropertyManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PropertyManager.Contract.contract.Transact(opts, method, params...)
}

// PropertyCount is a free data retrieval call binding the contract method 0xbb1c16df.
//
// Solidity: function PropertyCount() view returns(uint256)
func (_PropertyManager *PropertyManagerCaller) PropertyCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PropertyManager.contract.Call(opts, &out, "PropertyCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PropertyCount is a free data retrieval call binding the contract method 0xbb1c16df.
//
// Solidity: function PropertyCount() view returns(uint256)
func (_PropertyManager *PropertyManagerSession) PropertyCount() (*big.Int, error) {
	return _PropertyManager.Contract.PropertyCount(&_PropertyManager.CallOpts)
}

// PropertyCount is a free data retrieval call binding the contract method 0xbb1c16df.
//
// Solidity: function PropertyCount() view returns(uint256)
func (_PropertyManager *PropertyManagerCallerSession) PropertyCount() (*big.Int, error) {
	return _PropertyManager.Contract.PropertyCount(&_PropertyManager.CallOpts)
}

// AccessControl is a free data retrieval call binding the contract method 0x13007d55.
//
// Solidity: function accessControl() view returns(address)
func (_PropertyManager *PropertyManagerCaller) AccessControl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PropertyManager.contract.Call(opts, &out, "accessControl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccessControl is a free data retrieval call binding the contract method 0x13007d55.
//
// Solidity: function accessControl() view returns(address)
func (_PropertyManager *PropertyManagerSession) AccessControl() (common.Address, error) {
	return _PropertyManager.Contract.AccessControl(&_PropertyManager.CallOpts)
}

// AccessControl is a free data retrieval call binding the contract method 0x13007d55.
//
// Solidity: function accessControl() view returns(address)
func (_PropertyManager *PropertyManagerCallerSession) AccessControl() (common.Address, error) {
	return _PropertyManager.Contract.AccessControl(&_PropertyManager.CallOpts)
}

// GetPropertyAddress is a free data retrieval call binding the contract method 0xc6ef7a4b.
//
// Solidity: function getPropertyAddress(string propertyNumber) view returns(address)
func (_PropertyManager *PropertyManagerCaller) GetPropertyAddress(opts *bind.CallOpts, propertyNumber string) (common.Address, error) {
	var out []interface{}
	err := _PropertyManager.contract.Call(opts, &out, "getPropertyAddress", propertyNumber)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPropertyAddress is a free data retrieval call binding the contract method 0xc6ef7a4b.
//
// Solidity: function getPropertyAddress(string propertyNumber) view returns(address)
func (_PropertyManager *PropertyManagerSession) GetPropertyAddress(propertyNumber string) (common.Address, error) {
	return _PropertyManager.Contract.GetPropertyAddress(&_PropertyManager.CallOpts, propertyNumber)
}

// GetPropertyAddress is a free data retrieval call binding the contract method 0xc6ef7a4b.
//
// Solidity: function getPropertyAddress(string propertyNumber) view returns(address)
func (_PropertyManager *PropertyManagerCallerSession) GetPropertyAddress(propertyNumber string) (common.Address, error) {
	return _PropertyManager.Contract.GetPropertyAddress(&_PropertyManager.CallOpts, propertyNumber)
}

// GetPropertyInfo is a free data retrieval call binding the contract method 0x2ac30626.
//
// Solidity: function getPropertyInfo(string propertyNumber) view returns((string,uint64,uint8))
func (_PropertyManager *PropertyManagerCaller) GetPropertyInfo(opts *bind.CallOpts, propertyNumber string) (IPropertyPropertyInfo, error) {
	var out []interface{}
	err := _PropertyManager.contract.Call(opts, &out, "getPropertyInfo", propertyNumber)

	if err != nil {
		return *new(IPropertyPropertyInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IPropertyPropertyInfo)).(*IPropertyPropertyInfo)

	return out0, err

}

// GetPropertyInfo is a free data retrieval call binding the contract method 0x2ac30626.
//
// Solidity: function getPropertyInfo(string propertyNumber) view returns((string,uint64,uint8))
func (_PropertyManager *PropertyManagerSession) GetPropertyInfo(propertyNumber string) (IPropertyPropertyInfo, error) {
	return _PropertyManager.Contract.GetPropertyInfo(&_PropertyManager.CallOpts, propertyNumber)
}

// GetPropertyInfo is a free data retrieval call binding the contract method 0x2ac30626.
//
// Solidity: function getPropertyInfo(string propertyNumber) view returns((string,uint64,uint8))
func (_PropertyManager *PropertyManagerCallerSession) GetPropertyInfo(propertyNumber string) (IPropertyPropertyInfo, error) {
	return _PropertyManager.Contract.GetPropertyInfo(&_PropertyManager.CallOpts, propertyNumber)
}

// GetPropertyInfoBySolt is a free data retrieval call binding the contract method 0x298a3ed0.
//
// Solidity: function getPropertyInfoBySolt(uint256 soltId) view returns(address)
func (_PropertyManager *PropertyManagerCaller) GetPropertyInfoBySolt(opts *bind.CallOpts, soltId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PropertyManager.contract.Call(opts, &out, "getPropertyInfoBySolt", soltId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPropertyInfoBySolt is a free data retrieval call binding the contract method 0x298a3ed0.
//
// Solidity: function getPropertyInfoBySolt(uint256 soltId) view returns(address)
func (_PropertyManager *PropertyManagerSession) GetPropertyInfoBySolt(soltId *big.Int) (common.Address, error) {
	return _PropertyManager.Contract.GetPropertyInfoBySolt(&_PropertyManager.CallOpts, soltId)
}

// GetPropertyInfoBySolt is a free data retrieval call binding the contract method 0x298a3ed0.
//
// Solidity: function getPropertyInfoBySolt(uint256 soltId) view returns(address)
func (_PropertyManager *PropertyManagerCallerSession) GetPropertyInfoBySolt(soltId *big.Int) (common.Address, error) {
	return _PropertyManager.Contract.GetPropertyInfoBySolt(&_PropertyManager.CallOpts, soltId)
}

// GetPropertyInvestmentValue is a free data retrieval call binding the contract method 0x192c0834.
//
// Solidity: function getPropertyInvestmentValue(string propertyNumber) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint256,uint256,uint256))
func (_PropertyManager *PropertyManagerCaller) GetPropertyInvestmentValue(opts *bind.CallOpts, propertyNumber string) (IPropertyInvestmentValue, error) {
	var out []interface{}
	err := _PropertyManager.contract.Call(opts, &out, "getPropertyInvestmentValue", propertyNumber)

	if err != nil {
		return *new(IPropertyInvestmentValue), err
	}

	out0 := *abi.ConvertType(out[0], new(IPropertyInvestmentValue)).(*IPropertyInvestmentValue)

	return out0, err

}

// GetPropertyInvestmentValue is a free data retrieval call binding the contract method 0x192c0834.
//
// Solidity: function getPropertyInvestmentValue(string propertyNumber) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint256,uint256,uint256))
func (_PropertyManager *PropertyManagerSession) GetPropertyInvestmentValue(propertyNumber string) (IPropertyInvestmentValue, error) {
	return _PropertyManager.Contract.GetPropertyInvestmentValue(&_PropertyManager.CallOpts, propertyNumber)
}

// GetPropertyInvestmentValue is a free data retrieval call binding the contract method 0x192c0834.
//
// Solidity: function getPropertyInvestmentValue(string propertyNumber) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint256,uint256,uint256))
func (_PropertyManager *PropertyManagerCallerSession) GetPropertyInvestmentValue(propertyNumber string) (IPropertyInvestmentValue, error) {
	return _PropertyManager.Contract.GetPropertyInvestmentValue(&_PropertyManager.CallOpts, propertyNumber)
}

// Propertys is a free data retrieval call binding the contract method 0xf04ca6e8.
//
// Solidity: function propertys(string ) view returns(address)
func (_PropertyManager *PropertyManagerCaller) Propertys(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _PropertyManager.contract.Call(opts, &out, "propertys", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Propertys is a free data retrieval call binding the contract method 0xf04ca6e8.
//
// Solidity: function propertys(string ) view returns(address)
func (_PropertyManager *PropertyManagerSession) Propertys(arg0 string) (common.Address, error) {
	return _PropertyManager.Contract.Propertys(&_PropertyManager.CallOpts, arg0)
}

// Propertys is a free data retrieval call binding the contract method 0xf04ca6e8.
//
// Solidity: function propertys(string ) view returns(address)
func (_PropertyManager *PropertyManagerCallerSession) Propertys(arg0 string) (common.Address, error) {
	return _PropertyManager.Contract.Propertys(&_PropertyManager.CallOpts, arg0)
}

// SoltOfPropertys is a free data retrieval call binding the contract method 0x1b198851.
//
// Solidity: function soltOfPropertys(uint256 ) view returns(string)
func (_PropertyManager *PropertyManagerCaller) SoltOfPropertys(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _PropertyManager.contract.Call(opts, &out, "soltOfPropertys", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SoltOfPropertys is a free data retrieval call binding the contract method 0x1b198851.
//
// Solidity: function soltOfPropertys(uint256 ) view returns(string)
func (_PropertyManager *PropertyManagerSession) SoltOfPropertys(arg0 *big.Int) (string, error) {
	return _PropertyManager.Contract.SoltOfPropertys(&_PropertyManager.CallOpts, arg0)
}

// SoltOfPropertys is a free data retrieval call binding the contract method 0x1b198851.
//
// Solidity: function soltOfPropertys(uint256 ) view returns(string)
func (_PropertyManager *PropertyManagerCallerSession) SoltOfPropertys(arg0 *big.Int) (string, error) {
	return _PropertyManager.Contract.SoltOfPropertys(&_PropertyManager.CallOpts, arg0)
}

// CreateProperty is a paid mutator transaction binding the contract method 0x37738921.
//
// Solidity: function createProperty(string propertyNumber, (string,uint64,uint8) info, (string,string,string,string,string,string,string) location, (uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint256,uint256,uint256) investmentValue) returns(uint256)
func (_PropertyManager *PropertyManagerTransactor) CreateProperty(opts *bind.TransactOpts, propertyNumber string, info IPropertyPropertyInfo, location IPropertyPropertyLocationData, investmentValue IPropertyInvestmentValue) (*types.Transaction, error) {
	return _PropertyManager.contract.Transact(opts, "createProperty", propertyNumber, info, location, investmentValue)
}

// CreateProperty is a paid mutator transaction binding the contract method 0x37738921.
//
// Solidity: function createProperty(string propertyNumber, (string,uint64,uint8) info, (string,string,string,string,string,string,string) location, (uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint256,uint256,uint256) investmentValue) returns(uint256)
func (_PropertyManager *PropertyManagerSession) CreateProperty(propertyNumber string, info IPropertyPropertyInfo, location IPropertyPropertyLocationData, investmentValue IPropertyInvestmentValue) (*types.Transaction, error) {
	return _PropertyManager.Contract.CreateProperty(&_PropertyManager.TransactOpts, propertyNumber, info, location, investmentValue)
}

// CreateProperty is a paid mutator transaction binding the contract method 0x37738921.
//
// Solidity: function createProperty(string propertyNumber, (string,uint64,uint8) info, (string,string,string,string,string,string,string) location, (uint256,uint256,uint256,uint256,uint256,uint256,uint64,uint64,uint64,uint256,uint256,uint256) investmentValue) returns(uint256)
func (_PropertyManager *PropertyManagerTransactorSession) CreateProperty(propertyNumber string, info IPropertyPropertyInfo, location IPropertyPropertyLocationData, investmentValue IPropertyInvestmentValue) (*types.Transaction, error) {
	return _PropertyManager.Contract.CreateProperty(&_PropertyManager.TransactOpts, propertyNumber, info, location, investmentValue)
}

// UpdatePropertySoldQuantity is a paid mutator transaction binding the contract method 0x3b304b41.
//
// Solidity: function updatePropertySoldQuantity(uint256 soltId, uint256 soldQuantity) returns()
func (_PropertyManager *PropertyManagerTransactor) UpdatePropertySoldQuantity(opts *bind.TransactOpts, soltId *big.Int, soldQuantity *big.Int) (*types.Transaction, error) {
	return _PropertyManager.contract.Transact(opts, "updatePropertySoldQuantity", soltId, soldQuantity)
}

// UpdatePropertySoldQuantity is a paid mutator transaction binding the contract method 0x3b304b41.
//
// Solidity: function updatePropertySoldQuantity(uint256 soltId, uint256 soldQuantity) returns()
func (_PropertyManager *PropertyManagerSession) UpdatePropertySoldQuantity(soltId *big.Int, soldQuantity *big.Int) (*types.Transaction, error) {
	return _PropertyManager.Contract.UpdatePropertySoldQuantity(&_PropertyManager.TransactOpts, soltId, soldQuantity)
}

// UpdatePropertySoldQuantity is a paid mutator transaction binding the contract method 0x3b304b41.
//
// Solidity: function updatePropertySoldQuantity(uint256 soltId, uint256 soldQuantity) returns()
func (_PropertyManager *PropertyManagerTransactorSession) UpdatePropertySoldQuantity(soltId *big.Int, soldQuantity *big.Int) (*types.Transaction, error) {
	return _PropertyManager.Contract.UpdatePropertySoldQuantity(&_PropertyManager.TransactOpts, soltId, soldQuantity)
}

// UpdatePropertyStatus is a paid mutator transaction binding the contract method 0x21869932.
//
// Solidity: function updatePropertyStatus(string propertyNumber, uint8 newStatus) returns()
func (_PropertyManager *PropertyManagerTransactor) UpdatePropertyStatus(opts *bind.TransactOpts, propertyNumber string, newStatus uint8) (*types.Transaction, error) {
	return _PropertyManager.contract.Transact(opts, "updatePropertyStatus", propertyNumber, newStatus)
}

// UpdatePropertyStatus is a paid mutator transaction binding the contract method 0x21869932.
//
// Solidity: function updatePropertyStatus(string propertyNumber, uint8 newStatus) returns()
func (_PropertyManager *PropertyManagerSession) UpdatePropertyStatus(propertyNumber string, newStatus uint8) (*types.Transaction, error) {
	return _PropertyManager.Contract.UpdatePropertyStatus(&_PropertyManager.TransactOpts, propertyNumber, newStatus)
}

// UpdatePropertyStatus is a paid mutator transaction binding the contract method 0x21869932.
//
// Solidity: function updatePropertyStatus(string propertyNumber, uint8 newStatus) returns()
func (_PropertyManager *PropertyManagerTransactorSession) UpdatePropertyStatus(propertyNumber string, newStatus uint8) (*types.Transaction, error) {
	return _PropertyManager.Contract.UpdatePropertyStatus(&_PropertyManager.TransactOpts, propertyNumber, newStatus)
}

// UpdatePropertyTenancyStatus is a paid mutator transaction binding the contract method 0xd53c9fce.
//
// Solidity: function updatePropertyTenancyStatus(string propertyNumber, uint8 newStatus) returns()
func (_PropertyManager *PropertyManagerTransactor) UpdatePropertyTenancyStatus(opts *bind.TransactOpts, propertyNumber string, newStatus uint8) (*types.Transaction, error) {
	return _PropertyManager.contract.Transact(opts, "updatePropertyTenancyStatus", propertyNumber, newStatus)
}

// UpdatePropertyTenancyStatus is a paid mutator transaction binding the contract method 0xd53c9fce.
//
// Solidity: function updatePropertyTenancyStatus(string propertyNumber, uint8 newStatus) returns()
func (_PropertyManager *PropertyManagerSession) UpdatePropertyTenancyStatus(propertyNumber string, newStatus uint8) (*types.Transaction, error) {
	return _PropertyManager.Contract.UpdatePropertyTenancyStatus(&_PropertyManager.TransactOpts, propertyNumber, newStatus)
}

// UpdatePropertyTenancyStatus is a paid mutator transaction binding the contract method 0xd53c9fce.
//
// Solidity: function updatePropertyTenancyStatus(string propertyNumber, uint8 newStatus) returns()
func (_PropertyManager *PropertyManagerTransactorSession) UpdatePropertyTenancyStatus(propertyNumber string, newStatus uint8) (*types.Transaction, error) {
	return _PropertyManager.Contract.UpdatePropertyTenancyStatus(&_PropertyManager.TransactOpts, propertyNumber, newStatus)
}

// PropertyManagerPropertyCreatedIterator is returned from FilterPropertyCreated and is used to iterate over the raw logs and unpacked data for PropertyCreated events raised by the PropertyManager contract.
type PropertyManagerPropertyCreatedIterator struct {
	Event *PropertyManagerPropertyCreated // Event containing the contract specifics and raw log

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
func (it *PropertyManagerPropertyCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyManagerPropertyCreated)
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
		it.Event = new(PropertyManagerPropertyCreated)
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
func (it *PropertyManagerPropertyCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyManagerPropertyCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyManagerPropertyCreated represents a PropertyCreated event raised by the PropertyManager contract.
type PropertyManagerPropertyCreated struct {
	PropertyId      *big.Int
	PropertyNumber  string
	PropertyAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPropertyCreated is a free log retrieval operation binding the contract event 0xf6d5ae013bcc4d04e138a31df7f9a463c0a79a68f43f158b836b733de95725ac.
//
// Solidity: event PropertyCreated(uint256 indexed propertyId, string propertyNumber, address propertyAddress)
func (_PropertyManager *PropertyManagerFilterer) FilterPropertyCreated(opts *bind.FilterOpts, propertyId []*big.Int) (*PropertyManagerPropertyCreatedIterator, error) {

	var propertyIdRule []interface{}
	for _, propertyIdItem := range propertyId {
		propertyIdRule = append(propertyIdRule, propertyIdItem)
	}

	logs, sub, err := _PropertyManager.contract.FilterLogs(opts, "PropertyCreated", propertyIdRule)
	if err != nil {
		return nil, err
	}
	return &PropertyManagerPropertyCreatedIterator{contract: _PropertyManager.contract, event: "PropertyCreated", logs: logs, sub: sub}, nil
}

// WatchPropertyCreated is a free log subscription operation binding the contract event 0xf6d5ae013bcc4d04e138a31df7f9a463c0a79a68f43f158b836b733de95725ac.
//
// Solidity: event PropertyCreated(uint256 indexed propertyId, string propertyNumber, address propertyAddress)
func (_PropertyManager *PropertyManagerFilterer) WatchPropertyCreated(opts *bind.WatchOpts, sink chan<- *PropertyManagerPropertyCreated, propertyId []*big.Int) (event.Subscription, error) {

	var propertyIdRule []interface{}
	for _, propertyIdItem := range propertyId {
		propertyIdRule = append(propertyIdRule, propertyIdItem)
	}

	logs, sub, err := _PropertyManager.contract.WatchLogs(opts, "PropertyCreated", propertyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyManagerPropertyCreated)
				if err := _PropertyManager.contract.UnpackLog(event, "PropertyCreated", log); err != nil {
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

// ParsePropertyCreated is a log parse operation binding the contract event 0xf6d5ae013bcc4d04e138a31df7f9a463c0a79a68f43f158b836b733de95725ac.
//
// Solidity: event PropertyCreated(uint256 indexed propertyId, string propertyNumber, address propertyAddress)
func (_PropertyManager *PropertyManagerFilterer) ParsePropertyCreated(log types.Log) (*PropertyManagerPropertyCreated, error) {
	event := new(PropertyManagerPropertyCreated)
	if err := _PropertyManager.contract.UnpackLog(event, "PropertyCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyManagerPropertyCreated0Iterator is returned from FilterPropertyCreated0 and is used to iterate over the raw logs and unpacked data for PropertyCreated0 events raised by the PropertyManager contract.
type PropertyManagerPropertyCreated0Iterator struct {
	Event *PropertyManagerPropertyCreated0 // Event containing the contract specifics and raw log

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
func (it *PropertyManagerPropertyCreated0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyManagerPropertyCreated0)
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
		it.Event = new(PropertyManagerPropertyCreated0)
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
func (it *PropertyManagerPropertyCreated0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyManagerPropertyCreated0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyManagerPropertyCreated0 represents a PropertyCreated0 event raised by the PropertyManager contract.
type PropertyManagerPropertyCreated0 struct {
	PropertyId     *big.Int
	PropertyNumber string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPropertyCreated0 is a free log retrieval operation binding the contract event 0x3f7f078bc3fcf62a2b8b118f834f7db41c3c4a5aa20b849f5443290cebaa24b6.
//
// Solidity: event PropertyCreated(uint256 indexed PropertyId, string propertyNumber)
func (_PropertyManager *PropertyManagerFilterer) FilterPropertyCreated0(opts *bind.FilterOpts, PropertyId []*big.Int) (*PropertyManagerPropertyCreated0Iterator, error) {

	var PropertyIdRule []interface{}
	for _, PropertyIdItem := range PropertyId {
		PropertyIdRule = append(PropertyIdRule, PropertyIdItem)
	}

	logs, sub, err := _PropertyManager.contract.FilterLogs(opts, "PropertyCreated0", PropertyIdRule)
	if err != nil {
		return nil, err
	}
	return &PropertyManagerPropertyCreated0Iterator{contract: _PropertyManager.contract, event: "PropertyCreated0", logs: logs, sub: sub}, nil
}

// WatchPropertyCreated0 is a free log subscription operation binding the contract event 0x3f7f078bc3fcf62a2b8b118f834f7db41c3c4a5aa20b849f5443290cebaa24b6.
//
// Solidity: event PropertyCreated(uint256 indexed PropertyId, string propertyNumber)
func (_PropertyManager *PropertyManagerFilterer) WatchPropertyCreated0(opts *bind.WatchOpts, sink chan<- *PropertyManagerPropertyCreated0, PropertyId []*big.Int) (event.Subscription, error) {

	var PropertyIdRule []interface{}
	for _, PropertyIdItem := range PropertyId {
		PropertyIdRule = append(PropertyIdRule, PropertyIdItem)
	}

	logs, sub, err := _PropertyManager.contract.WatchLogs(opts, "PropertyCreated0", PropertyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyManagerPropertyCreated0)
				if err := _PropertyManager.contract.UnpackLog(event, "PropertyCreated0", log); err != nil {
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

// ParsePropertyCreated0 is a log parse operation binding the contract event 0x3f7f078bc3fcf62a2b8b118f834f7db41c3c4a5aa20b849f5443290cebaa24b6.
//
// Solidity: event PropertyCreated(uint256 indexed PropertyId, string propertyNumber)
func (_PropertyManager *PropertyManagerFilterer) ParsePropertyCreated0(log types.Log) (*PropertyManagerPropertyCreated0, error) {
	event := new(PropertyManagerPropertyCreated0)
	if err := _PropertyManager.contract.UnpackLog(event, "PropertyCreated0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyManagerPropertySoldQuantityUpdatedIterator is returned from FilterPropertySoldQuantityUpdated and is used to iterate over the raw logs and unpacked data for PropertySoldQuantityUpdated events raised by the PropertyManager contract.
type PropertyManagerPropertySoldQuantityUpdatedIterator struct {
	Event *PropertyManagerPropertySoldQuantityUpdated // Event containing the contract specifics and raw log

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
func (it *PropertyManagerPropertySoldQuantityUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyManagerPropertySoldQuantityUpdated)
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
		it.Event = new(PropertyManagerPropertySoldQuantityUpdated)
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
func (it *PropertyManagerPropertySoldQuantityUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyManagerPropertySoldQuantityUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyManagerPropertySoldQuantityUpdated represents a PropertySoldQuantityUpdated event raised by the PropertyManager contract.
type PropertyManagerPropertySoldQuantityUpdated struct {
	SoltId       *big.Int
	SoldQuantity *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPropertySoldQuantityUpdated is a free log retrieval operation binding the contract event 0x905f11e45f6563dd53035b3622d88a0f55bdd2d79b502d66a8a62cca0588575d.
//
// Solidity: event PropertySoldQuantityUpdated(uint256 indexed soltId, uint256 soldQuantity)
func (_PropertyManager *PropertyManagerFilterer) FilterPropertySoldQuantityUpdated(opts *bind.FilterOpts, soltId []*big.Int) (*PropertyManagerPropertySoldQuantityUpdatedIterator, error) {

	var soltIdRule []interface{}
	for _, soltIdItem := range soltId {
		soltIdRule = append(soltIdRule, soltIdItem)
	}

	logs, sub, err := _PropertyManager.contract.FilterLogs(opts, "PropertySoldQuantityUpdated", soltIdRule)
	if err != nil {
		return nil, err
	}
	return &PropertyManagerPropertySoldQuantityUpdatedIterator{contract: _PropertyManager.contract, event: "PropertySoldQuantityUpdated", logs: logs, sub: sub}, nil
}

// WatchPropertySoldQuantityUpdated is a free log subscription operation binding the contract event 0x905f11e45f6563dd53035b3622d88a0f55bdd2d79b502d66a8a62cca0588575d.
//
// Solidity: event PropertySoldQuantityUpdated(uint256 indexed soltId, uint256 soldQuantity)
func (_PropertyManager *PropertyManagerFilterer) WatchPropertySoldQuantityUpdated(opts *bind.WatchOpts, sink chan<- *PropertyManagerPropertySoldQuantityUpdated, soltId []*big.Int) (event.Subscription, error) {

	var soltIdRule []interface{}
	for _, soltIdItem := range soltId {
		soltIdRule = append(soltIdRule, soltIdItem)
	}

	logs, sub, err := _PropertyManager.contract.WatchLogs(opts, "PropertySoldQuantityUpdated", soltIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyManagerPropertySoldQuantityUpdated)
				if err := _PropertyManager.contract.UnpackLog(event, "PropertySoldQuantityUpdated", log); err != nil {
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

// ParsePropertySoldQuantityUpdated is a log parse operation binding the contract event 0x905f11e45f6563dd53035b3622d88a0f55bdd2d79b502d66a8a62cca0588575d.
//
// Solidity: event PropertySoldQuantityUpdated(uint256 indexed soltId, uint256 soldQuantity)
func (_PropertyManager *PropertyManagerFilterer) ParsePropertySoldQuantityUpdated(log types.Log) (*PropertyManagerPropertySoldQuantityUpdated, error) {
	event := new(PropertyManagerPropertySoldQuantityUpdated)
	if err := _PropertyManager.contract.UnpackLog(event, "PropertySoldQuantityUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyManagerPropertyStatusUpdatedIterator is returned from FilterPropertyStatusUpdated and is used to iterate over the raw logs and unpacked data for PropertyStatusUpdated events raised by the PropertyManager contract.
type PropertyManagerPropertyStatusUpdatedIterator struct {
	Event *PropertyManagerPropertyStatusUpdated // Event containing the contract specifics and raw log

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
func (it *PropertyManagerPropertyStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyManagerPropertyStatusUpdated)
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
		it.Event = new(PropertyManagerPropertyStatusUpdated)
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
func (it *PropertyManagerPropertyStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyManagerPropertyStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyManagerPropertyStatusUpdated represents a PropertyStatusUpdated event raised by the PropertyManager contract.
type PropertyManagerPropertyStatusUpdated struct {
	PropertyNumber common.Hash
	NewStatus      uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPropertyStatusUpdated is a free log retrieval operation binding the contract event 0x74cc0169ad2de84bd3372ff650775fb40a55ebc66fb041d5870cb82b9dff68b4.
//
// Solidity: event PropertyStatusUpdated(string indexed propertyNumber, uint8 newStatus)
func (_PropertyManager *PropertyManagerFilterer) FilterPropertyStatusUpdated(opts *bind.FilterOpts, propertyNumber []string) (*PropertyManagerPropertyStatusUpdatedIterator, error) {

	var propertyNumberRule []interface{}
	for _, propertyNumberItem := range propertyNumber {
		propertyNumberRule = append(propertyNumberRule, propertyNumberItem)
	}

	logs, sub, err := _PropertyManager.contract.FilterLogs(opts, "PropertyStatusUpdated", propertyNumberRule)
	if err != nil {
		return nil, err
	}
	return &PropertyManagerPropertyStatusUpdatedIterator{contract: _PropertyManager.contract, event: "PropertyStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchPropertyStatusUpdated is a free log subscription operation binding the contract event 0x74cc0169ad2de84bd3372ff650775fb40a55ebc66fb041d5870cb82b9dff68b4.
//
// Solidity: event PropertyStatusUpdated(string indexed propertyNumber, uint8 newStatus)
func (_PropertyManager *PropertyManagerFilterer) WatchPropertyStatusUpdated(opts *bind.WatchOpts, sink chan<- *PropertyManagerPropertyStatusUpdated, propertyNumber []string) (event.Subscription, error) {

	var propertyNumberRule []interface{}
	for _, propertyNumberItem := range propertyNumber {
		propertyNumberRule = append(propertyNumberRule, propertyNumberItem)
	}

	logs, sub, err := _PropertyManager.contract.WatchLogs(opts, "PropertyStatusUpdated", propertyNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyManagerPropertyStatusUpdated)
				if err := _PropertyManager.contract.UnpackLog(event, "PropertyStatusUpdated", log); err != nil {
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

// ParsePropertyStatusUpdated is a log parse operation binding the contract event 0x74cc0169ad2de84bd3372ff650775fb40a55ebc66fb041d5870cb82b9dff68b4.
//
// Solidity: event PropertyStatusUpdated(string indexed propertyNumber, uint8 newStatus)
func (_PropertyManager *PropertyManagerFilterer) ParsePropertyStatusUpdated(log types.Log) (*PropertyManagerPropertyStatusUpdated, error) {
	event := new(PropertyManagerPropertyStatusUpdated)
	if err := _PropertyManager.contract.UnpackLog(event, "PropertyStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PropertyManagerPropertyTenancyStatusUpdatedIterator is returned from FilterPropertyTenancyStatusUpdated and is used to iterate over the raw logs and unpacked data for PropertyTenancyStatusUpdated events raised by the PropertyManager contract.
type PropertyManagerPropertyTenancyStatusUpdatedIterator struct {
	Event *PropertyManagerPropertyTenancyStatusUpdated // Event containing the contract specifics and raw log

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
func (it *PropertyManagerPropertyTenancyStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropertyManagerPropertyTenancyStatusUpdated)
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
		it.Event = new(PropertyManagerPropertyTenancyStatusUpdated)
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
func (it *PropertyManagerPropertyTenancyStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropertyManagerPropertyTenancyStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropertyManagerPropertyTenancyStatusUpdated represents a PropertyTenancyStatusUpdated event raised by the PropertyManager contract.
type PropertyManagerPropertyTenancyStatusUpdated struct {
	PropertyNumber common.Hash
	NewStatus      uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPropertyTenancyStatusUpdated is a free log retrieval operation binding the contract event 0x246422eda0d4c071ef612582fcfd2ded7a62ec6cfd834ac9542e73bc268789fc.
//
// Solidity: event PropertyTenancyStatusUpdated(string indexed propertyNumber, uint8 newStatus)
func (_PropertyManager *PropertyManagerFilterer) FilterPropertyTenancyStatusUpdated(opts *bind.FilterOpts, propertyNumber []string) (*PropertyManagerPropertyTenancyStatusUpdatedIterator, error) {

	var propertyNumberRule []interface{}
	for _, propertyNumberItem := range propertyNumber {
		propertyNumberRule = append(propertyNumberRule, propertyNumberItem)
	}

	logs, sub, err := _PropertyManager.contract.FilterLogs(opts, "PropertyTenancyStatusUpdated", propertyNumberRule)
	if err != nil {
		return nil, err
	}
	return &PropertyManagerPropertyTenancyStatusUpdatedIterator{contract: _PropertyManager.contract, event: "PropertyTenancyStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchPropertyTenancyStatusUpdated is a free log subscription operation binding the contract event 0x246422eda0d4c071ef612582fcfd2ded7a62ec6cfd834ac9542e73bc268789fc.
//
// Solidity: event PropertyTenancyStatusUpdated(string indexed propertyNumber, uint8 newStatus)
func (_PropertyManager *PropertyManagerFilterer) WatchPropertyTenancyStatusUpdated(opts *bind.WatchOpts, sink chan<- *PropertyManagerPropertyTenancyStatusUpdated, propertyNumber []string) (event.Subscription, error) {

	var propertyNumberRule []interface{}
	for _, propertyNumberItem := range propertyNumber {
		propertyNumberRule = append(propertyNumberRule, propertyNumberItem)
	}

	logs, sub, err := _PropertyManager.contract.WatchLogs(opts, "PropertyTenancyStatusUpdated", propertyNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropertyManagerPropertyTenancyStatusUpdated)
				if err := _PropertyManager.contract.UnpackLog(event, "PropertyTenancyStatusUpdated", log); err != nil {
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

// ParsePropertyTenancyStatusUpdated is a log parse operation binding the contract event 0x246422eda0d4c071ef612582fcfd2ded7a62ec6cfd834ac9542e73bc268789fc.
//
// Solidity: event PropertyTenancyStatusUpdated(string indexed propertyNumber, uint8 newStatus)
func (_PropertyManager *PropertyManagerFilterer) ParsePropertyTenancyStatusUpdated(log types.Log) (*PropertyManagerPropertyTenancyStatusUpdated, error) {
	event := new(PropertyManagerPropertyTenancyStatusUpdated)
	if err := _PropertyManager.contract.UnpackLog(event, "PropertyTenancyStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
