package util

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"path/filepath"

	"os"

	"github.com/gofiber/fiber/v2/log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type orgs map[string]common.Address
type keys map[string]*ecdsa.PrivateKey

var (
	orgMap = make(orgs)
	keyMap = make(keys)
)

func GetAddress(name string) common.Address {
	return orgMap[name]
}

func GetPrivateKey(name string) *ecdsa.PrivateKey {
	return keyMap[name]
}

func GetNeworkId() string {
	return ""
}

func GetChaincodeId() string {
	return ""
}

func InitForGeth(names []string) error {
	keystorePath := "./../tmp/data/keystore"
	password := "password1"

	files, err := os.ReadDir(keystorePath)
	if err != nil {
		log.Info("Failed to read keystore directory: %v", err)
		return err
	}

	var i = 0

	for _, file := range files {
		if !file.IsDir() {
			// Read the keystore file
			keystoreFilePath := filepath.Join(keystorePath, file.Name())
			keyJSON, err := os.ReadFile(keystoreFilePath)
			if err != nil {
				log.Info("Failed to read keystore file %s: %v", file.Name(), err)
				continue
			}

			// Decrypt the key with the password
			key, err := keystore.DecryptKey(keyJSON, password)
			if err != nil {
				log.Info("Failed to decrypt keystore file %s: %v", file.Name(), err)
				continue
			}

			orgMap[names[i]] = key.Address
			keyMap[names[i]] = key.PrivateKey
			i++
		}
	}

	if len(orgMap) != len(names) {
		return fmt.Errorf("length not matched")
	}

	return nil
}

func InitForAnvil(names []string) error {
	payload, err := os.ReadFile("./../orgs.json")
	if err != nil {
		log.Info("error opening file:", err)
		return err
	}

	var accounts struct {
		AvailableAccounts []string `json:"available_accounts"`
		PrivateKeys       []string `json:"private_keys"`
	}

	err = json.Unmarshal(payload, &accounts)
	if err != nil {
		log.Info("error decoding JSON:", err)
		return err
	}

	if len(names) != len(accounts.AvailableAccounts) {
		return fmt.Errorf("len(names) != len(available_accounts)")
	}

	for i := range accounts.AvailableAccounts {
		addr := common.HexToAddress(accounts.AvailableAccounts[i])
		orgMap[names[i]] = addr

		key := accounts.PrivateKeys[i]
		key = key[2:] // strip "0x"

		privKeyBytes, err := hex.DecodeString(key)
		if err != nil {
			return err
		}
		privKey, err := crypto.ToECDSA(privKeyBytes)
		if err != nil {
			return err
		}
		keyMap[names[i]] = privKey
	}

	return nil
}

func InitForGanache(names []string) error {
	payload, err := os.ReadFile("./../orgs.json")
	if err != nil {
		log.Info("error opening file:", err)
		return err
	}

	var accounts struct {
		Addresses   map[string]any    `json:"addresses"`
		PrivateKeys map[string]string `json:"private_keys"`
	}
	err = json.Unmarshal(payload, &accounts)
	if err != nil {
		log.Info("error decoding JSON:", err)
		return err
	}

	if len(names) != len(accounts.Addresses) {
		return fmt.Errorf("len(names) != len(accounts.Accounts)")
	}

	var j = 0
	for i := range accounts.Addresses {
		// Convert address string to common.Address
		addr := common.HexToAddress(i)
		orgMap[names[j]] = addr

		key := accounts.PrivateKeys[i]
		key = key[2:]

		// Convert hex private key to *ecdsa.PrivateKey
		privKeyBytes, err := hex.DecodeString(key)
		if err != nil {
			return err
		}
		privKey, err := crypto.ToECDSA(privKeyBytes)
		if err != nil {
			return err
		}
		keyMap[names[j]] = privKey
		j++
	}

	return nil
}
