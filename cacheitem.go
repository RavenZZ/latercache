package groupcache

import (
	"sync"
	"time"
)

// CacheItem cache item for store data
type CacheItem struct {
	sync.RWMutex

	// items Key for ensure that we cannot add data twice
	Key string

	// data
	Value interface{}

	// item create time
	createTime time.Time
	// v1 does not support item expire
	// life time.Duration
	// createTime time.Time
}

// NewCacheItem returns a new CacheItem
func NewCacheItem(key string, data interface{}) *CacheItem {
	t := time.Now()
	return &CacheItem{
		Key:        key,
		Value:      data,
		createTime: t,
	}
}
