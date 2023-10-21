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

	log.Println("request: ", request)

	err := s.db.QueryRow(`insert into winnings (user_id, ticket_id, pool_id, prize_e5) values ($1, $2, $3, $4) returning id, user_id, ticket_id, pool_id, prize_e5, updated_at`,
		request.UserID, request.TicketID, request.PoolID, request.PrizeE5).Scan(&winning.ID, &winning.UserID, &winning.TicketID, &winning.PoolID, &winning.PrizeE5, &winning.UpdatedAt)

	if err != nil {
		log.Println(err)
		return winning, err
	}

	return winning, nil
}

func (s *WinningPostgresStore) GetWinningsForUserAndPool(user_id string, pool_id string) ([]Winning, error) {
	rows, err := s.db.Query(`select id, user_id, ticket_id, pool_id, prize_e5, updated_at from winnings where user_id = $1 and pool_id = $2`, user_id, pool_id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var winnings []Winning
	for rows.Next() {
		var winning Winning
		if err := rows.Scan(&winning.ID, &winning.UserID, &winning.TicketID, &winning.PoolID, &winning.PrizeE5, &winning.UpdatedAt); err != nil {
			return winnings, err
		}
		winnings = append(winnings, winning)
	}

	return winnings, nil
}

func (s *WinningPostgresStore) GetUserWinnings(userID string) ([]UserWinning, error) {
	query := `
WITH last_50_pools AS (
    SELECT id, created_at
    FROM pools
	WHERE is_active = false
    ORDER BY created_at DESC 
    LIMIT 50
)

SELECT DISTINCT ON(p.id)
    p.id AS pool_id,
	p.created_at as pool_date,
    COALESCE(SUM(w.prize_e5), 0) AS total_prize_e5
FROM 
    last_50_pools p
LEFT JOIN 
    winnings w ON p.id = w.pool_id AND w.user_id = $1
GROUP BY 
    p.id, p.created_at
ORDER BY 
    p.id;
`

	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var winnings []UserWinning
	for rows.Next() {
		userWinning := UserWinning{}
		err := rows.Scan(&userWinning.PoolID, &userWinning.PoolDate, &userWinning.TotalWinE5)
		if err != nil {
			return nil, err
		}
		winnings = append(winnings, userWinning)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return winnings, nil
}
