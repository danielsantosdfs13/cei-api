package ceitransaction

import (
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
