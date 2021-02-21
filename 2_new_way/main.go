package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

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
