package main

import (
	"fmt"
	"time"
)

// panic: send on closed channel
func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	time.Sleep(time.Millisecond * 500)

	close(ch)

	for i := range ch {
		fmt.Println(i)
	}

	time.Sleep(time.Millisecond * 100)
}
