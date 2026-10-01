package services

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlaceDetailsSuccess(t *testing.T) {

	server :=
		httptest.NewServer(
			http.HandlerFunc(
				func(
					w http.ResponseWriter,
					r *http.Request,
				) {

					w.Header().Set(
						"Content-Type",
						"application/json",
					)

					w.Write([]byte(`
					{
						"addressComponents":[
							{
								"longText":"Toronto",
								"shortText":"Toronto",
								"types":[
									"locality"
								]
							},
							{
								"longText":"Canada",
								"shortText":"CA",
								"types":[
									"country"
								]
							}
						]
					}
					`))

				},
			),
		)

	defer server.Close()

	oldURL :=
		GooglePlacesBaseURL

	GooglePlacesBaseURL =
		server.URL

	defer func() {

		GooglePlacesBaseURL =
			oldURL

	}()

	location, err :=
		GetPlaceDetails(
			"place123",
			"session123",
		)

	if err != nil {

		t.Fatal(err)

	}

	if location.City != "Toronto" {

		t.Fatal(
			"city parsing failed",
		)

	}

	if location.CountryCode != "CA" {

		t.Fatal(
			"country parsing failed",
		)

	}

}

func TestGetPlaceDetailsAPIError(
	t *testing.T,
) {

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
		GooglePlacesBaseURL

	GooglePlacesBaseURL =
		server.URL

	defer func() {

		GooglePlacesBaseURL =
			oldURL

	}()

	_, err :=
		GetPlaceDetails(
			"invalid",
			"session",
		)

	if err == nil {

		t.Fatal(
			"expected error",
		)

	}

}
