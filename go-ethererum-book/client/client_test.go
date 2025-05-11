package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"math"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/gofiber/fiber/v2/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {

	// basic client connection

	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, "http://localhost:9474")
	require.NoError(t, err)

	id, err := client.ChainID(ctx)
	require.NoError(t, err)

	log.Info("chainId : ", id)

	t.Run("account balances", func(t *testing.T) {

		var account01 = "0x13B344548F2590B5E75Fe5A7d89b7b676C79f34d"
		var account01Address = common.HexToAddress(account01)

		wei, err := client.BalanceAt(ctx, account01Address, nil)
		assert.NoError(t, err)

		log.Info("balance in wei : ", wei)
		log.Info("balance in ethers : ", WeiToEther(wei))

		wei, err = client.PendingBalanceAt(context.Background(), account01Address)
		assert.NoError(t, err)

		log.Info("pending balance in wei : ", wei)
		log.Info("pending balance in ethers : ", WeiToEther(wei))

	})

	t.Run("generate wallet", func(t *testing.T) {

		privateKey, err := crypto.GenerateKey()
		require.NoError(t, err)

		privateKeyBytes := crypto.FromECDSA(privateKey)
		log.Info("private-key : ", hexutil.Encode(privateKeyBytes))
		// why to use hexutil.Encode instead string() for []byte to string conversion
		// many character from privatekey does not have curresponding ascii characters so such characters string converion looks like garbage

		publicKeyCrypto := privateKey.Public()

		publicKey, ok := publicKeyCrypto.(*ecdsa.PublicKey)
		require.Equal(t, true, ok)

		publicKeyBytes := crypto.FromECDSAPub(publicKey)
		log.Info("public-key : ", hexutil.Encode(publicKeyBytes))

		publicKeyAddress := crypto.PubkeyToAddress(*publicKey).Hex()
		log.Info("public-key address : ", publicKeyAddress)

	})

	t.Run("keystores", func(t *testing.T) {

		// 1. generate empty keystore
		store := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
		password := "qazqazqaz"

		// 2. generate account for which key will get store in keystore
		account, err := store.NewAccount(password)
		require.NoError(t, err)

		log.Info("generated account : ", account.Address.Hex())

	})

	t.Run("importing wallet", func(t *testing.T) {

		// 1. take temp wallet
		store := keystore.NewKeyStore("/tmp", keystore.StandardScryptN, keystore.StandardScryptP)

		// 2. read store
		payload, err := os.ReadFile("./wallets/UTC--2025-05-07T11-41-44.144897440Z--be118bb389917ea17dcafa4ab5af9ac55445bb8c")
		require.NoError(t, err)

		// 3. import account
		password := "qazqazqaz"
		account, err := store.Import(payload, password, password)
		require.NoError(t, err)

		log.Info("imported account : ", account.Address.Hex())

	})

	// t.Run("quering", func(t *testing.T) {

	// 	header, err := client.HeaderByNumber(ctx, nil)
	// 	require.NoError(t, err)

	// 	log.Info("header : ", header.Number.String())

	// })

	t.Run("transfer ether", func(t *testing.T) {

		// fields needed for creating new txn
		var address = common.HexToAddress("0x6de2EC5577D5B6e32D522Fe9409E4fD9167b645F")
		var privateKeyBytes, err = hex.DecodeString("573cd7f19d88a3e3bd824d20802fee89979c1341c73232b3c4e2f57d85d2e688")
		require.NoError(t, err)

		nonce, err := client.PendingNonceAt(ctx, address)
		require.NoError(t, err)

		value := big.NewInt(1000000000000000000) // 1 eth
		gasLimit := uint64(21000)

		gasPrice, err := client.SuggestGasPrice(ctx)
		require.NoError(t, err)

		// create new txn
		var txn = types.NewTransaction(nonce, address, value, gasLimit, gasPrice, nil)

		chainId, err := client.NetworkID(ctx)
		require.NoError(t, err)

		privateKey, err := crypto.ToECDSA(privateKeyBytes)
		require.NoError(t, err)

		txn, err = types.SignTx(txn, types.NewEIP155Signer(chainId), privateKey)
		require.NoError(t, err)

		err = client.SendTransaction(ctx, txn)
		require.NoError(t, err)

	})

	t.Run("quering", func(t *testing.T) {

		header, err := client.HeaderByNumber(ctx, nil)
		require.NoError(t, err)

		log.Info("header : ", header.Number.String())

		block, err := client.BlockByNumber(ctx, big.NewInt(1))
		require.NoError(t, err)

		log.Infof("block details : %+v", block)
		log.Info(block.Number(), block.Time(), block.Difficulty(), block.Hash().Hex())
		count, _ := client.TransactionCount(ctx, block.Hash())
		log.Info("txnCount in block : ", count)

		for _, txn := range block.Transactions() {
			//  get basic txn info
			log.Infof(
				"txn detail in block : hash: %v, value: %v, gas: %v, nonce: %v, toAddress: %v ",
				txn.Hash().Hex(),
				txn.Value().String(),
				txn.Gas(),
				txn.Nonce(),
				txn.To().Hex(),
			)

			//  each txn having recipt. recipt having status (0/1 ie success or failed), return value & logs
			receipt, err := client.TransactionReceipt(ctx, txn.Hash())
			require.NoError(t, err)

			log.Info("receipt", receipt.Status, receipt.Logs, receipt.GasUsed, string(receipt.PostState))

			// get txn by index
			txn, err = client.TransactionInBlock(ctx, block.Hash(), 0)
			require.NoError(t, err)
			log.Info("txn got : ", txn.Hash().String())

			// get txn by hash
			txn, isPending, err := client.TransactionByHash(ctx, txn.Hash())
			require.NoError(t, err)
			log.Infof("txn got : %v, isPending %v", txn.Hash().String(), isPending)

		}

	})

	t.Run("listening events thorugh subscriber", func(t *testing.T) {

		headers := make(chan *types.Header)

		subscription, err := client.SubscribeNewHead(ctx, headers)
		require.NoError(t, err)

		for {
			select {
			case err := <-subscription.Err():
				log.Info(err.Error())
			case header := <-headers:
				log.Info("new block details :", header.Hash())
			}
		}

	})

	t.Run("listening events by polling", func(t *testing.T) {

	})

	t.Run("raw txn", func(t *testing.T) {

		// 1. create signed txn

		address := common.HexToAddress("0x6de2EC5577D5B6e32D522Fe9409E4fD9167b645F")
		privateKeyBytes, _ := hex.DecodeString("573cd7f19d88a3e3bd824d20802fee89979c1341c73232b3c4e2f57d85d2e688")
		nonce, _ := client.PendingNonceAt(ctx, address)
		gasPrice, _ := client.SuggestGasPrice(ctx)
		chainId, _ := client.NetworkID(ctx)
		privateKey, _ := crypto.ToECDSA(privateKeyBytes)

		var txn = types.NewTransaction(
			nonce, address,
			big.NewInt(1000000000000000000),
			uint64(21000), gasPrice, nil,
		)

		txn, _ = types.SignTx(txn, types.NewEIP155Signer(chainId), privateKey)

		// 2. get rawbytes for txn to log it
		rawTxnBytes, err := rlp.EncodeToBytes(txn)
		require.NoError(t, err)

		log.Info(hexutil.Encode(rawTxnBytes))

		// 3. how to convert those rawbytes again it txn. PS We need to send txn not rowbytes

		var txn2 = &types.Transaction{}
		err = rlp.DecodeBytes(rawTxnBytes, &txn2)
		require.NoError(t, err)

		err = client.SendTransaction(ctx, txn2)
		require.NoError(t, err)

	})

}

func WeiToEther(wei *big.Int) *big.Float {
	weiAmountFloat, _ := new(big.Float).SetString(wei.String())
	return new(big.Float).Quo(weiAmountFloat, big.NewFloat(math.Pow10(18)))
}

func EtherToWei(ether *big.Float) *big.Int {
	i, _ := ether.Uint64()
	return new(big.Int).SetUint64(i)
}
