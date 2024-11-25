package main

import (
	"fmt"
	"sync"
	"time"
)

// todo syncMap

var mp = map[string]int{"a": 1}
var mux = &sync.RWMutex{}

func main() {
	go rd()
	time.Sleep(time.Second)
	go wr()
	time.Sleep(time.Second)

}

func rd() {
	for {
		mux.RLock()
		fmt.Println(mp["a"])
		mux.RUnlock()
	}
}

func wr() {
	for {
		mux.Lock()
		mp["a"] = 2
		mux.Unlock()
	}
}
