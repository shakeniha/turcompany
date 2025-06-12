package models

import (
	"google.golang.org/genproto/googleapis/type/decimal"
	"time"
)

type Deals struct {
	ID        string          `json:"id"`
	LeadID    string          `json:"lead_id"`
	Amount    decimal.Decimal `json:"amount"`
	Currency  string          `json:"currency"`
	Status    string          `json:"status"`
	CreatedAt time.Time       `json:"created_at"`
}
