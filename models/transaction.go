package models

import (
	"log"
	"math/big"
	"rocks-test/database"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type TransactionData struct {
	ID               int64  `db:"id" json:"id"`
	BlockNumber      int64  `db:"block_number" json:"block_number"`
	TransactionHash  string `db:"transaction_hash" json:"transaction_hash"`
	SenderAddress    string `db:"sender_address" json:"sender"`
	RecipientAddress string `db:"recipient_address" json:"recipient"`
	Value            string `db:"value" json:"value"`
	GasPrice         string `db:"gas_price" json:"gas_price"`
	GasLimit         int64  `db:"gas_limit" json:"gas_limit"`
	Nonce            int64  `db:"nonce" json:"nonce"`
	Data             string `db:"data" json:"data"`
}

func SaveTransaction(tx *types.Transaction, sender, recipient common.Address, blockNumber *big.Int) error {
	query := `
    INSERT INTO transactions (block_number, transaction_hash, sender_address, recipient_address, value, gas_price, gas_limit, nonce, data)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
    ON CONFLICT (transaction_hash) DO NOTHING;`

	_, err := database.DB.Exec(query,
		blockNumber.Uint64(),        // Block number from the block context
		tx.Hash().Hex(),             // Transaction hash
		sender.Hex(),                // Sender address
		recipient.Hex(),             // Recipient address
		tx.Value().String(),         // Value as string
		tx.GasPrice().String(),      // Gas price as string
		tx.Gas(),                    // Gas limit
		tx.Nonce(),                  // Nonce
		common.Bytes2Hex(tx.Data()), // Data in hex format
	)

	return err
}

// GetTransactions retrieves all transactions for a given wallet address
func GetTransactions(walletAddress string) ([]TransactionData, error) {
	var transactions []TransactionData

	query := `
	SELECT id, block_number, transaction_hash, sender_address, recipient_address, value, gas_price, gas_limit, nonce, data
	FROM transactions
	WHERE sender_address = $1 OR recipient_address = $1
	ORDER BY block_number DESC
	LIMIT 100;`

	err := database.DB.Select(&transactions, query, walletAddress)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func CountTransactions() (int, error) {
	var count int
	err := database.DB.Get(&count, "SELECT COUNT(*) FROM transactions")
	if err != nil {
		log.Printf("Error counting transactions: %v", err)
		return 0, err
	}
	return count, nil
}
