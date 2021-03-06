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
	start := time.Now().UnixNano()
	fmt.Println("start: ", start)

	// 0~999までのランダムな数値を生成
	rand.Seed(time.Now().UnixNano())
	nums := rand.Perm(1000000)

	fmt.Println(margeSort(nums)[:20])

	end := time.Now().UnixNano()
	fmt.Println("end  : ", end)
	fmt.Println("nano : ", end-start)

}
