package blockchain

import (
	"os"

	"github.com/chenzhijie/go-web3"
	"github.com/chenzhijie/go-web3/eth"
)

type SmartContract struct {
	Contract *eth.Contract
}

func NewContract(contractAddress, rpc string) (*SmartContract, error) {
	web3Instance, err := web3.NewWeb3(rpc)
	if err != nil {
		return nil, err
	}

	jsonData, err := os.ReadFile("./resources/abi.json")
	if err != nil {
		return nil, err
	}

	contract, err := web3Instance.Eth.NewContract(string(jsonData), contractAddress)
	if err != nil {
		return nil, err
	}

	return &SmartContract{
		contract,
	}, nil
}

func (sc *SmartContract) CallAbi(funcName string, data ...any) ([]byte, error) {
	return sc.Contract.EncodeABI(funcName, data...)
}
