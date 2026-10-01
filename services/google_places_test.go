package services

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAutocomplete(t *testing.T) {

	tests := []struct {
		name string
		input string
		status int
		response string
		wantLen int
		wantErr bool
	}{
		{
			name: "successful autocomplete",
			input: "Toronto",
			status: 200,
			response: `
			{
				"suggestions":[
					{
						"placePrediction":{
							"placeId":"abc123",
							"text":{
								"text":"Toronto"
							}
						}
					}
				]
			}`,
			wantLen: 1,
			wantErr: false,
		},
		{
			name: "short input",
			input: "a",
			wantLen: 0,
			wantErr: false,
		},
		{
			name: "api failure",
			input: "Toronto",
			status: 500,
			wantErr: true,
		},
	}


	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {


			if tt.status != 0 {

				server := httptest.NewServer(
					http.HandlerFunc(
						func(
							w http.ResponseWriter,
							r *http.Request,
						){

							w.WriteHeader(tt.status)

							if tt.status == 200 {
								w.Write([]byte(tt.response))
							}
						},
					),
				)

				defer server.Close()

				oldURL := GooglePlacesBaseURL
				GooglePlacesBaseURL = server.URL

				defer func(){
					GooglePlacesBaseURL = oldURL
				}()

			}


			result, err := Autocomplete(
				tt.input,
				"session",
			)


			if tt.wantErr && err == nil {
				t.Fatal("expected error")
			}


			if !tt.wantErr && err != nil {
				t.Fatal(err)
			}


			if len(result) != tt.wantLen {
				t.Errorf(
					"got %d results, want %d",
					len(result),
					tt.wantLen,
				)
			}

		})
	}
}