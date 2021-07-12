package supercache

import (
	"container/list"
	"fmt"
	"github.com/trennepohl/lru-cache-poc/pkg/cache"
)

type simpleCache struct {
	maxKeyCount int
	cacheIndex  map[string]*list.Element
	cacheValues *list.List
}

func (s *simpleCache) Write(key string, value string){
	element := cache.Value{
		Value: value,
		Key: key,
	}

	if len(s.cacheIndex) >= s.maxKeyCount {
		toDelete := s.cacheValues.Back()
		s.cacheValues.Remove(toDelete)
		delete(s.cacheIndex, toDelete.Value.(cache.Value).Key)
	}

	s.cacheIndex[key] = s.cacheValues.PushFront(element)
}

func (s *simpleCache) Delete(key string) error {
	delete(s.cacheIndex, key)
	fmt.Println(s.cacheIndex)
	return nil
}

func (s *simpleCache) Clear() {
	s.cacheIndex = make(map[string]*list.Element, s.maxKeyCount)
}

func (s *simpleCache) Count() int {
	return len(s.cacheIndex)
}

func (s *simpleCache) Read(key string) (value string, err error) {
	if !s.keyExists(key){
		return value, fmt.Errorf("the key=%s specified doesn't exist", key)
	}

	element :=  s.cacheIndex[key]
	s.cacheValues.MoveToFront(element)
	return element.Value.(cache.Value).Value, err
}

func New(config cache.Config) (superCache cache.LRUCache, err error) {
	if config.MaxKeyCount == 0 {
		return superCache, fmt.Errorf("InvalidCapacity")
	}

	return &simpleCache{
		maxKeyCount : config.MaxKeyCount,
		cacheIndex:   make(map[string]*list.Element, config.MaxKeyCount),
		cacheValues: list.New(),
	}, err
}

func (s *simpleCache) keyExists(key string) bool {
	if _, exist := s.cacheIndex[key]; !exist {
		return false
	}
	return true
}