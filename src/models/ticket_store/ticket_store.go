package ticket_store

type TicketStore interface {
	GetTicket(id string) (Ticket, error)
	CreateTicket(request CreateTicketRequest) (Ticket, error)
	GetUserTicketsCount(user_id string, pool_id string) (int, error)
}
