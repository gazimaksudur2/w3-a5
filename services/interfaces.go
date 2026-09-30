package services

import (
	"context"
	"event-explorer/models"
)

type EventProvider interface {
	FetchEvents(
		ctx context.Context,
		city string,
		country string,
		category string,
	) ([]models.Event, error)

	GetEventDetails(
		ctx context.Context,
		eventID string,
	) (models.Event, error)
}
