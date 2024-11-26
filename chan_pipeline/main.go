package main

import "fmt"

func main() {
	in := make(chan int)
	out := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			in <- i
		}
		close(in)
	}()

	go func() {
		for val := range in {
			out <- val * val
		}
		close(out)
	}()

	for square := range out {
		fmt.Println(square)
	}
}
