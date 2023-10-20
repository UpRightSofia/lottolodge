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
	UpdatedAt    time.Time      `json:"updated_at"`
}

type CreateTicketRequest struct {
	UserID       string         `json:"user_id"`
	Details      sql.NullString `json:"details"` // Note: Consider using a custom type or a library that can handle JSON in Go for more advanced use-cases.
	IsHandPicked bool           `json:"is_hand_picked"`
}
