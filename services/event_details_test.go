package services

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEventDetailsSuccess(t *testing.T) {

	server :=
		httptest.NewServer(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {

					w.Header().
						Set(
							"Content-Type",
							"application/json",
						)

					w.Write([]byte(`
					{
						"id":"123",
						"name":"Test Event",
						"url":"https://www.ticketmaster.com/test",
						"description":"Demo event",

						"images":[
							{
								"url":"image.jpg"
							}
						],

						"dates":{
							"start":{
								"localDate":"2026-01-01"
							}
						},

						"_embedded":{
							"venues":[
								{
									"name":"Arena"
								}
							]
						}
					}
					`))

				},
			),
		)

	defer server.Close()

	oldURL :=
		TicketmasterBaseURL

	TicketmasterBaseURL =
		server.URL

	defer func() {
		TicketmasterBaseURL = oldURL
	}()

	event, err :=
		GetEventDetails(
			context.Background(),
			"123",
		)

	if err != nil {

		t.Fatal(err)

	}

	if event.Name != "Test Event" {

		t.Fatal("wrong event")

	}

	if event.Venue != "Arena" {

		t.Fatal("venue missing")

	}

}

func TestInvalidEventID(t *testing.T) {

	_, err :=
		GetEventDetails(
			context.Background(),
			"",
		)

	if err == nil {

		t.Fatal(
			"expected error for empty event id",
		)

	}

}
