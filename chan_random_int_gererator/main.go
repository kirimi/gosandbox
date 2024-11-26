package main

import (
	"fmt"
	"math/rand"
)

func randNumsGenerator(count int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
	}()

	return out
}

func main() {
	for num := range randNumsGenerator(3) {
		fmt.Println(num)
	}
}
