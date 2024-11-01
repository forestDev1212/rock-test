# Ethereum Block Query and Storage Application

This project is a technical assessment to demonstrate skills in Golang, Docker, and PostgreSQL by querying Ethereum blockchain data, filtering for specific contract transactions, and storing the data in a PostgreSQL database. Additionally, Redis is used as a caching layer, and an API endpoint allows querying transaction data by wallet address.


## Project Overview

This application connects to the Ethereum blockchain using a free RPC endpoint, starting from a specified block height (e.g., 100,000). It queries for `mint` transactions from a specific contract and stores this data in a PostgreSQL database. Additionally, Redis is used as a cache layer to optimize querying, and a REST API is provided for retrieving transactions by wallet address.

---

## Prerequisites

- **Go** (v1.22)
- **Docker** and **Docker Compose**
- **Git** for version control

---

## Setup and Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/forestDev1212/rocks-test.git
   cd rocks-test
2. **Create a .env File**
```RPC_URL=https://your-free-ethereum-rpc-url
DATABASE_URL=postgres://user:password@postgres:5432/ethereum_data?sslmode=disable
REDIS_ADDR=redis:6379
REDIS_PASSWORD=your_redis_password
CONTRACT_ADDRESS=0x047d41f2544b7f63a8e991af2068a363d210d6da
```

3. **Running the Project**
```docker-compose up --build```

4. **Application Structure**

```
ethereum-query-assessment/
├── blockchain/             # Ethereum-related code
├── config/                 # Configuration setup
├── database/               # Database models and setup
├── Dockerfile              # Dockerfile for the Go app
├── docker-compose.yml      # Docker Compose file
├── go.mod                  # Go modules dependencies
├── go.sum                  # Go modules checksum
└── main.go                 # Main application entry point
```

