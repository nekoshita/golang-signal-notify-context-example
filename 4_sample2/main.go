package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func blocking(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("worker started")
	<-ctx.Done()
	fmt.Println("worker canceled")
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	var wg sync.WaitGroup
	wg.Add(3)
	go blocking(ctx, &wg)
	go blocking(ctx, &wg)
	go blocking(ctx, &wg)

	wg.Wait()
}

// goroutineが複数起動される
// control + c　で一括キャンセル
