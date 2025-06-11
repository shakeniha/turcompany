package models

import (
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/genproto/googleapis/type/decimal"
)
type Deals struct {
	ID     string `json:"id"`
	LeadID string `json:"lead_id"`
	Amount decimal.Decimal `json:"amount"`
	Currency string `json:"currency"`
	Status string `json:"status"`
	CreatedAt date.Date `json:"created_at"`
}