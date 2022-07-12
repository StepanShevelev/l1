package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// используя канал для остановки
	exit := make(chan struct{})
	go func() {

		for {
			select {
			case <-exit:
				log.Println("goroutine with chanel stopped")
				close(exit)
				return
			default:

				time.Sleep(time.Second)
				fmt.Println(rand.Int())
			}
		}
	}()

	time.Sleep(2 * time.Second)
	exit <- struct{}{}

	// используя таймер
	timer := time.NewTimer(time.Second)

	go func() {
		for {
			select {
			case <-timer.C:
				log.Println("goroutine with timer stopped")
				return
			default:
				time.Sleep(time.Second)
				fmt.Println(rand.Int())
			}
		}
	}()

	time.Sleep(time.Second * 2)

	// Используя контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				log.Println("goroutine with context stopped")
				return
			default:
				time.Sleep(time.Second)
				fmt.Println(rand.Int())
			}
		}
	}(ctx)

	time.Sleep(3 * time.Second)
	cancel()

	// Используя контекст с таймаутом
	ctxTime, _ := context.WithTimeout(context.Background(), 2*time.Second)
	wg.Add(1)
	go func(ctxTime context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctxTime.Done():
				log.Println("goroutine with timeout stopped")
				return
			default:
				time.Sleep(time.Second)
				fmt.Println(rand.Int())
			}
		}

	}(ctxTime)

	// Используя контекст с дедлайном
	ctxDead, _ := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	wg.Add(1)
	go func(ctxDead context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctxDead.Done():
				log.Println("goroutine with deadline stopped")
				return
			default:
				time.Sleep(time.Second)
				fmt.Println(rand.Int())
			}
		}

	}(ctxDead)

	wg.Wait()
}
