package user_store

import "time"

// sql: doesnt do anything they are for documentation purposes
type User struct {
	ID        string    `json:"id" sql:"id,type:uuid"`
	UpdatedAt time.Time `json:"updated_at" sql:"updated_at,type:timestamp"`
	BalanceE5 int64     `json:"balance_e5" sql:"balance_e5,type:bigint"`
}

type CreateUserRequest struct {
	ID        string `json:"id"`
	BalanceE5 int64  `json:"balance_e5"`
}
