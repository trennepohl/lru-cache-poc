package cache

type LRUCache interface {
	Write(key string, value string)
	Delete(key string) error
	Clear()
	Count() int
	Read(key string) (string, error)
}