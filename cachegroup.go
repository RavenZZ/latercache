package groupcache

import (
	"fmt"
	"sync"
	"time"
)

// CacheGroup group for CacheItems
type CacheGroup struct {
	sync.RWMutex

	// the group of name
	GroupName string

	// cacheItems slice
	Values []*CacheItem

	// group's HP
	life time.Duration

	// group create time
	createTime time.Time

	// timer for check expiration
	cleanuptimer *time.Timer

	// check interval for expiration
	cleanupInterval time.Duration

	// Callback method triggerd when group expired
	allExpire func(group *CacheGroup)
}

// SetCacheGroupExpireCallback overrite global method for cachegroup expire
func (group *CacheGroup) SetCacheGroupExpireCallback(f func(*CacheGroup)) {
	group.Lock()
	defer group.Unlock()
	group.allExpire = f
}

// Push add Cache item to group
func (group *CacheGroup) Push(key string, data interface{}) *CacheItem {
	item := NewCacheItem(key, data)

	group.Lock()
	group.addInternal(item)
	return item
}

// Count get the Count of group
func (group *CacheGroup) Count() int {
	group.RLock()
	defer group.RUnlock()
	return len(group.Values)
}

// addInternal  internal function for
func (group *CacheGroup) addInternal(item *CacheItem) {

	group.Values = append(group.Values, item)
	group.Unlock()
}

// checkExpiration Expiration check loop
func (group *CacheGroup) checkExpiration() {
	fmt.Println("Check start")
	group.Lock()
	if group.cleanuptimer != nil {
		group.cleanuptimer.Stop()
	}

	now := time.Now()

	fmt.Println("Check 2")
	fmt.Println(group.createTime)
	fmt.Println(group.life)
	if now.Sub(group.createTime) >= group.life {
		fmt.Println("Check 3")
		group.groupExpire()
	} else {
		fmt.Println("Check 4")
		group.cleanuptimer = time.AfterFunc(group.cleanupInterval, func() {
			go group.checkExpiration()
		})
	}
	group.Unlock()
}

func (group *CacheGroup) groupExpire() {
	if group.allExpire != nil {
		fmt.Println("zzzzzzz")
		delete(cache, group.GroupName)
		group.allExpire(group)
		group.allExpire = nil
	} else if globalExpireCallback != nil {
		fmt.Println("zzzzzzz")
		delete(cache, group.GroupName)
		globalExpireCallback(group)
		group.allExpire = nil
	} else {
		fmt.Println("expire function has not set")
	}
}

// ExpireNow Do expiration now ~
func (group *CacheGroup) ExpireNow() {
	group.groupExpire()
}
