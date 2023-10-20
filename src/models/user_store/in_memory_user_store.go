package user_store

import (
	"errors"
)

type InMemoryUserStore struct {
	users map[string]User
}

func NewInMemoryUserStore(userMap map[string]User) *InMemoryUserStore {
	return &InMemoryUserStore{users: userMap}
}

func (s *InMemoryUserStore) GetUser(id string) (User, error) {
	user, ok := s.users[id]
	if !ok {
		return User{}, errors.New("User not found")
	}

	return user, nil
}

func (s *InMemoryUserStore) CreateUser(request CreateUserRequest) (User, error) {
	user := User{
		ID: request.ID,
	}

	s.users[request.ID] = user

	return user, nil
}
