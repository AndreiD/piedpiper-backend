package caching

import (
	"errors"
	"time"
)

const (
	DEFAULT = time.Duration(0)
	FOREVER = time.Duration(-1)
)

var (
	PageCachePrefix = "gincontrib.page.cache"
	ErrCacheMiss    = errors.New("cache: key not found")
	ErrNotStored    = errors.New("cache: not stored")
	ErrNotSupport   = errors.New("cache: not support")
)

// CacheStore is the interface of a caching backend
type CacheStore interface {
	// Get retrieves an item from the caching. Returns the item or nil, and a bool indicating
	// whether the key was found.
	Get(key string, value interface{}) error

	// Set sets an item to the caching, replacing any existing item.
	Set(key string, value interface{}, expire time.Duration) error

	// Add adds an item to the caching only if an item doesn't already exist for the given
	// key, or if the existing item has expired. Returns an error otherwise.
	Add(key string, value interface{}, expire time.Duration) error

	// Replace sets a new value for the caching key only if it already exists. Returns an
	// error if it does not.
	Replace(key string, data interface{}, expire time.Duration) error

	// Delete removes an item from the caching. Does nothing if the key is not in the caching.
	Delete(key string) error

	// Increment increments a real number, and returns error if the value is not real
	Increment(key string, data uint64) (uint64, error)

	// Decrement decrements a real number, and returns error if the value is not real
	Decrement(key string, data uint64) (uint64, error)

	// Flush seletes all items from the caching.
	Flush() error
}
