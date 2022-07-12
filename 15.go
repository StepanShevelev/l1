package main

import "fmt"

var justString string

// Cтрока в Go - слайс байт.
// Так как некоторые символы занимают 2 или 3 байта(кириллица или китайские символы) нужно привести v к слайсу рун
// Пример сделан на генерации 1024 символов буквы Б, которая занимает два байта

// Генерация строки
func createHugeString(size int) (s string) {
	for i := 0; i < size; i++ {
		s += "Б"
	}
	return
}

// Вывод первых 100 символов
func someFunc() {
	v := createHugeString(1 << 10)
	a := []rune(v)
	// Здесь мы сразу видим, что количество символов различается в байтах и рунах. Поэтому, чтобы у нас было правильно отображение,
	// нужно переводить строку в руны.
	fmt.Println(len(v))
	fmt.Println(len(a))

	justString = string(a[:100])

	fmt.Println(justString)

}

func main() {

	someFunc()

}
