package services

import (
	"context"
	"fmt"
	"testing"

	"event-explorer/models"
)

type MockEventProvider struct{}

func (m MockEventProvider) FetchEvents(
	ctx context.Context,
	city string,
	country string,
	category string,
) ([]models.Event, error) {

	return []models.Event{

		{
			ID:   "1",
			Name: category + " Event",
		},
	}, nil

}

func (m MockEventProvider) GetEventDetails(
	ctx context.Context,
	id string,
) (models.Event, error) {

	return models.Event{

		ID: id,
	}, nil

}

func TestGetCityEventsSuccess(
	t *testing.T,
) {

	provider :=
		MockEventProvider{}

	result, err :=
		GetCityEvents(
			context.Background(),
			provider,
			"Toronto",
			"CA",
		)

	if err != nil {

		t.Fatal(err)

	}

	if len(result.Music) == 0 {

		t.Fatal(
			"music events missing",
		)

	}

	if len(result.Sports) == 0 {

		t.Fatal(
			"sports events missing",
		)

	}

}

type FailingEventProvider struct{}

func (f FailingEventProvider) FetchEvents(
	ctx context.Context,
	city string,
	country string,
	category string,
) ([]models.Event, error) {

	return nil, fmt.Errorf("api failure")

}

func (f FailingEventProvider) GetEventDetails(
	ctx context.Context,
	id string,
) (models.Event, error) {

	return models.Event{}, fmt.Errorf("event failure")

}

func TestGetCityEventsFailure(
	t *testing.T,
) {

	provider :=
		FailingEventProvider{}

	result, err :=
		GetCityEvents(
			context.Background(),
			provider,
			"Toronto",
			"CA",
		)

	if err != nil {

		t.Fatal(
			"service should keep partial failures",
		)

	}

	if len(result.Errors) == 0 {

		t.Fatal(
			"expected errors",
		)

	}

}
