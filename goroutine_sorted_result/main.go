package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	resCh := make(chan int)
	lastValueCh := make(chan int)

	first, last := 1, 100
	for i := first; i <= last; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case lastV := <-lastValueCh:
					if i == lastV+1 {
						resCh <- i
						if i < last {
							lastValueCh <- i
						}
						return
					}

					lastValueCh <- lastV
				}
			}
		}(i)
	}

	go func() {
		lastValueCh <- first - 1
		wg.Wait()
		close(resCh)
		close(lastValueCh)
	}()

	for val := range resCh {
		fmt.Printf("%d ", val)
	}
}
