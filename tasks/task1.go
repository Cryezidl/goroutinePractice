package tasks

import (
	"fmt"
	"sync"
	"time"
)

/*
Вывод чисел:
Запустите 3 горутины, каждая из которых выводит
числа от 1 до 5 с небольшой задержкой (например,
time.Sleep(time.Millisecond * 100)).
Убедитесь, что вывод перемешивается.
*/

func RunTask1() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			time.Sleep(time.Second)
			for i := 1; i <= 5; i++ {
				fmt.Printf("Goroutine number %d: %d\n", id, i)
			}
			wg.Done()
		}(i + 1)

	}

	wg.Wait()
}
