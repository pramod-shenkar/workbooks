package main

import (
	"book/build/dlt"
	"book/util"
	"context"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2/log"
	"github.com/stretchr/testify/require"
)

func TestEvents(t *testing.T) {

	util.InitForGanache([]string{"audit", "init", "deploy"})

	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, "http://localhost:9474")
	require.NoError(t, err)

	wsClient, err := ethclient.DialContext(ctx, "ws://localhost:9474")
	require.NoError(t, err)

	var contractAddress = common.HexToAddress("0x3c31A5d52c0735B222aaF05e137f1FB7dB5DBB6D")

	// listen to event
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)

	subscription, err := wsClient.SubscribeFilterLogs(ctx, query, logs)
	require.NoError(t, err)

	instance, err := dlt.NewRequestContract(contractAddress, client)
	require.NoError(t, err)

	go func() {
		log.Info("listening events")
		for {
			select {
			case err := <-subscription.Err():
				log.Error(err)
			case currentLog := <-logs:
				log.Info(currentLog)
				event, err := instance.ParseSavedEvent(currentLog)
				require.NoError(t, err)

				log.Infof("%+v", event)

				contractAbi, err := abi.JSON(strings.NewReader(string(dlt.RequestContractABI)))
				require.NoError(t, err)

				unpacked, err := contractAbi.Unpack("SavedEvent", currentLog.Data)
				require.NoError(t, err)

				if len(unpacked) != 1 {
					log.Fatalf("unexpected unpacked items count: %v", len(unpacked))
				}

				status, ok := unpacked[0].(bool)
				if !ok {
					log.Fatalf("failed to type assert event value to bool")
				}

				log.Info("Event status:", status)

			}
		}
		log.Info("listener stopped")
	}()

	// invoke & query smart contracts

	addess := util.GetAddress("deploy")
	privateKey := util.GetPrivateKey("deploy")
	id, err := client.ChainID(ctx)
	require.NoError(t, err)

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

	var requestId = big.NewInt(1)
	txn, err := instance.SaveRequest(txnOpts, dlt.RequestContractRequest{
		Id: requestId,
		// Name: "sample-nft",
	})
	require.NoError(t, err)
	log.Info("saved request with txnId : ", txn.Hash().Hex())

	time.Sleep(50 * time.Second)

}
