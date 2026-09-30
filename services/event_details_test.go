package services

import (
	"context"
	"testing"
)


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