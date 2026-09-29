package models


type Event struct {

	ID string

	Name string

	Image string

	Date string

	Venue string

	Description string

	TicketURL string

}

type EventResponse struct {

	Music []Event

	Sports []Event

	Error error

}