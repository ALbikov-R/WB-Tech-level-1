package main

import (
	"fmt"
	"time"
)

func worker(stopChan chan struct{}) {
	for {
		select {
		case <-stopChan: //При закрытии канала в нём будет находиться значение по умолчанию

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
	stopChan := make(chan struct{})

	// Запускаем горутину
	go worker(stopChan)

	// Ждем некоторое время
	time.Sleep(2 * time.Second)

	// Останавливаем горутину
	close(stopChan)

	// Ждем завершения горутины
	time.Sleep(500 * time.Millisecond)
}
