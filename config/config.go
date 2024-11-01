package config

import (
	"os"
)

var (
	RPCUrl          string
	DatabaseURL     string
	RedisAddr       string
	RedisPassword   string
	ContractAddress string
)

func LoadEnv() {
	// Environment variables are already set in the Docker container.
	// No need to use godotenv.Load()
	RPCUrl = os.Getenv("RPC_URL")
	DatabaseURL = os.Getenv("DATABASE_URL")
	RedisAddr = os.Getenv("REDIS_ADDR")
	RedisPassword = os.Getenv("REDIS_PASSWORD")
	ContractAddress = os.Getenv("CONTRACT_ADDRESS")
}
