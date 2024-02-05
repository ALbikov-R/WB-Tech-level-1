package main

import (
	"fmt"
	"time"
)

func worker(stopChan chan struct{}) {
	for {
		// Работа горутины
		select {
		case <-stopChan: //Горутина получила пустую структуру и может отработать case
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

	// Останавливаем горутину путем передачи элемента
	stopChan <- struct{}{}
}
