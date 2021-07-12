## LRU Cache proof of Concept

This repo contains two different implementations of a simple LRU cache for benchmarking purposes.


### Scenario

An in-memory cache that has a limit of N keys. When writing a new key while the cache is full, the least last used key 
should be deleted.

### internal/notscalablecache

This implementation tries to find the least last used key by a timestamp, looping through a `map[string]cache.Value` and
comparing its members. Needless to say that it's not scalable as the number of cached elements grows.


### internal/supercache

This implementation uses a map[string]*list.Element which points to a node of a doubly linked list.

When we read a key, we move it to the "front|top" of the queue, so the key at the "back|rear" should be the least last
used key, making it easier to find when writing to a full cache.

### Benchmark

#### Cache without doubly linked queue/list
```
pkg: github.com/trennepohl/lru-cache-poc/internal/notscalablecache
cpu: Intel(R) Core(TM) i5-7400 CPU @ 3.00GHz
BenchmarkNotScalableCacheWrites-4         434892              2758 ns/op
PASS
ok      github.com/trennepohl/lru-cache-poc/internal/notscalablecache   1.231s
```


### Cache with doubly linked queue/list
```
pkg: github.com/trennepohl/lru-cache-poc/internal/supercache
cpu: Intel(R) Core(TM) i5-7400 CPU @ 3.00GHz
BenchmarkScalableCacheWrites-4           3144170               393.0 ns/op
PASS
ok      github.com/trennepohl/lru-cache-poc/internal/supercache 1.627s
```