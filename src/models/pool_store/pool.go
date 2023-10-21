package pool_store

import (
	"database/sql"
	"time"
)

type PoolDetails struct {
	DrawnNumbers   []int `json:"drawn_numbers"`
	TwoMultiplier  int   `json:"two_multiplier"`
	FiveMultiplier int   `json:"five_multiplier"`
}

type Pool struct {
	ID        string         `json:"id"`
	Details   sql.NullString `json:"details"`
	CreatedAt time.Time      `json:"created_at"`
	IsActive  bool           `json:"is_active"`
}

type CreatePoolRequest struct {
	Details string `json:"details"`
}
