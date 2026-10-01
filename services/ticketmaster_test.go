package services

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchEventsSuccess(t *testing.T) {

	server :=
		httptest.NewServer(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {

					w.Header().Set(
						"Content-Type",
						"application/json",
					)

					w.Write([]byte(`
					{
						"_embedded":{
							"events":[
								{
									"id":"1",
									"name":"Concert",
									"url":"https://www.ticketmaster.com/test",
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

	events, err :=
		FetchEvents(
			context.Background(),
			"Toronto",
			"CA",
			"Music",
		)

	if err != nil {

		t.Fatal(err)

	}

	if len(events) != 1 {

		t.Fatal(
			"expected one event",
		)

	}

	if events[0].Name != "Concert" {

		t.Fatal(
			"wrong event",
		)

	}

}

func TestFetchEventsAPIError(t *testing.T) {

	server :=
		httptest.NewServer(
			http.HandlerFunc(
				func(
					w http.ResponseWriter,
					r *http.Request,
				) {

					w.WriteHeader(500)

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

	_, err :=
		FetchEvents(
			context.Background(),
			"FailureCity",
			"XX",
			"Music",
		)

	if err == nil {

		t.Fatal(
			"expected error",
		)

	}

}
