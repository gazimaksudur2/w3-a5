package services

import (
	"context"

	"golang.org/x/sync/errgroup"

	"event-explorer/models"
)

func GetCityEvents(
	ctx context.Context,
	provider EventProvider,
	city string,
	country string,
) (models.EventResponse, error) {

	var response models.EventResponse

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {

		events, err :=
			provider.FetchEvents(
				ctx,
				city,
				country,
				"Music",
			)

		if err != nil {
			response.Errors =
				append(response.Errors, err)

			return nil
		}

		response.Music = events

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

		if err != nil {

			response.Errors =
				append(response.Errors, err)

			return nil
		}

		response.Sports = events

		return nil

	})

	err := g.Wait()

	return response, err

}
