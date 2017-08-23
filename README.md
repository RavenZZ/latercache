groupcache
==========

golang caching library 
support expire cache items in **one expiration** 
so you can handle items in one method

## Installation

```
    go get -u github.com/ravenzz/groupcache
```

## Example

```go
package main

import (
    cache "github.com/ravenzz/groupcache"
)

type myData struct{
    Text string
}

func main() {
    // Set Global Expire, every group will follow this option
    cache.SetGlobalCacheExpire(time.Second * 8)
    group1 := cache.Cache("groupname1")

    group1.SetCacheGroupExpireCallback(func(group *CacheGroup) {
		fmt.Println("expired=====items count:", len(group.values))
	})

    group1.Push("some id1", myData{"some value"})
    group1.Push("some id2", myData{"some value2"})
    group1.Push("some id3", myData{"some value3"})
    time.Sleep(13 * time.Second)

    fmt.Println("==expired=====items count:", len(Cache("user1:module1").values))

    fmt.Println("=====================")
    
    group1.Push("some id1", myData{"some value"})
    group1.Push("some id2", myData{"some value2"})
    group1.Push("some id3", myData{"some value3"})
	group1.ExpireNow()
    fmt.Println("==expired=====items count:", len(Cache("user1:module1").values))
}



```

