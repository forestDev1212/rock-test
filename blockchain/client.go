package blockchain

import (
	"log"

	"rocks-test/config"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBlockchainClient() *ethclient.Client {
	client, err := ethclient.Dial(config.RPCUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	return client
}
