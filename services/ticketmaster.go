package services

import (
	"context"
	"encoding/json"

	"fmt"

	"net/http"

	"net/url"

	"event-explorer/config"

	"event-explorer/models"
)

var TicketmasterBaseURL = "https://app.ticketmaster.com/discovery/v2"

type ticketmasterResponse struct {
	Embedded struct {
		Events []struct {
			ID string `json:"id"`

			Name string `json:"name"`

			URL string `json:"url"`

			Images []struct {
				URL string `json:"url"`
			} `json:"images"`

			Dates struct {
				Start struct {
					LocalDate string `json:"localDate"`
				} `json:"start"`
			} `json:"dates"`

			Embedded struct {
				Venues []struct {
					Name string `json:"name"`
				} `json:"venues"`
			} `json:"_embedded"`
		} `json:"events"`
	} `json:"_embedded"`
}

func (t TicketmasterProvider) FetchEvents(
	ctx context.Context,
	city string,
	country string,
	category string,
) ([]models.Event, error) {

	return FetchEvents(
		ctx,
		city,
		country,
		category,
	)

}

func FetchEvents(
	ctx context.Context,
	city string,
	country string,
	category string,
) ([]models.Event, error) {

	key :=
		city + "_" + country + "_" + category

	cached, ok := EventCacheInstance.Get(key)

	if ok {

		return cached, nil

	}

	params :=
		url.Values{}

	params.Add(
		"apikey",
		config.TicketmasterAPIKey(),
	)

	params.Add(
		"city",
		city,
	)

	params.Add(
		"countryCode",
		country,
	)

	params.Add(
		"classificationName",
		category,
	)

	params.Add(
		"size",
		"6",
	)

	apiURL := TicketmasterBaseURL + "/events.json?" + params.Encode()

	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		apiURL,
		nil,
	)

	if err != nil {
		return nil, err
	}

	response, err := HTTPClient.Do(req)

	if err != nil {

		return nil, err

	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {

		return nil,
			fmt.Errorf(
				"ticketmaster status: %d",
				response.StatusCode,
			)

	}

	var data ticketmasterResponse

	err =
		json.NewDecoder(
			response.Body,
		).Decode(&data)

	if err != nil {

		return nil, err

	}

	events :=
		[]models.Event{}

	for _, item := range data.Embedded.Events {

		event :=
			models.Event{

				ID: item.ID,

				Name: item.Name,

				TicketURL: item.URL,
			}

		if len(item.Images) > 0 {

			event.Image =
				item.Images[0].URL

		}

		event.Date =
			item.Dates.Start.LocalDate

		if len(item.Embedded.Venues) > 0 {

			event.Venue =
				item.Embedded.Venues[0].Name

		}

		events =
			append(events, event)

	}

	EventCacheInstance.Set(key, city, country, category, events)

	return events, nil

}

type TicketmasterProvider struct{}

func (t TicketmasterProvider) GetEventDetails(
	ctx context.Context,
	eventID string,
) (models.Event, error) {

	return GetEventDetails(
		ctx,
		eventID,
	)
}
