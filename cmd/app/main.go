package main

import (
	"fmt"

	"github.com/TheRealLordtop/Financial-Planner/internal/adapters/storage"
	"github.com/TheRealLordtop/Financial-Planner/internal/core"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Transaction core.Transaction
type Method core.Method

func main() {
	// if err := os.MkdirAll("data", 0755); err != nil {
	// 	log.Fatal(err)
	// }

	db, err := gorm.Open(sqlite.Open("data/app.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// ctx := context.Background()

	// Migrate the schema
	db.AutoMigrate(&Transaction{})

	// err = gorm.G[Transaction](db).Create(ctx, &Transaction{Time: time.Now(), Value: 255, Method: core.Card})

	txs := storage.GetAllTransactions()

	fmt.Println(txs)
}
