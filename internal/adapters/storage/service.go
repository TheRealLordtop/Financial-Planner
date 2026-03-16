package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/TheRealLordtop/Financial-Planner/internal/core"
)

type Transaction core.Transaction

func GetAllTransactions() []Transaction {
	db, err := gorm.Open(sqlite.Open("data/app.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// ctx := context.Background()

	var transactions []Transaction
	result := db.Find(&transactions)

	return transactions
}

func AddTransactions([]Transaction) {
	db, err := gorm.Open(sqlite.Open("data/app.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}
