package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func merge(ch ...<-chan int) <-chan int {
	out := make(chan int)

	wg := &sync.WaitGroup{}

	for _, c := range ch {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out

}

func source(sourceFunc func(int) int) <-chan int {
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

	in1 := source(func(_ int) int {
		return rand.Int()
	})

	in2 := source(func(i int) int {
		return i
	})

	out := merge(in1, in2)

	for val := range out {
		fmt.Printf("%d, ", val)
	}

}
