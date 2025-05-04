package entity

type Park struct {
	ID          uint64
	Name        string
	IsAvailable bool
}

type Ticket struct {
	ID       uint64
	CarID    string
	ParkID   uint64
	TicketID string
}
