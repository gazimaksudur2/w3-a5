package services

import (
	"testing"
	"time"

	"event-explorer/models"
)


func TestCacheHit(t *testing.T) {


	cache :=
		NewEventCache()


	events :=
		[]models.Event{

			{
				ID:"123",
				Name:"Test Event",
			},

		}



	cache.Set(
		"toronto_CA_Music",
		events,
	)



	result,ok :=
		cache.Get(
			"toronto_CA_Music",
		)



	if !ok {

		t.Fatal("expected cache hit")

	}



	if result[0].ID != "123" {

		t.Fatal("wrong cached data")

	}

}



func TestCacheExpiry(t *testing.T) {


	cache :=
		&EventCache{

			items: map[string]CacheItem{

				"expired":
				{

					Events:nil,

					ExpiresAt:
						time.Now().Add(
							-1*time.Minute,
						),

				},

			},

		}



	_,ok :=
		cache.Get(
			"expired",
		)



	if ok {

		t.Fatal(
			"expected cache miss after expiry",
		)

	}

}