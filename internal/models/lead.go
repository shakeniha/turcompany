package models

import (
	"time"
)

type Leads struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	OwnerID     int       `json:"owner_id"`
	Status      string    `json:"status"`
}
