package services

import (
	"sync"
	"time"

	"event-explorer/models"
)

type CacheItem struct {
	Events []models.Event

	ExpiresAt time.Time
}

var eventCache = make(map[string]CacheItem)

var cacheMutex sync.RWMutex

const cacheDuration = 5 * time.Minute

func GetCachedEvents(key string) ([]models.Event, bool) {

	cacheMutex.RLock()

	item, exists := eventCache[key]

	cacheMutex.RUnlock()

	if !exists {

		return nil, false

	}

	if time.Now().After(item.ExpiresAt) {

		cacheMutex.Lock()

		delete(eventCache, key)

		cacheMutex.Unlock()

		return nil, false

	}

	return item.Events, true

}

func SetCachedEvents(
	key string,
	events []models.Event,
) {

	cacheMutex.Lock()

	eventCache[key] = CacheItem{

		Events: events,

		ExpiresAt: time.Now().Add(cacheDuration),
	}

	cacheMutex.Unlock()

}
