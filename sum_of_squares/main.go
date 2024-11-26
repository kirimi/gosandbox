package main

import (
	"fmt"
)

//Реализуйте функцию SumOfSquares, получающую целое число c и возвращающую сумму всех квадратов между 1 и c.
// Вам потребуется использовать инструкции select, горутины и каналы.
//
//Например, ввод 5 приведет к возвращению 55, потому что $1² + 2² + 3² + 4² + 5² = 55$.

func SumOfSquares(c, quit chan int) {
	i := 1
	for {
		select {
		case c <- i * i:
			i++
		case <-quit:
			return
		}
	}
}
func main() {
	mychannel := make(chan int)
	quitchannel := make(chan int)
	sum := 0
	go func() {
		for i := 1; i <= 5; i++ {
			sum += <-mychannel
		}
		fmt.Println(sum)
		quitchannel <- 0
	}()
	SumOfSquares(mychannel, quitchannel)
}
