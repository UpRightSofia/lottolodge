package pool_store

import (
	"database/sql"
	"log"
)

type PoolPostgresStore struct{ db *sql.DB }

func NewPoolPostgresStore(db *sql.DB) *PoolPostgresStore {
	return &PoolPostgresStore{db: db}
}

func (s *PoolPostgresStore) GetPool(id string) (Pool, error) {
	var pool Pool

	err := s.db.QueryRow(`select id, details, created_at, is_active from pools where id = $1`, id).Scan(&pool.ID, &pool.Details, &pool.CreatedAt, &pool.IsActive)

	if err != nil {
		log.Println(err)
		return pool, err
	}

	return pool, nil
}

func (s *PoolPostgresStore) GetActivePool() (Pool, error) {
	var pool Pool
	err := s.db.QueryRow(`SELECT id, details, created_at, is_active FROM pools WHERE is_active = true ORDER BY created_at LIMIT 1`).Scan(&pool.ID, &pool.Details, &pool.CreatedAt, &pool.IsActive)

	if err != nil {
		log.Println(err)
		return pool, err
	}

	return pool, nil
}

func (s *PoolPostgresStore) CreatePool() (Pool, error) {
	var pool Pool

	err := s.db.QueryRow(`insert into pools (details) values (null) returning id, details, created_at, is_active`).Scan(&pool.ID, &pool.Details, &pool.CreatedAt, &pool.IsActive)

	if err != nil {
		return pool, err
	}

	return pool, nil
}

func (s *PoolPostgresStore) MarkPoolAsDone(id string, details string) (Pool, error) {

	pool, err := s.GetPool(id)
	if err != nil {
		return pool, err
	}

	var updatedPool Pool
	err = s.db.QueryRow(`UPDATE pools SET is_active = false, details = $1 WHERE id = $2 RETURNING id, details, created_at, is_active`, details, pool.ID).Scan(&updatedPool.ID, &updatedPool.Details, &updatedPool.CreatedAt, &updatedPool.IsActive)
	if err != nil {
		return updatedPool, err
	}

	return updatedPool, nil
}

func (s *PoolPostgresStore) GetLastCompleted() (Pool, error) {
	var pool Pool
	err := s.db.QueryRow(`select id, details, created_at, is_active from pools where is_active = false order by created_at desc limit 1`).Scan(&pool.ID, &pool.Details, &pool.CreatedAt, &pool.IsActive)
	if err != nil {
		return pool, err
	}

	return pool, nil
}