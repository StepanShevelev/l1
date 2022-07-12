package main

import (
	"fmt"
	"math/rand"
)

func main() {
	set1 := make([]int, 10)
	set2 := make([]int, 10)

	var cross []int

	// Заполняем наборы
	for i, _ := range set1 {
		set1[i] = rand.Intn(15)
	}
	for i, _ := range set2 {
		set2[i] = rand.Intn(15)
	}
	// Ищем пересечение
	for _, v1 := range set1 {
		for _, v2 := range set2 {
			if v1 == v2 {
				cross = append(cross, v1)
				break
			}
		}
	}

	fmt.Println("Set1 =", set1)
	fmt.Println("Set2 =", set2)
	fmt.Println("Cross =", cross)
}
