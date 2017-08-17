package groupcache

import (
	"fmt"
	"sync"
	"time"
)

var (
	cache = make(map[string]*CacheGroup)
	// set zero for never expire
	globalLifeSpan = time.Nanosecond //time.Duration(time.Minute * 5)
	mutex          sync.RWMutex
)

// SetGlobalCacheExpire Set all cache groups lifespan
func SetGlobalCacheExpire(life time.Duration) {
	globalLifeSpan = life
}

// Cache if exists the cache group return the cache group ;
// otherwise create a new one
func Cache(group string) *CacheGroup {
	mutex.RLock()
	t, ok := cache[group]
	mutex.RUnlock()

	if !ok {
		mutex.Lock()
		t, ok = cache[group]

		if !ok {
			t = &CacheGroup{
				groupname:       group,
				values:          []*CacheItem{},
				createTime:      time.Now(),
				cleanupInterval: time.Duration(5 * time.Second),
				life:            globalLifeSpan,
			}
			fmt.Println("xxxxx")
			t.checkExpiration()
			cache[group] = t
		}
		mutex.Unlock()
	}

	return t
}
