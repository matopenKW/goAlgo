package main

import (
	"fmt"
	"math/rand"
	"time"
)

func gnome_sort(numbers []int) []int {
	index := 0

	for index < len(numbers) {
		if index == 0 {
			index++
		}

		if numbers[index] >= numbers[index-1] {
			index++
		} else {
			numbers[index], numbers[index-1] = numbers[index-1], numbers[index]
			index--
		}
	}

	return numbers
}

func main() {
	// 0~999までのランダムな数値を生成
	rand.Seed(time.Now().UnixNano())
	nums := rand.Perm(1000)[:10]

	fmt.Println(gnome_sort(nums))
}
