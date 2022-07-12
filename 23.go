package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	i := 2
	fmt.Println(a)
	// Выполнить сдвиг a[i+1:] влево на один индекс.
	copy(a[i:], a[i+1:])
	// Удалить последний элемент (записать нулевое значение).
	a[len(a)-1] = 0
	// Усечь срез.
	a = a[:len(a)-1]

	fmt.Println(a)

}
