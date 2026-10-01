package services

import (
	"testing"
	"time"

	"event-explorer/models"
)

func TestRemoveExpiredCacheItems(
	t *testing.T,
) {

	cache :=
		&EventCache{

			items: make(map[string]CacheItem),
		}

	cache.items["expired"] =
		CacheItem{

			Events: []models.Event{

				{
					ID: "1",
				},
			},

			ExpiresAt: time.Now().Add(
				-1 * time.Hour,
			),
		}

	cache.removeExpired(
		time.Now(),
	)

	_, exists :=
		cache.items["expired"]

	if exists {

		t.Fatal(
			"expired cache item was not removed",
		)

	}

}
