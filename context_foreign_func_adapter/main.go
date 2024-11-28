package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*2000)

	var res int
	var err error

	errFunc := foreignFuncAdapterContext(ctx, func() {
		res, err = task()
	})

	if errFunc != nil {
		log.Fatalf("return error %v", errFunc)
	}

	if err != nil {
		log.Fatal("task return error")
	}
	fmt.Printf("Result %d\n", res)
}

func task() (int, error) {
	time.Sleep(time.Second)

	return 0, nil
}

func foreignFuncAdapterContext(ctx context.Context, foreignFunc func()) error {
	doneCh := make(chan struct{})

	go func() {
		foreignFunc()
		close(doneCh)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-doneCh:
		return nil
	}
}
