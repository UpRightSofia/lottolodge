package ticket_store

type TicketStore interface {
	GetTicket(id string) (Ticket, error)
	CreateTicket(request CreateTicketRequest) (Ticket, error)
}
