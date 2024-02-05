package main

import (
	"fmt"
	"time"
)

func worker(stopChan chan struct{}) {
	// stopChan будет закрыт автоматически при завершении функции main
	// из-за defer close(stopChan)
	defer close(stopChan)

	for {
		// Работа горутины
		select {
		case <-stopChan:
			fmt.Println("Горутина завершена.")
			return
		default:
			fmt.Println("Горутина работает...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	stopChan := make(chan struct{})

	// Запускаем горутину
	go worker(stopChan)

	// Ждем некоторое время
	time.Sleep(2 * time.Second)

	// Останавливаем горутину
}
