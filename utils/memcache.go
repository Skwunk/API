package utils

import (
	"log"
	"fmt"
	memcache "github.com/bradfitz/gomemcache/memcache"
)

var Memcache *memcache.Client

func NewMemcacheClient(c *Config) {
	Memcache = memcache.New(c.Memcache)
}

func test() {
	value, err := Memcache.Get("test1")
	if err != nil {
		log.Print(err)
		Memcache.Set(&memcache.Item{Key: "test1", Value: []byte("test value")})
	} else {
		sz := string(value.Value)
		fmt.Printf(sz)
	}
	value, err = Memcache.Get("test1")
	if err != nil {
		log.Print(err)
		Memcache.Set(&memcache.Item{Key: "test1", Value: []byte("test value")})
	} else {
		sz := string(value.Value)
		fmt.Printf(sz)
	}
}