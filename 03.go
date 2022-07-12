package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// для каждого числа вызываем горутину которая посчитает квадрат и сложит с atomic переменной
func main() {
	array := []uint64{2, 4, 6, 8, 10}
	var res uint64
	var wg sync.WaitGroup
	for _, num := range array {
		wg.Add(1)
		go func(number uint64) {
			defer wg.Done()
			atomic.AddUint64(&res, number*number)
		}(num)
	}

	wg.Wait()
	fmt.Println(res)
}
