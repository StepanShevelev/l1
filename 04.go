package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//если произошло чтение <-ctx.Done(), выходим из воркера
func worker(ctx context.Context, id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for range ch {
		select {
		case <-ctx.Done():
			return
		default:
			if val, ok := <-ch; ok {
				fmt.Printf("worker %d received %d\n", id, val)
			} else {
				fmt.Println("Channel closed!")
			}

		}
	}
}

func main() {
	var wg sync.WaitGroup
	var workersNum int
	ch := make(chan int)                         // главный канал
	signalChannel := make(chan os.Signal, 1)     // канал для обработки SIGINT
	signal.Notify(signalChannel, syscall.SIGINT) // запись в signalChannel если пришел SIGINT
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Set workers num:")
	fmt.Scan(&workersNum)
	log.Println("Program started")

	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go worker(ctx, i, ch, &wg)
	}

	// если пользоватеьл вводит Ctrl+C
	// происходит отмена контекста и закрытие каналов
	for {
		select {
		case <-signalChannel:
			cancel()
			close(ch)
			close(signalChannel)
			wg.Wait()
			log.Println("Program interrupted")
			return
		default:
			ch <- rand.Intn(100)

			time.Sleep(time.Second)
		}

	}

}
