package main

import (
	"fmt"
	"sync"
	"time"
)

var arr = []int{2, 4, 6, 8, 10}

// для каждого числа запускаем горутину в которой выводим квадрат числа
// sleep для ожидания всех горутин
func firstSolution() {
	for _, num := range arr {
		go func(number int) {
			fmt.Println(number * number)
		}(num)
	}

	time.Sleep(time.Second * 1)
}

// решение с помошью небуферезированного канала
// для каждого числа из массива запускаем горутину которая пишет в канал значение квадрта числа
// после чего читаем из канала len(arr) раз
func secondSolution() {
	res := make(chan int)
	for i, _ := range arr {
		go func(i int) {
			res <- arr[i] * arr[i]
		}(i)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(<-res)
	}
	close(res)
}

//решение с помошью waitGroup и буферезированного канала
// также для каждого числа запускаем горутину в которой пишем в буфферизриованный канал квадрат числа
// ждем выполнения всех горутин с помощью wg.Wait(), закрываем канал и в цикле читаем все его значения
func thirdSolution() {
	res := make(chan int, len(arr))
	var wg sync.WaitGroup
	for _, num := range arr {
		wg.Add(1)
		go func(number int) {
			defer wg.Done()
			res <- number * number
			fmt.Println(<-res)
		}(num)
	}

	wg.Wait()
	close(res)

}

func main() {
	firstSolution()
	secondSolution()
	thirdSolution()
}
