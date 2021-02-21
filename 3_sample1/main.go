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

	select {
	case <-time.After(time.Second):
		fmt.Println("done")
	case <-ctx.Done():
		stop()
		fmt.Println("canceled")
	}
}

// 1秒経つと「done」と表示される
// 途中で control + C すると「canceled」と表示される
