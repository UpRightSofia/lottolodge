package pool_store

import (
	"database/sql"
	"time"
)


type Pool struct {
	ID        string         `json:"id"`
	Details   sql.NullString `json:"details"`
	CreatedAt time.Time      `json:"created_at"`
	IsActive  bool           `json:"is_active"`
}

type CreatePoolRequest struct {
	Details string `json:"details"`
}
