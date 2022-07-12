package main

import (
	"fmt"
	"sync"
)

// Мьютекс лочится, когда какая-то горутина уже начала работу с мапой, только после завершения
// этой работы, получаем анлок для других горутин
// Без мьютекса из-за состояния гонки получим ошибку
func write(m map[int]int, i int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	mu.Lock()
	m[i] = i
	mu.Unlock()
}

// тоже самое с использованием sync.Map
func writeS(m *sync.Map, i int, wg *sync.WaitGroup) {
	defer wg.Done()

	m.Store(i, i)

}

func main() {

	var sm sync.Map

	m := make(map[int]int)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go write(m, i, &wg, &mu)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go writeS(&sm, i, &wg)
	}

	wg.Wait()
	fmt.Println(m)
	fmt.Println(sm.Load(0))
	fmt.Println(sm.Load(1))
	fmt.Println(sm.Load(2))
	fmt.Println(sm.Load(3))
	fmt.Println(sm.Load(4))

}
