package main

import (
	"fmt"
	"sync"
	"time"
)

var m = map[string]int{"a": 1}
var mu = &sync.RWMutex{}

func main() {
	go read()
	time.Sleep(time.Second)
	go write()
	time.Sleep(time.Second)

}

func read() {
	for {
		mu.RLock()
		fmt.Println(m["a"])
		mu.RUnlock()
	}
}

func write() {
	for {
		mu.Lock()
		m["a"] = 2
		mu.Unlock()
	}
}
