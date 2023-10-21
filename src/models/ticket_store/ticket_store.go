package ticket_store

type TicketStore interface {
	GetTicket(id string) (Ticket, error)
	GetUnusedTickets(poolID string) ([]Ticket, error)
	CreateTicket(request CreateTicketRequest) (Ticket, error)
	BatchInsertTicket(requests []CreateTicketRequest) error
	GetUserTicketsCount(user_id string, pool_id string) (int, error)
	GetTicketsForUser(user_id string, pool_id string) ([]Ticket, error)
}
