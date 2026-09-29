package services

import (
	"event-explorer/models"
)

type EventResult struct {
	Category string

	Events []models.Event

	Err error
}

func GetCityEvents(
	city string,
	country string,
) (models.EventResponse, error) {

	resultChannel :=
		make(chan EventResult, 2)

	go func() {

		events, err :=
			FetchEvents(
				city,
				country,
				"Music",
			)

		resultChannel <- EventResult{

			Category: "Music",

			Events: events,

			Err: err,
		}

	}()

	go func() {

		events, err :=
			FetchEvents(
				city,
				country,
				"Sports",
			)

		resultChannel <- EventResult{

			Category: "Sports",

			Events: events,

			Err: err,
		}

	}()

	response :=
		models.EventResponse{}

	for i := 0; i < 2; i++ {

		result :=
			<-resultChannel

		if result.Err != nil {

			response.Errors =
				append(response.Errors, result.Err)

			continue

		}

		switch result.Category {

		case "Music":

			response.Music = result.Events

		case "Sports":

			response.Sports = result.Events

		}

	}

	return response, nil

}
