package main

import "sync"

var cache = struct {
	sync.Mutex // 内嵌
	mapping    map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock() // 直接调用 Lock()/Unlock()
	value := cache.mapping[key]
	cache.Unlock()
	return value
}
