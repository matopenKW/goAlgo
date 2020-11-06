package main

import (
	"fmt"
	"math/rand"
)

func combSort(numbers []int) []int {

	gap := len(numbers)
	swapped := true

	for gap != 1 || swapped {
		gap = int(float64(gap) / 1.3)
		if gap < 1 {
			gap = 1
		}

		swapped = false

		i := 0
		l := len(numbers) - gap
		for i < l {
			if numbers[i] > numbers[i+gap] {
				numbers[i], numbers[i+gap] = numbers[i+gap], numbers[i]
				swapped = true
			}
			i++
		}

	}

	return numbers
}

func main() {
	// 0~999までのランダムな数値を10個生成
	nums := rand.Perm(1000)[:10]

	fmt.Println(combSort(nums))
}
