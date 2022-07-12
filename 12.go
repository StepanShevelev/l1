package main

import "fmt"

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool) // итоговое множество слов

	for _, word := range words {
		if _, ok := set[word]; !ok { // если слова еще нет в множестве то добавляем его
			set[word] = true
		}
	}

	for k, _ := range set {
		fmt.Println(k)
	}
}
