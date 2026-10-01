package services

import "testing"

func TestIsValidTicketURL(t *testing.T) {

	tests := []struct {
		name string
		url  string
		want bool
	}{
		{
			name: "valid ticketmaster com",
			url:  "https://www.ticketmaster.com/event123",
			want: true,
		},
		{
			name: "valid ticketmaster ca",
			url:  "https://ticketmaster.ca/event123",
			want: true,
		},
		{
			name: "invalid http scheme",
			url:  "http://www.ticketmaster.com/event123",
			want: false,
		},
		{
			name: "invalid external domain",
			url:  "https://evil.com/fake",
			want: false,
		},
		{
			name: "invalid empty url",
			url:  "",
			want: false,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			got := IsValidTicketURL(tt.url)

			if got != tt.want {
				t.Errorf(
					"IsValidTicketURL(%q) = %v, want %v",
					tt.url,
					got,
					tt.want,
				)
			}

		})
	}
}