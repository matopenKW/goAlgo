package main

import (
	"fmt"
	"math/rand"
)

func selectionSort(numbers []int) []int {
	for i, _ := range numbers {
		minIdx := i
		for j := i + 1; j < len(numbers); j++ {
			if numbers[minIdx] > numbers[j] {
				minIdx = j
			}
		}
		numbers[i], numbers[minIdx] = numbers[minIdx], numbers[i]
	}

	return numbers
}

func main() {
	// 0~999までのランダムな数値を生成
	nums := rand.Perm(1000)[:10]

	fmt.Println(selectionSort(nums))
	fmt.Println(selectionSort([]int{2, 1, 13, 4, 56, 33, 7}))

}
