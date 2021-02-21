package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		select {
		case <-c:
			fmt.Fprintln(os.Stderr, "signal received")
			cancel()
		case <-ctx.Done():
		}
	}()

	doSomethingAwesome(ctx)
}

func doSomethingAwesome(ctx context.Context) {
	for i := 1; i <= 1000; i++ {
		select {
		case <-ctx.Done():
			time.Sleep(1 * time.Second)
			fmt.Print("canceled\n")
			return
		case <-time.After(1 * time.Second):
			fmt.Print("processing...\n")
		}
	}
}
