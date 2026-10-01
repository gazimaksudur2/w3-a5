package services

import (
	"context"
	"errors"
	"testing"

	"event-explorer/models"
)

type TableProvider struct {
	failCategory string
}

func (p TableProvider) FetchEvents(
	ctx context.Context,
	city string,
	country string,
	category string,
) ([]models.Event, error) {

	if category == p.failCategory {
		return nil, errors.New("api failed")
	}

	return []models.Event{
		{
			ID:   "1",
			Name: category + " Event",
		},
	}, nil
}

func (p TableProvider) GetEventDetails(
	ctx context.Context,
	id string,
) (models.Event, error) {

	return models.Event{}, nil
}

func TestGetCityEvents(t *testing.T) {

	tests := []struct {
		name       string
		provider   EventProvider
		wantErrors int
		wantMusic  bool
		wantSports bool
	}{
		{
			name:       "successful fetching",
			provider:   TableProvider{},
			wantErrors: 0,
			wantMusic:  true,
			wantSports: true,
		},
		{
			name: "music api failure",
			provider: TableProvider{
				failCategory: "Music",
			},
			wantErrors: 1,
			wantSports: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			result, err := GetCityEvents(
				context.Background(),
				tt.provider,
				"Toronto",
				"CA",
			)

			if err != nil {
				t.Fatal(err)
			}

			if len(result.Errors) != tt.wantErrors {
				t.Errorf(
					"errors=%d want=%d",
					len(result.Errors),
					tt.wantErrors,
				)
			}

			if tt.wantMusic && len(result.Music) == 0 {
				t.Error("music events missing")
			}

			if tt.wantSports && len(result.Sports) == 0 {
				t.Error("sports events missing")
			}

		})
	}

}
