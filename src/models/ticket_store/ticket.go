package ticket_store

import (
	"database/sql"
	"time"
)

type Ticket struct {
	ID           string         `json:"id"`
	UserID       string         `json:"user_id"`
	Details      sql.NullString `json:"details"` // Note: Consider using a custom type or a library that can handle JSON in Go for more advanced use-cases.
	IsHandPicked bool           `json:"is_hand_picked"`
	IsUsed       bool           `json:"is_used"`
	UpdatedAt    time.Time      `json:"updated_at"`
	PoolID       string         `json:"pool_id"`
}

type CreateTicketRequest struct {
	UserID       string `json:"user_id"`
	Details      string `json:"details"` // Note: Consider using a custom type or a library that can handle JSON in Go for more advanced use-cases.
	IsHandPicked bool   `json:"is_hand_picked"`
	PoolID       string `json:"pool_id"`
}
