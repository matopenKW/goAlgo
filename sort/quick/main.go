package main

import (
	"fmt"
	"math/rand"
	"time"
)

func partition(numbers []int, low, high int) int {
	i := low - 1
	pivot := numbers[high]

	for j := low; j < high; j++ {
		if numbers[j] <= pivot {
			i++
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}
	numbers[i+1], numbers[high] = numbers[high], numbers[i+1]
	return i + 1
}

func quickSort(numbers []int, low, high int) {
	if low < high {
		partitionIndex := partition(numbers, low, high)
		quickSort(numbers, low, partitionIndex-1)
		quickSort(numbers, partitionIndex+1, high)
	}
}

func quickSortWrpper(numbers []int) []int {
	quickSort(numbers, 0, len(numbers)-1)
	return numbers
}

func main() {
	// 0~999までのランダムな数値を生成
	rand.Seed(time.Now().UnixNano())
	nums := rand.Perm(1000)[:10]

	fmt.Println(quickSortWrpper(nums))

	//	fmt.Println(quickSortWrpper([]int{1, 8, 3, 9, 4, 5, 7}))
}
