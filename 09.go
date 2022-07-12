package main

import "fmt"

// конвейер чисел
// Принимает массив и канал, в который будет записывать числа из массива.
func in(arr []int, ch1 chan int) {
	defer close(ch1)
	for _, val := range arr {
		ch1 <- val
	}
}

// Принимает данные из канала ch1 и записывает в канал ch2.
func out(ch2 chan int, ch1 chan int) {
	defer close(ch2)
	for v := range ch1 {
		ch2 <- v * 2
	}
}

func main() {
	// Создаем каналы ch1 и ch2 и массив, из которого будем брать числа.
	ch1 := make(chan int)
	ch2 := make(chan int)
	array := []int{1, 2, 3, 4, 5}

	// Запускаем горутины
	go in(array, ch1)
	go out(ch2, ch1)

	//печатаем содержимое второго канала.
	for v := range ch2 {
		fmt.Println(v)
	}
}
