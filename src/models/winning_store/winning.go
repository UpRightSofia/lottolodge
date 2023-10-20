package winning_store

import (
	"time"
)

type Winning struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	TicketID  string    `json:"ticket_id"`
	PoolID    string    `json:"pool_id"`
	PrizeE5   int64     `json:"prize_e5"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateWinningRequest struct {
	UserID   string `json:"user_id"`
	TicketID string `json:"ticket_id"`
	PoolID   string `json:"pool_id"`
	PrizeE5  int64  `json:"prize_e5"`
}
