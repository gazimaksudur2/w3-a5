package services

import (
	"log"
	"sync"
	"time"

	"event-explorer/models"
)

type CacheItem struct {
	Events    []models.Event
	ExpiresAt time.Time
}

type EventCache struct {
	items map[string]CacheItem
	mutex sync.RWMutex
}

const cacheDuration = 5 * time.Minute

func NewEventCache() *EventCache {

	cache := &EventCache{
		items: make(map[string]CacheItem),
	}

	go cache.cleanup()

	return cache
}

func (c *EventCache) Get(key string) ([]models.Event, bool) {

	c.mutex.RLock()

	item, exists := c.items[key]

	c.mutex.RUnlock()

	if !exists {
		return nil, false
	}

	if time.Now().After(item.ExpiresAt) {

		c.Delete(key)

		return nil, false
	}

	log.Printf(
		"cache hit: %s",
		key,
	)

	return item.Events, true
}

func (c *EventCache) Set(
	key string,
	events []models.Event,
) {

	c.mutex.Lock()

	c.items[key] = CacheItem{

		Events: events,

		ExpiresAt: time.Now().Add(cacheDuration),
	}

	c.mutex.Unlock()

}

func (c *EventCache) Delete(key string) {

	c.mutex.Lock()

	delete(
		c.items,
		key,
	)

	c.mutex.Unlock()

}

func (c *EventCache) cleanup() {

	ticker :=
		time.NewTicker(
			time.Minute,
		)

	defer ticker.Stop()

	for range ticker.C {

		c.removeExpired(
			time.Now(),
		)

	}

}

func (c *EventCache) removeExpired(
	now time.Time,
) {

	c.mutex.Lock()

	defer c.mutex.Unlock()

	for key, item := range c.items {

		if now.After(
			item.ExpiresAt,
		) {

			delete(
				c.items,
				key,
			)

			log.Printf(
				"cache expired: %s",
				key,
			)

		}

	}

}
