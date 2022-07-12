package main

import (
	"fmt"
)

func main() {
	result, count := SearchInIntSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 6)
	fmt.Println(result, count)
}

func SearchInIntSlice(haystack []int, needle int) (result bool, iterationsCount int) {

	lowKey := 0                  // первый индекс
	highKey := len(haystack) - 1 // последний индекс
	if (haystack[lowKey] > needle) || (haystack[highKey] < needle) {
		return // нужное значение не в диапазоне данных
	}
	for lowKey <= highKey {
		// уменьшаем список рекурсивно
		iterationsCount++
		mid := (lowKey + highKey) / 2 // середина
		if haystack[mid] == needle {
			result = true // мы нашли значение
			return
		}
		if haystack[mid] < needle {
			// если поиск больше середины - мы берем только блок с большими значениями увеличивая lowKey
			lowKey = mid + 1
			continue
		}
		if haystack[mid] > needle {
			// если поиск меньше середины - мы берем блок с меньшими значениями уменьшая highKey
			highKey = mid - 1
		}
	}
	return
}
