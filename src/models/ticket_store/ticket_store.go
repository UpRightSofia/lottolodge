package ticket_store

type SettingStore interface {
	GetTicket(id string) (Ticket, error)
	CreateTicket(request CreateTicketRequest) (Ticket, error)
}
