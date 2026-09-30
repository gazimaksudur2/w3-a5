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

type ticketmasterEventDetailResponse struct {
	ID string `json:"id"`

	Name string `json:"name"`

	URL string `json:"url"`

	Description string `json:"description"`

	Images []struct {
		URL string `json:"url"`
	} `json:"images"`

	Dates struct {
		Start struct {
			LocalDate string `json:"localDate"`
			LocalTime string `json:"localTime"`
		} `json:"start"`
	} `json:"dates"`

	Embedded struct {
		Venues []struct {
			Name string `json:"name"`

			City struct {
				Name string `json:"name"`
			} `json:"city"`

			Country struct {
				Name string `json:"name"`
			} `json:"country"`
		} `json:"venues"`
	} `json:"_embedded"`
}

func GetEventDetails(
	ctx context.Context,
	eventID string,
) (models.Event, error) {

	var event models.Event

	if eventID == "" {

		return event, fmt.Errorf("event id is required")

	}

	apiURL := fmt.Sprintf(
		"%s/events/%s.json",
		TicketmasterBaseURL,
		url.PathEscape(eventID),
	)

	req, err := http.NewRequest(
		"GET",
		apiURL,
		nil,
	)

	if err != nil {

		return event, err

	}

	query := req.URL.Query()

	query.Add(
		"apikey",
		config.TicketmasterAPIKey(),
	)

	req.URL.RawQuery = query.Encode()

	response, err := HTTPClient.Do(req.WithContext(ctx))

	if err != nil {

		return event, err

	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {

		return event,
			fmt.Errorf(
				"ticketmaster status: %d",
				response.StatusCode,
			)

	}

	var data ticketmasterEventDetailResponse

	err = json.NewDecoder(
		response.Body,
	).Decode(&data)

	if err != nil {

		return event, err

	}

	event = models.Event{

		ID: data.ID,

		Name: data.Name,

		TicketURL: data.URL,

		Description: data.Description,
	}

	if len(data.Images) > 0 {

		event.Image =
			data.Images[0].URL

	}

	event.Date =
		data.Dates.Start.LocalDate

	if len(data.Embedded.Venues) > 0 {

		venue :=
			data.Embedded.Venues[0]

		event.Venue =
			venue.Name

	}

	return event, nil

}
