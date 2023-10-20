package winning_store

import (
	"database/sql"
	"log"
)

type WinningPostgresStore struct{ db *sql.DB }

func NewWinningPostgresStore(db *sql.DB) *WinningPostgresStore {
	return &WinningPostgresStore{db: db}
}

func (s *WinningPostgresStore) GetWinning(id string) (Winning, error) {
	var winning Winning

	err := s.db.QueryRow(`select id, user_id, ticket_id, pool_id, prize_e5, updated_at from winnings where id = $1`, id).Scan(&winning.ID, &winning.UserID, &winning.TicketID, &winning.PoolID, &winning.PrizeE5, &winning.UpdatedAt)

	if err != nil {
		log.Println(err)
		return winning, err
	}

	return winning, nil
}

func (s *WinningPostgresStore) CreateWinning(request CreateWinningRequest) (Winning, error) {
	var winning Winning

	err := s.db.QueryRow(`insert into winnings (user_id, ticket_id, pool_id, prize_e5) values ($1, $2, $3, $4) returning id, user_id, ticket_id, pool_id, prize_e5, updated_at`,
		request.UserID, request.TicketID, request.PoolID, request.PrizeE5).Scan(&winning.ID, &winning.UserID, &winning.TicketID, &winning.PoolID, &winning.PrizeE5, &winning.UpdatedAt)

	if err != nil {
		log.Println(err)
		return winning, err
	}

	return winning, nil
}
