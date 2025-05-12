package entity

type Park struct {
	ID       int
	CarID    int
	TicketID string
}

type Ticket struct {
	ID       int
	CarID    string
	ParkID   int
	TicketID string
}
