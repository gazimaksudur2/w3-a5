package services

import (
	"log"
	"strings"
	"sync"
	"time"

	"event-explorer/models"
)


type CacheItem struct {
	Events []models.Event

	City string
	Country string
	Category string

	ExpiresAt time.Time
}



type EventCache struct {
	items map[string]CacheItem

	mutex sync.RWMutex
}



const cacheDuration = 5 * time.Minute



func NewEventCache() *EventCache {

	return &EventCache{
		items: make(map[string]CacheItem),
	}
}




func (c *EventCache) Get(key string) ([]models.Event, bool) {


	c.mutex.RLock()

	item, exists := c.items[key]

	c.mutex.RUnlock()



	if !exists {

		log.Printf(
			"[CACHE MISS] key=%s",
			key,
		)

		return nil, false
	}




	// Lazy expiration check
	if time.Now().After(item.ExpiresAt) {


		log.Printf(
			"[CACHE EXPIRED] key=%s",
			key,
		)


		c.Delete(key)


		return nil, false
	}




	log.Printf(
		"[CACHE HIT] key=%s city=%s country=%s category=%s expires=%s",
		key,
		item.City,
		item.Country,
		item.Category,
		item.ExpiresAt.Format(time.RFC3339),
	)



	return item.Events, true
}





func (c *EventCache) Set(
	key string,
	city string,
	country string,
	category string,
	events []models.Event,
) {


	c.mutex.Lock()

	defer c.mutex.Unlock()



	expiry :=
		time.Now().Add(cacheDuration)



	c.items[key] = CacheItem{

		Events: events,

		City: city,

		Country: country,

		Category: category,

		ExpiresAt: expiry,
	}



	log.Printf(
		"[CACHE SET] key=%s city=%s country=%s category=%s events=%d expires=%s",
		key,
		city,
		country,
		category,
		len(events),
		expiry.Format(time.RFC3339),
	)

}




func (c *EventCache) Delete(key string) {


	c.mutex.Lock()

	defer c.mutex.Unlock()



	delete(
		c.items,
		key,
	)



	log.Printf(
		"[CACHE DELETE] key=%s",
		key,
	)

}





func (c *EventCache) Invalidate(
	city string,
	country string,
	category string,
) {


	c.mutex.Lock()

	defer c.mutex.Unlock()

	removed := 0

	if city == "" &&
		country == "" &&
		category == "" {


		removed = len(c.items)


		c.items = make(map[string]CacheItem)



		log.Printf(
			"[CACHE INVALIDATE ALL] removed=%d",
			removed,
		)


		return
	}




	log.Printf(
		"[CACHE INVALIDATE REQUEST] city=%s country=%s category=%s",
		city,
		country,
		category,
	)





	for key, item := range c.items {


		match := true



		if city != "" &&
			!strings.EqualFold(item.City, city) {

			match = false
		}




		if country != "" &&
			!strings.EqualFold(item.Country, country) {

			match = false
		}




		if category != "" &&
			!strings.EqualFold(item.Category, category) {

			match = false
		}





		if match {


			delete(
				c.items,
				key,
			)


			removed++



			log.Printf(
				"[CACHE INVALIDATED] key=%s",
				key,
			)

		}

	}





	log.Printf(
		"[CACHE INVALIDATE COMPLETE] removed=%d remaining=%d",
		removed,
		len(c.items),
	)

}