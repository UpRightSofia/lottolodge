package pool_store

import (
	"database/sql"
	"time"
)

type Pool struct {
	ID        string         `json:"id"`
	Details   sql.NullString `json:"details"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type CreatePoolRequest struct {
	Details string `json:"details"`
}
