package models

type Event struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Image string `json:"image"`

	Date string `json:"date"`

	Venue string `json:"venue"`

	Description string `json:"description"`

	TicketURL string `json:"ticketUrl"`
}

type EventResponse struct {
	Music []Event

	Sports []Event

	Errors []error
}
