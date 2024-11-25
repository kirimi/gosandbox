package goroutine_escape

import (
	"context"
	"fmt"
	"time"
)

func GproutineEscape(ctx context.Context) {
	ch := make(chan struct{}, 1)
	go func() {
		defer close(ch)
		time.Sleep(time.Second * 1)
		ch <- struct{}{}

		fmt.Println("goroutine is live")
	}()

	return

	//select {
	//case <-ctx.Done():
	//	fmt.Println("exit from GproutineEscape: context done")
	//	return
	//case <-ch:
	//	fmt.Println("exit from GproutineEscape: write to chan")
	//	return
	//}

}
