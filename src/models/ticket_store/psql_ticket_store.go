package ticket_store

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
)

type TicketPostgresStore struct{ db *sql.DB }

func NewTicketPostgresStore(db *sql.DB) *TicketPostgresStore {
	return &TicketPostgresStore{db: db}
}

func (s *TicketPostgresStore) GetTicket(id string) (Ticket, error) {
	var ticket Ticket

	err := s.db.QueryRow(`select id, user_id, pool_id, details, is_hand_picked, is_used, updated_at from tickets where id = $1`, id).
		Scan(&ticket.ID, &ticket.UserID, &ticket.PoolID, &ticket.Details, &ticket.IsHandPicked, &ticket.IsUsed, &ticket.UpdatedAt)

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

func (s *TicketPostgresStore) GetUnusedTickets(poolID string) ([]Ticket, error) {
	rows, err := s.db.Query(`select id, user_id, pool_id, details, is_hand_picked, is_used, updated_at from tickets where is_used = false and pool_id = $1`, poolID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var tickets []Ticket
	for rows.Next() {
        var ticket Ticket
        if err := rows.Scan(&ticket.ID, &ticket.UserID, &ticket.PoolID, &ticket.Details, &ticket.IsHandPicked, &ticket.IsUsed, &ticket.UpdatedAt); err != nil {
            return tickets, err
        }
        tickets = append(tickets, ticket)
    }

	return tickets, nil
}

func (s *TicketPostgresStore) CreateTicket(request CreateTicketRequest) (Ticket, error) {
	var ticket Ticket

	err := s.db.QueryRow(`insert into tickets (user_id, pool_id, details, is_hand_picked) values ($1, $2, $3, $4) returning id, user_id, pool_id, details, is_hand_picked, is_used, updated_at`, request.UserID, request.PoolID, request.Details, request.IsHandPicked).
		Scan(&ticket.ID, &ticket.UserID, &ticket.PoolID, &ticket.Details, &ticket.IsHandPicked, &ticket.IsUsed, &ticket.UpdatedAt)

	if err != nil {
		return ticket, errors.New("unable to create ticket")
	}

	return ticket, nil
}

func (s *TicketPostgresStore) BatchInsertTicket(requests []CreateTicketRequest) error {
	var placeholders []string
	var args []any

	i := 1
	for _, request := range requests {
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d)", i, i+1, i+2, i+3))
		args = append(args, request.UserID, request.PoolID, request.Details, request.IsHandPicked)
		i += 4
	}

	query := fmt.Sprintf("INSERT INTO tickets (user_id, pool_id, details, is_hand_picked) VALUES %s", strings.Join(placeholders, ","))

	_, err := s.db.Exec(query, args...)
	if err != nil {
		return errors.New("unable to create ticket: " + err.Error())
	}

	return nil
}
