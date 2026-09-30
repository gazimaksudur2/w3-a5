package services

import (
	"context"

	"golang.org/x/sync/errgroup"

	"event-explorer/models"
)

type eventFetchResult struct {
	category string

	events []models.Event

	err error
}

func GetCityEvents(
	ctx context.Context,
	provider EventProvider,
	city string,
	country string,
) (models.EventResponse, error) {

	var response models.EventResponse

	resultChannel :=
		make(chan eventFetchResult, 2)

	g, ctx :=
		errgroup.WithContext(ctx)

	g.Go(func() error {

		events, err :=
			provider.FetchEvents(
				ctx,
				city,
				country,
				"Music",
			)

		resultChannel <- eventFetchResult{

			category: "Music",

			events: events,

			err: err,
		}

		return nil

	})

	g.Go(func() error {

		events, err :=
			provider.FetchEvents(
				ctx,
				city,
				country,
				"Sports",
			)

		resultChannel <- eventFetchResult{

			category: "Sports",

			events: events,

			err: err,
		}

		return nil

	})

	go func() {

		g.Wait()

		close(resultChannel)

	}()

	for result := range resultChannel {

		if result.err != nil {

			response.Errors =
				append(
					response.Errors,
					result.err,
				)

			continue

		}

		switch result.category {

		case "Music":

			response.Music =
				result.events

		case "Sports":

			response.Sports =
				result.events

		}

	}

	return response, nil

}
