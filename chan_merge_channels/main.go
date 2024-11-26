package main

import (
	"fmt"
	"sync"
	"time"
)

func joinChannels(in ...chan int) <-chan int {
	out := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(len(in))
	for _, ch := range in {
		go func() {
			defer wg.Done()

			for val := range ch {
				out <- val
			}

			//for {
			//	val, ok := <-ch
			//	if !ok {
			//		break
			//	}
			//	out <- val
			//}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		for num := range joinChannels(a, b, c) {
			fmt.Println(num)
		}
	}()

	go func() {
		for _, num := range []int{1, 2, 3} {
			a <- num
		}
		close(a)
	}()

	go func() {
		for _, num := range []int{20, 10, 30} {
			b <- num
		}
		close(b)
	}()

	go func() {
		for _, num := range []int{300, 200, 100} {
			c <- num
		}
		close(c)
	}()

	time.Sleep(time.Second)
}
