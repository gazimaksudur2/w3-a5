package services

import (
	"testing"
	"time"

	"event-explorer/models"
)

func TestEventCache(t *testing.T) {

	tests := []struct {
		name string
		key string
		item CacheItem
		wantExists bool
	}{
		{
			name: "cache hit",
			key: "toronto_CA_music",
			item: CacheItem{
				Events: []models.Event{
					{
						ID: "123",
						Name: "Concert",
					},
				},
				ExpiresAt: time.Now().Add(time.Hour),
			},
			wantExists: true,
		},
		{
			name: "expired cache",
			key: "expired",
			item: CacheItem{
				Events: nil,
				ExpiresAt: time.Now().Add(-time.Hour),
			},
			wantExists: false,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			cache := &EventCache{
				items: map[string]CacheItem{
					tt.key: tt.item,
				},
			}

			_, ok := cache.Get(tt.key)

			if ok != tt.wantExists {
				t.Errorf(
					"cache.Get() = %v, want %v",
					ok,
					tt.wantExists,
				)
			}

		})
	}
}