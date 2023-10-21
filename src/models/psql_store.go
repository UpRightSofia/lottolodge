package models

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/UpRightSofia/lottolodge/src/models/config"
	"github.com/UpRightSofia/lottolodge/src/models/pool_store"
	"github.com/UpRightSofia/lottolodge/src/models/ticket_store"
	"github.com/UpRightSofia/lottolodge/src/models/user_store"
	"github.com/UpRightSofia/lottolodge/src/models/winning_store"
)

func NewPostgresStore(dbConfig config.DBConfig, testBlock func(*PostgresStore)) {
	db, err := sql.Open("pgx", dbConfig.GetDSN())
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()

	testBlock(
		&PostgresStore{
			db:        db,
			UserStore: user_store.NewUserPostgresStore(db),
		},
	)
}

type PostgresStore struct {
	db           *sql.DB
	UserStore    *user_store.UserPostgresStore
	PoolStore    *pool_store.PoolPostgresStore
	TicketStore  *ticket_store.TicketPostgresStore
	WinningStore *winning_store.WinningPostgresStore
}

func (s *PostgresStore) GetDB() *sql.DB {
	return s.db
}
