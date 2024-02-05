package main

import (
	"fmt"
	"runtime"
	"time"
)

func worker(stop <-chan time.Time) {
	for {
		select {
		case <-stop:
			fmt.Println("Горутина не работает...")
			runtime.Goexit() //Остановка текущий горутины
		default:
			fmt.Println("Горутина работает...")
			time.Sleep(500 * time.Millisecond)
		}
		// Работа горутины
	}
}

func main() {
	stop := time.After(time.Second * 4) //Таймер на оставноку

	// Запускаем горутину
	go worker(stop)

	// Ждем некоторое время
	time.Sleep(6 * time.Second)

}
