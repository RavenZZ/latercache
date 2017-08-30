package groupcache

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type a struct {
	name string
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

var ()

func TestCacheWithExpire(t *testing.T) {
	fmt.Println("start")
	SetGlobalCacheExpireCallback(func(group *CacheGroup) {
		fmt.Println("expired=====items count:", len(group.Values))
	})
	SetGlobalCacheExpire(time.Second * 6)
	group1 := Cache("user1:module1")
	fmt.Println(group1)
	// group1.SetCacheGroupExpireCallback(func(group *CacheGroup) {
	// 	fmt.Println("expired=====items count:", len(group.values))
	// })
	someid := RandStringRunes(10)
	v1 := "some value"
	group1.Push(someid, v1)
	v2 := "some value2"
	group1.Push(someid, v2)
	v3 := "some value3"
	group1.Push(someid, v3)
	v4 := "some value4"
	group1.Push(someid, v4)
	time.Sleep(10 * time.Second)

	fmt.Println("==expired=====items count:", len(Cache("user1:module1").Values))

}

func TestExpireNow(t *testing.T) {
	SetGlobalCacheExpire(time.Second * 8)
	group1 := Cache("user1:module1")
	//fmt.Println(group1)
	group1.SetCacheGroupExpireCallback(func(group *CacheGroup) {
		fmt.Println("expired=====items count:", len(group.Values))
	})
	someid := RandStringRunes(10)
	v1 := "some value"
	group1.Push(someid, v1)
	v2 := "some value2"
	group1.Push(someid, v2)
	v3 := "some value3"
	group1.Push(someid, v3)
	v4 := "some value4"
	group1.Push(someid, v4)
	time.Sleep(2 * time.Second)
	group1.ExpireNow()
	fmt.Println("==expired=====items count:", len(Cache("user1:module1").Values))
}

func TestPointer(t *testing.T) {
	SetGlobalCacheExpire(time.Second * 4)
	group1 := Cache("user1:module1")
	//fmt.Println(group1)
	group1.SetCacheGroupExpireCallback(func(group *CacheGroup) {
		fmt.Println(group.Count())
		for _, val := range group.All() {
			v := val
			fmt.Println(v.value.(*a).name)
		}
		fmt.Println("expired=====items count:", group.Count())
	})
	someid := RandStringRunes(10)
	v1 := a{name: "some value1"}
	group1.Push(someid, &v1)
	v2 := a{name: "some value2"}
	group1.Push(someid, &v2)
	v3 := a{name: "some value3"}
	group1.Push(someid, &v3)
	v4 := a{name: "some value4"}
	group1.Push(someid, &v4)
	fmt.Println(group1.Count())
	time.Sleep(6 * time.Second)
	//	group1.ExpireNow()
	fmt.Println("==expired=====items count:", Cache("user1:module1").Count())
	// for _, val := range group1.All() {
	// 	v := val
	// 	fmt.Println(v.value.(*a).name)
	// }
	// fmt.Println("expired=====items count:", group1.Count())
}
