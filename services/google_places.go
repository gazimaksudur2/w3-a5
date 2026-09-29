package services


import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"event-explorer/config"
	"event-explorer/models"
)



func Autocomplete(
	input string,
	sessionToken string,
) ([]models.LocationSuggestion,error){


	var result struct {

		Suggestions []struct{

			PlacePrediction struct{

				Text struct{

					Text string `json:"text"`

				} `json:"text"`


				PlaceID string `json:"placeId"`

			} `json:"placePrediction"`


		} `json:"suggestions"`

	}



	body:=map[string]interface{}{

		"input":input,

		"includedPrimaryTypes":[]string{
			"cities",
		},

		"sessionToken":sessionToken,

	}


	jsonBody,_:=json.Marshal(body)



	req,err:=http.NewRequest(
		"POST",
		"https://places.googleapis.com/v1/places:autocomplete",
		bytes.NewBuffer(jsonBody),
	)


	if err!=nil{
		return nil,err
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



	resp,err:=HTTPClient.Do(req)


	if err!=nil{
		return nil,err
	}


	defer resp.Body.Close()



	if resp.StatusCode!=200{

		return nil,
		fmt.Errorf(
			"google status %d",
			resp.StatusCode,
		)

	}



	json.NewDecoder(resp.Body).Decode(&result)



	suggestions:=[]models.LocationSuggestion{}



	for _,item:=range result.Suggestions{


		suggestions=append(
			suggestions,
			models.LocationSuggestion{

				Description:
				item.PlacePrediction.Text.Text,


				PlaceID:
				item.PlacePrediction.PlaceID,
			},
		)

	}


	return suggestions,nil

}