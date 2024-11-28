package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func merge2(ch ...<-chan int) <-chan int {
	out := make(chan int)

	wg := &sync.WaitGroup{}

	wg.Add(len(ch))
	for _, c := range ch {
		go func(c <-chan int) {
			for {
				select {
				case val, ok := <-c:
					if ok {
						out <- val
						continue
					}
					wg.Done()

					return
				}
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out

}

func source2(sourceFunc func(int) int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- sourceFunc(i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()

	return ch
}

func main() {

	in1 := source2(func(_ int) int {
		return rand.Int()
	})

	in2 := source2(func(i int) int {
		return i
	})

	out := merge2(in1, in2)

	for val := range out {
		fmt.Printf("%d, ", val)
	}

}
