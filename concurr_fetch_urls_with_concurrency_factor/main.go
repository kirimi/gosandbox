package main

import (
	"fmt"
	"sync"
	"time"
)

// Делает параллельные запросы к urls, c concurrencyFactor

func main() {
	urls := []string{
		"url1",
		"url2",
		"url3",
		"url4",
	}

	concurrencyFactor := 2

	resChan := fetchUrls(urls, concurrencyFactor)

	for res := range resChan {
		fmt.Println(res)
	}
}

func fetchUrls(urls []string, cf int) <-chan string {
	resCh := make(chan string)

	// сюда поместятся только cf воркеров
	workersCh := make(chan struct{}, cf)

	taskCh := make(chan string)
	go func() {
		for _, url := range urls {
			taskCh <- url
		}
		close(taskCh)
	}()

	go func() {
		wg := &sync.WaitGroup{}

		defer func() {
			wg.Wait()
			close(resCh)
		}()

		for {
			select {
			case url, ok := <-taskCh:
				if !ok {
					return
				}

				// занимаем воркер,
				// будем ждать тут, если канал уже заполнен
				workersCh <- struct{}{}

				wg.Add(1)
				go func() {
					defer wg.Done()
					res := fetch(url)
					resCh <- res

					// освобождаем воркер
					<-workersCh
				}()
			}
		}

	}()

	return resCh
}

func fetch(url string) string {
	time.Sleep(time.Second)
	return "result from " + url
}
