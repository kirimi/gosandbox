package waitgroup_custom_impl

import (
	"fmt"
	"sync"
)

type MyWg struct {
	doneChan chan interface{}
	count    int
	mu       *sync.Mutex
}

func newWg() *MyWg {
	return &MyWg{doneChan: make(chan interface{}), mu: &sync.Mutex{}}
}

func (w *MyWg) add(num int) {
	w.mu.Lock()
	w.count += num
	fmt.Printf("Add, Counter = %d\n", w.count)
	w.mu.Unlock()
}

func (w *MyWg) done() {
	w.mu.Lock()
	w.count--
	fmt.Printf("done, Counter = %d\n", w.count)
	w.mu.Unlock()
	if w.count <= 0 {
		w.doneChan <- true
	}
}

func (w *MyWg) wait() {
	<-w.doneChan
	fmt.Printf("Wait done, Counter = %d\n", w.count)
}
