package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"math/rand"
	"sync"
	"time"
)

func main() {
	withWaitGroup()
	fmt.Printf("-------------\n\n")
	withErrorGroup(context.Background())
	fmt.Printf("-------------\n\n")
	withCancelContext(context.Background())
}

func withWaitGroup() {
	fmt.Printf("Start withWaitGroup\n\n")

	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Printf("Start goroutine %d\n", num)
			if i%2 == 0 {
				fmt.Printf("Error goroutine %d\n", num)
				return
			}
			sleep()
			fmt.Printf("Stop goroutine %d\n", num)
		}(i)
	}

	wg.Wait()
	fmt.Println("All done")
}

func withErrorGroup(context context.Context) {
	fmt.Printf("Start withErrorGroup\n\n")

	eg, _ := errgroup.WithContext(context)

	for i := 0; i < 10; i++ {
		eg.Go(func() error {
			num := i

			fmt.Printf("Start goroutine %d\n", num)
			sleep()
			if i%2 == 0 {
				fmt.Printf("Error goroutine %d\n", num)
				return errors.New("error goroutine")
			}
			fmt.Printf("Stop goroutine %d\n", num)

			return nil
		},
		)
	}

	if err := eg.Wait(); err != nil {
		fmt.Printf("All done with error %v\n", err)
	}

	fmt.Println("All done")
}

func withCancelContext(c context.Context) {

	fmt.Printf("Start withCancelContext\n\n")

	ctx, cancel := context.WithCancel(c)
	defer cancel()

	for i := 0; i < 10; i++ {
		go func(num int) {
			fmt.Printf("Start goroutine %d\n", num)
			sleep()
			if i%2 == 0 {
				fmt.Printf("Error goroutine %d\n", num)
				cancel()
				return
			}
			fmt.Printf("Stop goroutine %d\n", num)
		}(i)
	}

	<-ctx.Done()
	if err := ctx.Err(); err != nil && errors.Is(err, context.Canceled) {
		fmt.Printf("All done with cancel %v\n", err)
	}

	fmt.Println("All done")
}

func sleep() {
	minMillis, maxMillis := 1000, 2000
	multiplier := rand.Intn(maxMillis-minMillis) + maxMillis
	time.Sleep(time.Duration(multiplier) * time.Millisecond)
}
