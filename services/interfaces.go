package services

import "event-explorer/models"

type EventProvider interface {
	FetchEvents(
		city string,
		country string,
		category string,
	) ([]models.Event, error)

	GetEventDetails(
		eventID string,
	) (models.Event, error)
}
