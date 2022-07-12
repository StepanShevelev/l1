package main

import "fmt"

//взять первое значение как опорную точку
//перебирать остальные элементы
//разделить элементы на те, которые больше и те, которые меньше опорной точки
//рекурсивно отсортировать меньшие и большие элементы

func quickSort(input []int) []int {
	l := len(input)
	if l < 2 {
		return input
	}
	less := make([]int, 0)
	bigger := make([]int, 0)
	pivot := input[0]
	for _, v := range input[1:] {
		if v > pivot {
			bigger = append(bigger, v)
		} else {
			less = append(less, v)
		}
	}
	input = append(quickSort(less), pivot)
	input = append(input, quickSort(bigger)...)
	return input
}

func main() {
	array := []int{5, 6, 7, 2, 1, 0, 4, 3, 6}
	fmt.Println(quickSort(array))
}
