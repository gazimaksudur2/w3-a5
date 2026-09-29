package services

import (
	"encoding/json"

	"fmt"

	"net/http"

	"net/url"

	"event-explorer/config"

	"event-explorer/models"
)

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

func FetchEvents(
	city string,
	country string,
	category string,
) ([]models.Event, error) {

	key :=
		city + "_" + country + "_" + category

	cached, ok :=
		GetCachedEvents(key)

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

	apiURL :=
		"https://app.ticketmaster.com/discovery/v2/events.json?"+params.Encode()

	response, err :=
		http.Get(apiURL)

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

	SetCachedEvents(
		key,
		events,
	)

	return events, nil

}
