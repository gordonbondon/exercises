package lru_cache1

import "fmt"

type LRUCache struct {
	capacity int
	used     int
	items    map[string]*CacheItem
	first    *CacheItem
	last     *CacheItem
}

type CacheItem struct {
	Key   string
	Value interface{}

	previous *CacheItem
	next     *CacheItem
}

func NewLRUCache(capacity int) (*LRUCache, error) {
	if capacity == 0 {
		return nil, fmt.Errorf("Capacity can not be 0")
	}

	cache := &LRUCache{
		capacity: capacity,
		used:     0,
	}

	cache.items = make(map[string]*CacheItem)

	return cache, nil
}

func (c *LRUCache) Set(key string, value interface{}) {
	if v, ok := c.items[key]; ok {
		v.Value = value
		c.bump(v)

		return
	}

	item := &CacheItem{
		Key:   key,
		Value: value,
	}

	c.items[key] = item

	if c.first == nil {
		c.first = item
	} else {
		item.next = c.first
		c.first.previous = item
		c.first = item
	}

	if c.last == nil {
		c.last = item
	}

	c.used++

	if c.used > c.capacity {
		temp := c.last
		c.last = temp.previous
		c.last.next = nil
		delete(c.items, temp.Key)
	}
}

func (c *LRUCache) Get(key string) (interface{}, bool) {
	if v, ok := c.items[key]; ok {
		c.bump(v)

		return v.Value, true
	}

	return nil, false
}

func (c *LRUCache) bump(item *CacheItem) {
	if c.first.Key == item.Key {
		return
	}

	if c.last.Key == item.Key {
		item.next = c.first
		c.first.previous = item
		c.first = item

		c.last = item.previous
		c.last.next = nil
		item.previous = nil

		return
	}

	prev := item.previous
	next := item.next

	item.next = c.first
	c.first.previous = item
	c.first = item

	prev.next = next
	next.previous = prev
}
