package services

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTicketmasterProviderFetchEvents(
	t *testing.T,
) {

	provider :=
		TicketmasterProvider{}

	oldURL :=
		TicketmasterBaseURL

	TicketmasterBaseURL =
		createMockTicketmasterServer(
			t,
		)

	defer func() {

		TicketmasterBaseURL =
			oldURL

	}()

	events, err :=
		provider.FetchEvents(
			context.Background(),
			"London",
			"GB",
			"Music",
		)

	if err != nil {

		t.Fatal(err)

	}

	if len(events) == 0 {

		t.Fatal(
			"expected events",
		)

	}

}

func createMockTicketmasterServer(
	t *testing.T,
) string {

	server :=
		httptest.NewServer(
			http.HandlerFunc(
				func(
					w http.ResponseWriter,
					r *http.Request,
				) {

					w.Header().
						Set(
							"Content-Type",
							"application/json",
						)

					w.Write([]byte(`
					{
						"_embedded":{
							"events":[
								{
									"id":"100",
									"name":"Mock Concert"
								}
							]
						}
					}
					`))

				},
			),
		)

	t.Cleanup(
		server.Close,
	)

	return server.URL

}

func TestTicketmasterProviderGetEventDetails(
	t *testing.T,
) {

	provider :=
		TicketmasterProvider{}

	oldURL :=
		TicketmasterBaseURL

	TicketmasterBaseURL =
		createMockTicketmasterDetailServer(
			t,
		)

	defer func() {

		TicketmasterBaseURL =
			oldURL

	}()

	event, err :=
		provider.GetEventDetails(
			context.Background(),
			"100",
		)

	if err != nil {

		t.Fatal(err)

	}

	if event.ID != "100" {

		t.Fatal(
			"wrong event",
		)

	}

}

func createMockTicketmasterDetailServer(
	t *testing.T,
) string {

	server :=
		httptest.NewServer(
			http.HandlerFunc(
				func(
					w http.ResponseWriter,
					r *http.Request,
				) {

					w.Header().
						Set(
							"Content-Type",
							"application/json",
						)

					w.Write([]byte(`
					{
						"id":"100",
						"name":"Mock Event"
					}
					`))

				},
			),
		)

	t.Cleanup(
		server.Close,
	)

	return server.URL

}
