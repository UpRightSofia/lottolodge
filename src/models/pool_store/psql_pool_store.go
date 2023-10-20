package pool_store

import (
	"database/sql"
	"errors"
	"log"
)

type PoolPostgresStore struct{ db *sql.DB }

func NewPoolPostgresStore(db *sql.DB) *PoolPostgresStore {
	return &PoolPostgresStore{db: db}
}

func (s *PoolPostgresStore) GetPool(id string) (Pool, error) {
	var pool Pool

	err := s.db.QueryRow(`select id, details, updated_at from pools where id = $1`, id).Scan(&pool.ID, &pool.Details, &pool.UpdatedAt)

	if err != nil {
		log.Println(err)
		return pool, err
	}

	return pool, nil
}

func (s *PoolPostgresStore) CreatePool(request CreatePoolRequest) (Pool, error) {
	var pool Pool

	err := s.db.QueryRow(`insert into pools (details) values ($1) returning id, details, updated_at`, request.Details).Scan(&pool.ID, &pool.Details, &pool.UpdatedAt)

	if err != nil {
		return pool, errors.New("unable to create pool")
	}

	return pool, nil
}
