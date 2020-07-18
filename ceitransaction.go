package ceiapi

import (
	"fmt"
	"time"
)

type CeiTransaction struct {
	Institution     string
	Account         string
	Date            time.Time
	Operation       string
	Category        string
	Expiration      time.Time
	Code            string
	Quantity        int
	Price           float32
	TotalValue      float32
	QuotationFactor int
}

type CeiTransactions struct {
	Entries []CeiTransaction
	Months  CeiTransactionMonth
}

type CeiTransactionMonth struct {
	Categories                []string
	CeiTransactionMonthSeries []CeiTransactionMonthSeries
}

type CeiTransactionMonthSeries struct {
	Name string
	Data []float32
}

func New() {
	fmt.Println("r1")
}
