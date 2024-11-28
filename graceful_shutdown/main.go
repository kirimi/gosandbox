package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main1() {
	ctx, cancel := context.WithCancel(context.Background())
	osSigCh := make(chan os.Signal)
	signal.Notify(osSigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-osSigCh
		cancel()
	}()

	<-ctx.Done()
}

// с версии 1.16 можно так
func main2() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		time.Sleep(time.Second)
		cancel()
	}()

	<-ctx.Done()
}
