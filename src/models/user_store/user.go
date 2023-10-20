package user_store

type User struct {
	ID string `json:"id"`
}

type CreateUserRequest struct {
	ID string `json:"id"`
}
