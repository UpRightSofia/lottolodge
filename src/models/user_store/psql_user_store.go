package user_store

import (
	"database/sql"
	"errors"
)

type UserPostgresStore struct{ db *sql.DB }

func NewUserPostgresStore(db *sql.DB) *UserPostgresStore {
	return &UserPostgresStore{db: db}
}

func (s *UserPostgresStore) GetUser(id string) (User, error) {
	var user User

	err := s.db.QueryRow(`select id from users where id = $1`, id).Scan(&user.ID)

	if err != nil {
		return user, errors.New("unable to fetch user")
	}

	return user, nil
}

func (s *UserPostgresStore) CreateUser(request CreateUserRequest) (User, error) {
	var user User

	err := s.db.QueryRow(`insert into users (id) values ($1) returning id`, request.ID).Scan(&user.ID)

	if err != nil {
		return user, errors.New("unable to create user")
	}

	return user, nil
}
