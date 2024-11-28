package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	resCh := make(chan int)
	currValueCh := make(chan int)

	firstValue, lastValue := 1, 100
	for i := firstValue; i <= lastValue; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case currValue := <-currValueCh:
					nextValue := currValue + 1
					if i == nextValue {
						resCh <- i
						if i < lastValue {
							currValueCh <- i
						}
						return
					}

					currValueCh <- currValue
				}
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(resCh)
		close(currValueCh)
	}()

	// run
	currValueCh <- firstValue - 1

	for val := range resCh {
		fmt.Printf("%d ", val)
	}
}
