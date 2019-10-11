package hn

import (
	"log"
	"time"

	"github.com/logrusorgru/aurora"
)

const defaultExpirationDuration = 15 * time.Minute

// Cache Hacker News stories
type Cache struct {
	value              map[int]*cachedItem
	ExpirationDuration time.Duration
}

// NewCache instanciates a new cache
func NewCache(expirationDuration time.Duration) *Cache {
	return &Cache{
		value:              map[int]*cachedItem{},
		ExpirationDuration: expirationDuration,
	}
}

// NewDefaultCache instanciates a new cache with default expiration duration
func NewDefaultCache() *Cache {
	return NewCache(defaultExpirationDuration)
}

// Add the story into the cache
func (c *Cache) Add(i Item) {
	it := c.value[i.ID]
	if it == nil {
		log.Println(aurora.Sprintf(aurora.BrightBlue("Story %d is not cached. Putting to cache..."), i.ID))
		c.value[i.ID] = &cachedItem{Item: i, CreationDate: time.Now()}
		time.AfterFunc(c.ExpirationDuration, func() {
			c.Invalidate(i.ID)
		})
	} else {
		log.Println(aurora.Sprintf(aurora.BrightBlack("Story %d has already been cached. Updating its date..."), i.ID))
		it.CreationDate = time.Now()
	}
}

// Invalidate cache from a given id
func (c *Cache) Invalidate(id int) {
	log.Println(aurora.Sprintf(aurora.BrightRed("Invalidating story %d from cache"), id))
	c.value[id] = nil
}

// Get the item from cache
func (c *Cache) Get(id int) *Item {
	cachedItem := c.value[id]
	if cachedItem == nil {
		log.Println(aurora.Sprintf(aurora.BrightRed("Story %d not found in cache"), id))
		return nil
	}
	log.Println(aurora.Sprintf(aurora.BrightBlack("Fetching story %d from cache"), id))
	return &cachedItem.Item
}

type cachedItem struct {
	Item
	CreationDate time.Time
}
