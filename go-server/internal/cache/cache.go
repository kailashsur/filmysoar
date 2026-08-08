package cache

import (
	"sync"
	"time"
)

type entry struct {
	value   any
	expires time.Time
}

var store = struct {
	sync.RWMutex
	items map[string]entry
}{items: make(map[string]entry)}

const maxEntries = 128

func Get(key string) (any, bool) {
	store.RLock()
	item, ok := store.items[key]
	store.RUnlock()
	if !ok || time.Now().After(item.expires) {
		if ok {
			Delete(key)
		}
		return nil, false
	}
	return item.value, true
}

func Set(key string, value any, ttl time.Duration) {
	store.Lock()
	defer store.Unlock()
	if len(store.items) >= maxEntries {
		for oldKey := range store.items {
			delete(store.items, oldKey)
			break
		}
	}
	store.items[key] = entry{value: value, expires: time.Now().Add(ttl)}
}

func Delete(key string) { store.Lock(); delete(store.items, key); store.Unlock() }
func DeletePrefix(prefix string) {
	store.Lock()
	for key := range store.items {
		if len(key) >= len(prefix) && key[:len(prefix)] == prefix {
			delete(store.items, key)
		}
	}
	store.Unlock()
}
