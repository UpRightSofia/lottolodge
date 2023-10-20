package user_store

import (
	"database/sql"
	"errors"
	"log"
)

type UserPostgresStore struct{ db *sql.DB }

func NewUserPostgresStore(db *sql.DB) *UserPostgresStore {
	return &UserPostgresStore{db: db}
}

func (s *UserPostgresStore) GetUser(id string) (User, error) {
	var user User

	err := s.db.QueryRow(`select id, updated_at, balance_e5 from users where id = $1`, id).Scan(&user.ID, &user.UpdatedAt, &user.BalanceE5)

	if err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}

func (s *UserPostgresStore) CreateUser(request CreateUserRequest) (User, error) {
	var user User

	err := s.db.QueryRow(`insert into users (id, balance_e5) values ($1, $2) returning id, updated_at, balance_e5`, request.ID, request.BalanceE5).Scan(&user.ID, &user.UpdatedAt, &user.BalanceE5)

	if err != nil {
		return user, errors.New("unable to create user")
	}

	return user, nil
}
