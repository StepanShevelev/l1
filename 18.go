package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type counter struct {
	sync.Mutex
	Number int
}

func (c *counter) Increment() {
	c.Lock()
	defer c.Unlock()
	c.Number++
}

type atomicCounter struct {
	Number uint64
}

func (ac *atomicCounter) Increment() {
	atomic.AddUint64(&ac.Number, 1)
}

func main() {
	// используем структуру с мьютексом в которой хранится наш счетчик
	// при каждой операции с числом блокируемся мьютексом
	c := counter{
		Mutex:  sync.Mutex{},
		Number: 0,
	}

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}
	wg.Wait()

	fmt.Println(c.Number)

	// использование sync/atomic
	ac := atomicCounter{Number: 0}

	for i := 0; i < 1000; i++ {
		go func() {
			ac.Increment()
		}()
	}

	time.Sleep(time.Millisecond)
	fmt.Println(ac.Number)
}
