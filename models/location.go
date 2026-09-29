package models


type LocationSuggestion struct {

	Description string `json:"description"`

	PlaceID string `json:"placeId"`

}


type Location struct {

	City string `json:"city"`

	CountryCode string `json:"countryCode"`

}