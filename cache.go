package groupcache

import (
	"sync"
	"time"
)

var (
	cache = make(map[string]*CacheGroup)
	// set zero for never expire
	globalLifeSpan       = time.Nanosecond //time.Duration(time.Minute * 5)
	mutex                sync.RWMutex
	globalExpireCallback func(group *CacheGroup)
)

// SetGlobalCacheExpire Set all cache groups lifespan
func SetGlobalCacheExpire(life time.Duration) {
	globalLifeSpan = life
}

// SetGlobalCacheExpireCallback Set Global Expire Function
func SetGlobalCacheExpireCallback(f func(*CacheGroup)) {
	globalExpireCallback = f
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
				GroupName:       group,
				Values:          []*CacheItem{},
				createTime:      time.Now(),
				cleanupInterval: time.Duration(1 * time.Second),
				life:            globalLifeSpan,
				allExpire:       globalExpireCallback,
			}

			t.checkExpiration()
			cache[group] = t
		}
		mutex.Unlock()
	}

	return t
}
