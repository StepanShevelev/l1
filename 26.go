package main

import (
	"fmt"
	"strings"
)

// используем мапу для проверки уникальных символов
func checkUnique(str string) bool {
	str = strings.ToLower(str)
	runes := []rune(str)
	unique := make(map[rune]struct{})

	for _, char := range runes {
		if _, ok := unique[char]; ok {
			return false
		}
		unique[char] = struct{}{}
	}

	return true
}

func main() {

	testStr := "abcd"
	fmt.Println(testStr, "-", checkUnique(testStr))

	testStr = "abCdefAaf"
	fmt.Println(testStr, "-", checkUnique(testStr))

	testStr = "aabcd"
	fmt.Println(testStr, "-", checkUnique(testStr))
}
