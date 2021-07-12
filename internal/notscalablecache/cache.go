package notscalablecache

import (
	"fmt"
	"github.com/trennepohl/lru-cache-poc/pkg/cache"
	"time"
)

type lowScalabilityCache struct {
	maxKeyCount int
	inMemoryCache map[string]cache.Value
}

func (s *lowScalabilityCache) Write(key string, value string) {
	v := cache.Value{
		Value: value,
		LastUsed: time.Now(),
	}

	if len(s.inMemoryCache) >= s.maxKeyCount {
		toDelete := s.findLeastRecentlyUsed()
		delete(s.inMemoryCache, toDelete)
	}

	s.inMemoryCache[key] = v
}

func (s *lowScalabilityCache) findLeastRecentlyUsed() string{

	oldestElement := cache.Value{
		LastUsed: time.Now().Add(1000 * time.Hour),
	}

	var oldestKey string
	for key, element := range s.inMemoryCache {
		if element.LastUsed.Before(oldestElement.LastUsed) {
			oldestElement = element
			oldestKey = key
		}
	}
	return oldestKey
}

func (s *lowScalabilityCache) Delete(key string) error {
	delete(s.inMemoryCache, key)
	fmt.Println(s.inMemoryCache)
	return nil
}

func (s *lowScalabilityCache) Clear() {
	s.inMemoryCache = make(map[string]cache.Value, s.maxKeyCount)
}

func (s *lowScalabilityCache) Count() int {
	return len(s.inMemoryCache)
}

func (s *lowScalabilityCache) Read(key string) (value string, err error) {
	if !s.keyExists(key){
		return value, fmt.Errorf("KeyNotFound")
	}
	element :=  s.inMemoryCache[key]
	element.LastUsed  = time.Now()
	s.inMemoryCache[key] = element
	return element.Value, err
}

func New(config cache.Config) (superCache cache.LRUCache, err error) {

	if config.MaxKeyCount == 0 {
		return superCache, fmt.Errorf("InvalidCapacity")
	}

	return &lowScalabilityCache{
		maxKeyCount : config.MaxKeyCount,
		inMemoryCache: make(map[string]cache.Value, config.MaxKeyCount),
	}, err
}

func (s *lowScalabilityCache) keyExists(key string) bool {
	if _, exist := s.inMemoryCache[key]; !exist {
		return false
	}
	return true
}