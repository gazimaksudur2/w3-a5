package services

import (
	"net/url"
	"strings"
)

func IsValidTicketURL(
	rawURL string,
) bool {

	parsed, err :=
		url.Parse(rawURL)

	if err != nil {

		return false

	}

	if parsed.Scheme != "https" {

		return false

	}

	host :=
		strings.ToLower(
			parsed.Host,
		)

	allowedHosts := []string{

		"ticketmaster.com",

		"ticketmaster.ca",

		"www.ticketmaster.com",

		"www.ticketmaster.ca",
	}

	for _, allowed := range allowedHosts {

		if host == allowed {

			return true

		}

	}

	return false

}
