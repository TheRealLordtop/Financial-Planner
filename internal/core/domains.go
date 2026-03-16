package core

import "time"

type Method byte
type Currency byte

const (
	Card Method   = 0
	Cash Method   = 1
	EUR  Currency = 0
	USD  Currency = 1
)

type Transaction struct {
	ID     int32
	Time   time.Time
	Value  int32
	Method Method
}

type Balance struct {
	CardValue int64
	CashValue int64
}

type HolderInfo struct {
	Currency Currency
	Balance  Balance
}
