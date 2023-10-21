package ticket_store

import (
	"database/sql"
	"errors"
	"log"
)

type TicketPostgresStore struct{ db *sql.DB }

func NewTicketPostgresStore(db *sql.DB) *TicketPostgresStore {
	return &TicketPostgresStore{db: db}
}

func (s *TicketPostgresStore) GetTicket(id string) (Ticket, error) {
	var ticket Ticket

	err := s.db.QueryRow(`select id, user_id, details, is_hand_picked, updated_at from tickets where id = $1`, id).Scan(&ticket.ID, &ticket.UserID, &ticket.Details, &ticket.IsHandPicked, &ticket.UpdatedAt)

	if err != nil {
		log.Println(err)
		return ticket, err
	}

	return ticket, nil
}

func (s *TicketPostgresStore) GetUserTicketsCount(user_id string, pool_id string) (int, error) {
	var count int

	err := s.db.QueryRow(`SELECT count(*) FROM tickets WHERE user_id = $1 AND pool_id = $2`, user_id, pool_id).Scan(&count)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return count, nil
}

func (s *TicketPostgresStore) CreateTicket(request CreateTicketRequest) (Ticket, error) {
	var ticket Ticket

	err := s.db.QueryRow(`insert into tickets (user_id, details, is_hand_picked) values ($1, $2, $3) returning id, user_id, details, is_hand_picked, updated_at`, request.UserID, request.Details, request.IsHandPicked).Scan(&ticket.ID, &ticket.UserID, &ticket.Details, &ticket.IsHandPicked, &ticket.UpdatedAt)

	if err != nil {
		return ticket, errors.New("unable to create ticket")
	}

	return ticket, nil
}
