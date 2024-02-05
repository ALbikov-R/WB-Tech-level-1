package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): //При закрытии контекста в нём будет содержаться пустая структура
			fmt.Println("Горутина завершена.")
			return
		default:
			// Работа горутины
			fmt.Println("Горутина работает...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// Запускаем горутину
	go worker(ctx)

	// Ждем некоторое время
	time.Sleep(2 * time.Second)

	// Останавливаем горутину
	cancel()

	// Ждем завершения горутины
	time.Sleep(500 * time.Millisecond)
}
