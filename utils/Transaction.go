package utils

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CheckNetworkSettings(networkName string) bool {
	selectedNetwork := getNetworkConfig(networkName)

	if selectedNetwork.Provider == "" || selectedNetwork.Provider == "YOUR_JSON_RPC_PROVIDER" {
		return false
	}

	return true
}

func GetTransactionData(txHash common.Hash, networkName string) ([]uint8, error) {
	selectedNetwork := getNetworkConfig(networkName)

	client, err := ethclient.Dial(selectedNetwork.Provider)
	if err != nil {
		panic("error occured while trying to connect to provider")
	}
	ColoredPrint("[+] provider connection successful!", PrintColors.GREEN)

	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		panic("error occured while querying the transaction")
	}

	return tx.Data(), nil
}

func getNetworkConfig(networkName string) Network {
	jsonFile, err := os.Open(CONFIG_FILE)
	if err != nil {
		panic("error occured while opening config file")
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic("error occured while reading config file")
	}

	var config ConfigFile

	json.Unmarshal(byteValue, &config)

	return config.Networks[networkName]
}
