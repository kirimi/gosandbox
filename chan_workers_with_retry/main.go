package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//  на входе есть набор адресов реплик kv-хранилища
//  нужно параллельно запросить их и вернуть успешный ответ первой завершенной, остальные отменить
//  реплика может вернуть ошибку в случае если что-то пошло не так, тогда нужно сделать ретраи.
//  если ретраи тоже закончились ошибкой, то вернуть ошибку ErrGetFailed
//  или реплика может вернуть ошибку, что ключа нет - тогда нужно отменить другие запросы и вернуть ошибку ErrNotFound

var ErrNotFound = errors.New("not found")
var ErrGetFailed = errors.New("get failed")

func main() {
	urls := []string{"url1", "url2", "url3", "url4"}
	key := "key"

	ctx := context.Background()
	ctxWithTO, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*2000))
	defer cancel()

	val, err := get(ctxWithTO, urls, key)
	if err != nil {
		fmt.Printf("Error. %v\n", err)
		return
	}

	fmt.Printf("Success. Value = %v\n", val)

	time.Sleep(time.Second * 1)
}

func get(ctx context.Context, urls []string, key string) (string, error) {
	contextWithCancel, cancel := context.WithCancel(ctx)
	defer cancel()

	resultCh := make(chan string)

	wg := &sync.WaitGroup{}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	wg.Add(len(urls))
	for _, url := range urls {
		go func() {
			defer wg.Done()

			var value string
			var err error

			getDataWithRetry(3, func() error {
				if contextWithCancel.Err() != nil {
					return nil
				}

				value, err = getData(contextWithCancel, url, key)
				if !errors.Is(err, ErrNotFound) {
					return err
				}

				return nil
			})

			if err != nil {
				return
			}

			select {
			case resultCh <- value:
			default:
			}
		}()
	}

	select {
	case res, ok := <-resultCh:
		if !ok {
			return res, ErrGetFailed
		}
		return res, nil
	case <-contextWithCancel.Done():
		return "", contextWithCancel.Err()
	}
}

func getDataWithRetry(
	maxRetry int,
	callback func() error,
) {
	for i := 0; i < maxRetry; i++ {
		err := callback()
		if err == nil {
			break
		}
	}
}

func getData(ctx context.Context, url, key string) (string, error) {

	requestDuration := rand.Intn(3000)

	fmt.Printf("%s running with %d delay\n", url, requestDuration)

	sleepWithContext(ctx, time.Millisecond*time.Duration(requestDuration))

	if ctx.Err() != nil {
		fmt.Printf("%s stop running cancel\n", url)
		return "", ctx.Err()
	}

	if requestDuration%2 == 0 {
		fmt.Printf("%s stop running error \n", url)
		return "", ErrGetFailed
	}

	if requestDuration%3 == 0 {
		fmt.Printf("%s stop running not found\n", url)
		return "", ErrNotFound
	}

	fmt.Printf("%s stop running success\n", url)
	return "value", nil
}

func sleepWithContext(ctx context.Context, delay time.Duration) {
	select {
	case <-ctx.Done():
	case <-time.After(delay):
	}
}
