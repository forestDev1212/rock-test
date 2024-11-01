package main

import (
	"context"
	"log"
	"rocks-test/api"
	"rocks-test/blockchain"
	"rocks-test/config"
	"rocks-test/database"
	"rocks-test/models"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize configuration
	config.LoadEnv()

	// Setup database and Redis
	database.SetupPostgres()
	database.SetupRedis()

	// Start the block fetcher in a separate goroutine
	go StartBlockFetcher()

	// Initialize the Gin router
	r := gin.Default()
	api.RegisterRoutes(r) // Register API routes

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// StartBlockFetcher starts a goroutine that fetches new blocks periodically
func StartBlockFetcher() {
	client := blockchain.GetBlockchainClient()
	ticker := time.NewTicker(15 * time.Second) // Adjust the interval as needed
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			lastProcessedBlock, err := models.GetLastProcessedBlock()
			if err != nil {
				log.Printf("Error getting last processed block: %v", err)
				continue
			}

			// Get the latest block number from the blockchain
			latestBlockHeader, err := client.HeaderByNumber(context.Background(), nil)
			if err != nil {
				log.Printf("Error fetching latest block header: %v", err)
				continue
			}
			latestBlockNumber := latestBlockHeader.Number.Uint64()

			if latestBlockNumber > lastProcessedBlock {
				// Fetch transactions from lastProcessedBlock+1 to latestBlockNumber
				log.Printf("Latest Block Number: %d", latestBlockNumber)
				log.Printf("Last Processed Block: %d", lastProcessedBlock)

				err := blockchain.FetchTransactions(lastProcessedBlock+1, latestBlockNumber)
				if err != nil {
					log.Printf("Error fetching transactions: %v", err)
					continue
				}

				// Update the last processed block number
				err = models.UpdateLastProcessedBlock(latestBlockNumber)
				if err != nil {
					log.Printf("Error updating last processed block: %v", err)
					continue
				}
			} else {
				log.Println("No new blocks to process.")
			}
		}
	}
}
