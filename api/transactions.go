package api

import (
	"net/http"

	"rocks-test/models"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("api/v1/transactions/:walletAddress", GetTransactionsByWallet)
}

func GetTransactionsByWallet(c *gin.Context) {
	walletAddress := c.Param("walletAddress")

	// Here, fetch the transactions based on the wallet address
	transactions, err := models.GetTransactions(walletAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching transactions"})
		return
	}
	c.JSON(http.StatusOK, transactions)
}
