package main

import (
	"fmt"
	"sync"
)

// Конкурентно порахувати суму кожного слайсу int, та роздрукувати результат.
// Потрібно використовувати WaitGroup.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// Порядок друку не важливий.
// “slice 1: 10”
// “slice 2: 16”
func main() {

	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	var wg sync.WaitGroup
	for i := 0; i < len(n); i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			sum(i, n[i])
		}(i)
	}

	wg.Wait()
}

func sum(i int, n []int) {
	sum := 0
	for i := 0; i < len(n); i++ {

		sum = sum + n[i]
	}
	fmt.Println("slice", i, ":", sum)
}
