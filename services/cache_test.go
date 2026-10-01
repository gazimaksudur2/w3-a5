package services

import (
	"testing"
	"time"

	"event-explorer/models"
)


func TestEventCache(t *testing.T) {

	tests := []struct {
		name string
		setup func(*EventCache)
		key string
		wantExists bool
	}{
		{
			name: "cache hit with valid entry",
			setup: func(c *EventCache) {

				c.Set(
					"toronto_CA_music",
					"Toronto",
					"CA",
					"Music",
					[]models.Event{
						{
							ID: "123",
							Name: "Concert",
						},
					},
				)

			},
			key: "toronto_CA_music",
			wantExists: true,
		},
		{
			name: "cache miss for missing key",
			setup: func(c *EventCache) {
				// no cache entry
			},
			key: "missing_key",
			wantExists: false,
		},
		{
			name: "expired cache entry",
			setup: func(c *EventCache) {

				c.items["expired"] = CacheItem{
					Events: []models.Event{
						{
							ID: "1",
						},
					},
					ExpiresAt: time.Now().Add(-time.Hour),
				}

			},
			key: "expired",
			wantExists: false,
		},
	}


	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {


			cache := NewEventCache()


			tt.setup(cache)


			_, exists := cache.Get(tt.key)


			if exists != tt.wantExists {

				t.Errorf(
					"cache.Get() = %v, want %v",
					exists,
					tt.wantExists,
				)

			}

		})
	}
}





func TestCacheInvalidation(t *testing.T) {

	tests := []struct {
		name              string
		city              string
		country           string
		category          string
		expectedRemaining int
	}{
		{
			name:              "invalidate by city",
			city:              "Toronto",
			expectedRemaining: 1,
		},
		{
			name:              "invalidate by country",
			country:           "CA",
			expectedRemaining: 1,
		},
		{
			name:              "invalidate by category",
			category:          "Music",
			expectedRemaining: 1,
		},
		{
			name:              "invalidate all cache",
			expectedRemaining: 0,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			cache := NewEventCache()

			cache.items = map[string]CacheItem{

				"Toronto_CA_Music": {
					City:     "Toronto",
					Country:  "CA",
					Category: "Music",
				},

				"Toronto_CA_Sports": {
					City:     "Toronto",
					Country:  "CA",
					Category: "Sports",
				},

				"London_GB_Music": {
					City:     "London",
					Country:  "GB",
					Category: "Music",
				},
			}


			cache.Invalidate(
				tt.city,
				tt.country,
				tt.category,
			)


			if len(cache.items) != tt.expectedRemaining {

				t.Errorf(
					"remaining cache items=%d, want=%d",
					len(cache.items),
					tt.expectedRemaining,
				)

			}

		})
	}
}