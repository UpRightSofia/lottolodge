package models

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/UpRightSofia/lottolodge/src/models/config"
)

func NewPostgresStore(dbConfig config.DBConfig, testBlock func(*PostgresStore)) {
	db, err := sql.Open("pgx", dbConfig.GetDSN())
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()

	testBlock(
		&PostgresStore{
			db: db,
			// HudStore:  hud_store.NewHudPostgresStore(db),
			// UserStore: user_store.NewUserPostgresStore(db),
		},
	)
}

type PostgresStore struct {
	db *sql.DB
	// HudStore  *hud_store.HudPostgresStore
	// UserStore *user_store.UserPostgresStore
}

func (s *PostgresStore) GetDB() *sql.DB {
	return s.db
}
