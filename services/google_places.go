package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"event-explorer/config"
	"event-explorer/models"
)

func Autocomplete(
	input string,
	sessionToken string,
) ([]models.LocationSuggestion, error) {

	if len(input) < 2 {
		return []models.LocationSuggestion{}, nil
	}

	var result struct {
		Suggestions []struct {
			PlacePrediction struct {
				Text struct {
					Text string `json:"text"`
				} `json:"text"`

				PlaceID string `json:"placeId"`
			} `json:"placePrediction"`
		} `json:"suggestions"`
	}

	body := map[string]interface{}{

		"input": input,

		"includedPrimaryTypes": []string{
			"(cities)",
		},

		"sessionToken": sessionToken,
	}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		"https://places.googleapis.com/v1/places:autocomplete",
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	req.Header.Set(
		"X-Goog-Api-Key",
		config.GoogleAPIKey(),
	)

	req.Header.Set(
		"X-Goog-FieldMask",
		"suggestions.placePrediction.text,suggestions.placePrediction.placeId",
	)

	resp, err := HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		body, _ := io.ReadAll(resp.Body)

		fmt.Println(
			"Google autocomplete error:",
		)

		fmt.Println(
			string(body),
		)

		return nil,
			fmt.Errorf(
				"google autocomplete status: %d",
				resp.StatusCode,
			)

	}

	err = json.NewDecoder(
		resp.Body,
	).Decode(&result)

	if err != nil {
		return nil, err
	}

	suggestions := []models.LocationSuggestion{}

	for _, item := range result.Suggestions {

		suggestions = append(
			suggestions,
			models.LocationSuggestion{

				Description: item.PlacePrediction.Text.Text,

				PlaceID: item.PlacePrediction.PlaceID,
			},
		)

	}

	return suggestions, nil

}

func GetPlaceDetails(
	placeID string,
	sessionToken string,
) (models.Location, error) {

	var location models.Location

	apiURL := "https://places.googleapis.com/v1/places/" + url.PathEscape(placeID)

	if sessionToken != "" {

		apiURL += "?sessionToken=" + url.QueryEscape(sessionToken)

	}

	req, err := http.NewRequest(
		"GET",
		apiURL,
		nil,
	)

	if err != nil {
		return location, err
	}

	req.Header.Set(
		"X-Goog-Api-Key",
		config.GoogleAPIKey(),
	)

	req.Header.Set(
		"X-Goog-FieldMask",
		"addressComponents",
	)

	resp, err := HTTPClient.Do(req)

	if err != nil {
		return location, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		body, _ := io.ReadAll(resp.Body)

		fmt.Println(
			"Google details error:",
		)

		fmt.Println(
			string(body),
		)

		return location,
			fmt.Errorf(
				"google details status: %d",
				resp.StatusCode,
			)

	}

	var result struct {
		AddressComponents []struct {
			LongText string `json:"longText"`

			ShortText string `json:"shortText"`

			Types []string `json:"types"`
		} `json:"addressComponents"`
	}

	err = json.NewDecoder(
		resp.Body,
	).Decode(&result)

	if err != nil {
		return location, err
	}

	for _, component := range result.AddressComponents {

		for _, t := range component.Types {

			switch t {

			case "locality":

				location.City =
					component.LongText

			case "country":

				location.CountryCode =
					component.ShortText

			}

		}

	}

	return location, nil

}
