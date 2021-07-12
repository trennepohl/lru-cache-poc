package main

import (
	"fmt"
	"github.com/trennepohl/lru-cache-poc/internal/supercache"
	"github.com/trennepohl/lru-cache-poc/pkg/cache"
)

//Take a look at the tests :)
func main(){
	for i:= 0; i< 200000000; i++ {	c, _ := supercache.New(cache.Config{MaxKeyCount: 100})

		kv := fmt.Sprintf("key%d", i)
		c.Write(kv,kv)
	}
}
