package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var Num int

	fmt.Println("Set timer in seconds:")
	fmt.Scanln(&Num)

	ch := make(chan int)                                  // канал для записи и чтения
	timer := time.After(time.Duration(Num) * time.Second) // таймер  для остановки
	var wg sync.WaitGroup                                 // для ожидания завершения горутин
	wg.Add(1)
	// Запись в канал
	go func() {
		defer wg.Done()
		for {
			select {
			case <-timer:
				log.Printf("Write to channel stopped after %d seconds ", Num)
				close(ch)
				return
			default:
				randNum := rand.Intn(100)
				ch <- randNum
				time.Sleep(time.Second)
			}
		}
	}()

	wg.Add(1)
	// Чтение из канала
	go func() {
		defer wg.Done()
		for range ch {
			if val, ok := <-ch; ok {
				fmt.Printf("received %d\n", val)
			} else {
				fmt.Println("Channel closed!")
			}

		}
	}()

	wg.Wait()
}
