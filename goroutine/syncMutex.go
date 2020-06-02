package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	v map[string]int
	sMutex sync.Mutex
}

func (c *Counter) Increment(key string) {
	c.sMutex.Lock()
	defer c.sMutex.Unlock()
	c.v[key]++
}

func (c *Counter) Value(key string) int{
	c.sMutex.Lock()
	defer c.sMutex.Unlock()
	return c.v[key]
}

func main() {
//	c := make(map[string]int)
	c := Counter{v: make(map[string]int)}
	// ２つのgoroutineから2つのmapを呼ぶとエラーになる
	// fatal error: concurrent map writes
	// sync.Mutexを使うと解決できる
	go func() {
		for i := 0; i < 10; i++ {
			//c["key"] += 1
			c.Increment("key")
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
//			c["key"] += 1
			c.Increment("key")
		}
	}()
	time.Sleep(1 * time.Second)
//	fmt.Println(c, c["key"])
	fmt.Println(c," : " , c.Value("key"))
}
