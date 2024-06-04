package chain

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"go-server/config"
	"go-server/pkg/blockchain"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ErrParsePublicKey = errors.New("got error while parsing private key to public key")
	ErrGasPrice       = errors.New("get suggested gas price error")
	ErrGasLimit       = errors.New("estimate gas limit error")
)

type Web3Instance struct {
	client *ethclient.Client
	pk     *ecdsa.PrivateKey
}

func NewWeb3Instance() (*Web3Instance, error) {
	client, err := ethclient.Dial(config.C.Chain.RPC)
	if err != nil {
		return nil, err
	}

	// Load your private key (assuming you have it stored securely)
	privateKey, err := crypto.HexToECDSA(config.C.Secret.MasterWalletPrivateKey)
	if err != nil {
		return nil, err
	}

	return &Web3Instance{
		client: client,
		pk:     privateKey,
	}, nil
}

func (w *Web3Instance) SendTransaction(ctx context.Context, data ...any) (string, error) {
	publicKey := w.pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", ErrParsePublicKey
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Get the pending nonce for the transaction
	nonce, err := w.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	contractInstance, err := blockchain.NewContract(config.C.Contract.Address, config.C.Chain.RPC)
	if err != nil {
		return "", err
	}

	txData, err := contractInstance.CallAbi("create", data...)
	if err != nil {
		return "", err
	}

	var (
		amout     = big.NewInt(0)
		toAddress = common.HexToAddress(config.C.Contract.Address)
	)
	gasLimit, err := w.client.EstimateGas(ctx, ethereum.CallMsg{
		From:  fromAddress,
		To:    &toAddress,
		Data:  txData,
		Value: amout,
	})
	if err != nil {
		return "", ErrGasLimit
	}

	gasPrice, err := w.client.SuggestGasPrice(ctx)
	if err != nil {
		return "", ErrGasPrice
	}

	// Double gas limit for faster
	tx := types.NewTransaction(nonce, common.HexToAddress(config.C.Contract.Address), amout, gasLimit*2, gasPrice, txData)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(new(big.Int).SetUint64(uint64(config.C.Chain.ID))), w.pk)
	if err != nil {
		return "", err
	}

	err = w.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return "", err
	}

	return signedTx.Hash().String(), nil
}
