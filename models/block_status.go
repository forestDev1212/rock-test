package models

import (
	"rocks-test/database"
)

type BlockStatus struct {
	ID              int64  `db:"id"`
	LastBlockNumber uint64 `db:"last_block_number"`
}

func GetLastProcessedBlock() (uint64, error) {
	var blockStatus BlockStatus
	err := database.DB.Get(&blockStatus, "SELECT last_block_number FROM block_status WHERE id = 1")
	if err != nil {
		return 0, err
	}
	return blockStatus.LastBlockNumber, nil
}

func UpdateLastProcessedBlock(blockNumber uint64) error {
	_, err := database.DB.Exec("UPDATE block_status SET last_block_number = $1 WHERE id = 1", blockNumber)
	return err
}
