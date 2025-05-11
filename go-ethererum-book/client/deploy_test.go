package main

import (
	"book/build/dlt"
	"book/util"
	"context"
	"encoding/hex"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2/log"
	"github.com/stretchr/testify/require"
)

func TestDeploy(t *testing.T) {

	err := util.InitForGanache([]string{"audit", "init", "deploy"})
	require.NoError(t, err)

	addess := util.GetAddress("deploy")
	privateKey := util.GetPrivateKey("deploy")

	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, "http://localhost:9474")
	require.NoError(t, err)

	id, err := client.ChainID(ctx)
	require.NoError(t, err)

	t.Run("deploy contract", func(t *testing.T) {

		// deploy contract

		txnOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, id)
		require.NoError(t, err)

		nonce, err := client.PendingNonceAt(context.Background(), addess)
		require.NoError(t, err)

		gasPrice, err := client.SuggestGasPrice(context.Background())
		require.NoError(t, err)

		txnOpts.Nonce = big.NewInt(int64(nonce))
		txnOpts.Value = big.NewInt(0)
		txnOpts.GasLimit = uint64(300000)
		txnOpts.GasPrice = gasPrice

		contractAddress, txn, _, err := dlt.DeployRequestContract(txnOpts, client)
		require.NoError(t, err)

		log.Infof("deployed contract with address : %v txn.hash %v", contractAddress.Hex(), txn.Hash().Hex())

		// load deployed contract
		instance, err := dlt.NewRequestContract(contractAddress, client)
		require.NoError(t, err)

		// invoke & query smart contracts

		nonce, err = client.PendingNonceAt(context.Background(), addess)
		require.NoError(t, err)
		txnOpts.Nonce = big.NewInt(int64(nonce))

		var requestId = big.NewInt(1)
		txn, err = instance.SaveRequest(txnOpts, dlt.RequestContractRequest{
			Id: requestId,
			// Name: "sample-nft",
		})
		require.NoError(t, err)
		log.Info("saved request with txnId : ", txn.Hash().Hex())

		request, err := instance.QueryRequest(nil, requestId)
		require.NoError(t, err)
		log.Infof("queried request with txnId : %+v", request)

		nonce, err = client.PendingNonceAt(context.Background(), addess)
		require.NoError(t, err)
		txnOpts.Nonce = big.NewInt(int64(nonce))

		txn, err = instance.ApproveRequest(txnOpts, requestId, big.NewInt(time.Now().Unix()), addess)
		require.NoError(t, err)
		log.Info("approved request with txnId : ", txn.Hash().Hex())

		request, err = instance.QueryRequest(nil, requestId)
		require.NoError(t, err)

		log.Infof("queried request with txnId : %+v", request)

		// get bytecode of contract

		byteCode, err := client.CodeAt(ctx, contractAddress, nil)
		require.NoError(t, err)

		log.Infof("byteCode : %+v", hex.EncodeToString(byteCode))

	})

}
