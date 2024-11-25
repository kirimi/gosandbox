package main

import (
	"fmt"
	"time"
)

func main() {
	//runtime.GOMAXPROCS(1)

	//ctx := context.Background()
	//goroutine_escape.GproutineEscape(ctx)
	//
	//time.Sleep(time.Second * 3)

	//pool := NewPool(100, func(i int) string {
	//	return strconv.Itoa(i)
	//})
	//
	//go func(values <-chan string) {
	//	for {
	//		select {
	//		case val := <-values:
	//			fmt.Printf("Value: %s\n", val)
	//		}
	//	}
	//}(pool.out)
	//
	//for i := 0; i < 100; i++ {
	//	pool.doJob(i)
	//}
	//
	//time.Sleep(time.Second * 3)

	// Custom WG
	//wg := newWg()
	//
	//for i := 0; i < 1000; i++ {
	//	wg.add(1)
	//	go func(wg *MyWg) {
	//		defer wg.done()
	//		time.Sleep(time.Second / 3)
	//		fmt.Println(i)
	//	}(wg)
	//}
	//
	//wg.wait()

	//var wg sync.WaitGroup
	//
	//for x := 0; x < 10; x++ {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		fmt.Printf("Go %d\n", x)
	//
	//	}()
	//	//runtime.Gosched()
	//}
	//wg.Wait()
}

func worker(workerId int, ch <-chan int) {
	time.Sleep(time.Second)
	for {
		value, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("Worker %v, value %d\n", workerId, value)
	}
}

func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i := range slice {
		fmt.Printf("[%d] %p %s\n", i, &slice[i], slice[i])
	}
}
