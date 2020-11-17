package main

import (
	"fmt"
	"math/rand"
	"time"
)

func margeSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	center := len(numbers) / 2
	left := margeSort(numbers[:center])
	right := margeSort(numbers[center:])

	result := make([]int, 0, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++

	}

	return result
}

func main() {
	// 0~999までのランダムな数値を生成
	rand.Seed(time.Now().UnixNano())
	nums := rand.Perm(1000)[:10]

	//fmt.Println(margeSort(nums))
	fmt.Println(nums)
	fmt.Println(margeSort([]int{5, 4, 1, 8, 7, 3, 2, 9}))

}
