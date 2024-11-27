package main

import (
	"fmt"
	"strings"
	"sync"
)

func worker(in string) string {
	return strings.ToUpper(in)
}

func doWithWorkerPool(jobs, result chan string, worker func(string) string, capacity int) {
	pool := make(chan string, capacity)

	go func() {
		for {
			job, ok := <-jobs
			if !ok {
				break
			}
			pool <- job
		}
		close(pool)
	}()

	wg := &sync.WaitGroup{}

	wg.Add(capacity)
	for i := 0; i < capacity; i++ {
		go func(workerId int) {
			defer wg.Done()
			for {
				val, ok := <-pool
				if !ok {
					break
				}
				result <- fmt.Sprintf("%d %s", workerId, worker(val))
			}
		}(i)
	}

	wg.Wait()
}

func main() {
	jobs := make(chan string)
	results := make(chan string)

	go func() {
		for i := 0; i < 100; i++ {
			jobs <- "kirill"
			jobs <- "Mironov"
			jobs <- "asdf"
			jobs <- "rrf"
			jobs <- "dsvsb"
			jobs <- "bsfbdfb"
		}

		close(jobs)
	}()

	go func() {
		for r := range results {
			fmt.Println(r)
		}
	}()

	doWithWorkerPool(jobs, results, worker, 5)
}
