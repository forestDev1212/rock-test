package blockchain

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"

	"rocks-test/config"
	"rocks-test/models"
)

func FetchTransactions(fromBlock, toBlock uint64) error {
	client := GetBlockchainClient() // Ensure this returns *ethclient.Client

	// Use the contract address from the config
	contractAddress := common.HexToAddress(config.ContractAddress)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   big.NewInt(int64(toBlock)),
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Printf("Failed to fetch logs: %v", err)
		return err
	}

	fmt.Printf("Logs Length: %d\n", len(logs))

	for _, vLog := range logs {
		tx, isPending, err := client.TransactionByHash(context.Background(), vLog.TxHash)
		if err != nil {
			log.Printf("Failed to fetch transaction by hash: %v", err)
			continue
		}
		if isPending {
			log.Println("Transaction is still pending; skipping.")
			continue
		}

		// Get the chain ID
		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Printf("Failed to get network ID: %v", err)
			continue
		}

		// Get the signer
		signer := ethTypes.LatestSignerForChainID(chainID)

		// Get the sender address
		sender, err := ethTypes.Sender(signer, tx)
		if err != nil {
			log.Printf("Failed to get sender from transaction: %v", err)
			continue
		}

		// Get the transaction receipt
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil || receipt == nil {
			log.Printf("Failed to get transaction receipt for tx %s: %v", tx.Hash().Hex(), err)
			continue
		}

		// Get the recipient address
		recipient := tx.To()
		if recipient == nil {
			if receipt.ContractAddress != (common.Address{}) {
				recipient = &receipt.ContractAddress
			} else {
				log.Printf("Transaction recipient address is nil for tx %s", tx.Hash().Hex())
				continue
			}
		}

		// Save the transaction with the required details
		err = models.SaveTransaction(tx, sender, *recipient, receipt.BlockNumber)
		if err != nil {
			log.Printf("Failed to save transaction: %v", err)
		} else {
			fmt.Println("Transaction saved successfully.")
		}
		time.Sleep(10 * time.Millisecond)
	}

	return nil
}
