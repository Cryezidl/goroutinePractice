package tasks

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
Сумма чисел
Разделите массив чисел на две части и
вычислите сумму каждой части в отдельной горутине.
Выведите общую сумму.
*/

func RunTask2() {
	n := rand.Intn(11)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = rand.Intn(100)
	}
	var wg sync.WaitGroup
	partsCount := 2
	sliceChunk := len(nums) / partsCount
	total := 0

	for i := 0; i < partsCount; i++ {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			sum := 0
			end := start + sliceChunk
			length := len(nums)
			if end > length || i+1 == partsCount {
				end = length
			}

			for _, val := range nums[start:end] {
				sum += val
			}
			total += sum
		}(i * sliceChunk)
	}

	wg.Wait()
	totalTwo := 0
	for _, val := range nums {
		totalTwo += val
	}
	fmt.Println(total, totalTwo, nums)
}
