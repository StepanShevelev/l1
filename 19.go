package main

import "fmt"

// проход по половине массива рун (используем руны чтобы работать с unicode символами)
// и замена местами символов
func reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var s string
	fmt.Println("Введите строку:")
	fmt.Scanln(&s)
	fmt.Println(reverse(s))
}
