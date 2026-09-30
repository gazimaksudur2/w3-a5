package services

import "testing"

func TestValidTicketURL(t *testing.T) {

	valid :=
		IsValidTicketURL(
			"https://www.ticketmaster.com/event123",
		)

	if !valid {

		t.Fatal(
			"expected valid ticket URL",
		)

	}

}

func TestInvalidTicketURL(t *testing.T) {

	invalid :=
		IsValidTicketURL(
			"https://evil.com/fake",
		)

	if invalid {

		t.Fatal(
			"expected invalid ticket URL",
		)

	}

}
