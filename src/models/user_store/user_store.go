package user_store

type UserStore interface {
	GetUser(id string) (User, error)
	CreateUser(request CreateUserRequest) (User, error)
}
