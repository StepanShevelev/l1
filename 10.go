package main

import "fmt"

func main() {
	t := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := make(map[int][]float64)

	for _, v := range t {
		group := int(v/10) * 10        // определяем к какой группе будет принадлежать число
		m[group] = append(m[group], v) // добавляем число в группу
	}

	for k, v := range m {
		fmt.Println("Group: ", k, " Temperature: ", v) // вывод всех групп
	}

}
