package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mutex *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cache: map[string]cacheEntry{},
		mutex: &sync.Mutex{},
	}
	go cache.reapLoop(interval)

	return cache
}

func (C Cache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(interval)
    for range ticker.C {
        C.mutex.Lock()
        fmt.Printf("reapLoop Lock")
        for k, v := range C.cache {
            if time.Now().Sub(v.createdAt) > interval {
                delete(C.cache, k)
            }
        }
        C.mutex.Unlock()
        fmt.Printf("reapLoop Unlock")
    }
}

func (C Cache) Add(key string, val []byte) {
    fmt.Printf("waiting for lock")
    C.mutex.Lock()
    fmt.Printf("add Lock")
	C.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
    C.mutex.Unlock()
    fmt.Printf("add unlock")
}

func (C Cache) Get(key string) ([]byte, bool) {
    C.mutex.Lock()
    fmt.Printf("get Lock")
    entry, ok := C.cache[key]
    if !ok {
        return []byte{}, ok
    }
    val := entry.val
    C.mutex.Unlock()
    fmt.Printf("get unlock")
    return val, ok
}
