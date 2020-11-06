package main

import (
	"fmt"
	"math/rand"
)

func cocktailSort(numbers []int) []int {
	swapped := true
	start := 0
	end := len(numbers) - 1
	for swapped {

		swapped = false

		i := start
		for i < end {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				swapped = true
			}
			i++
		}

		if !swapped {
			break
		}
		swapped = false
		end--

		i = end - 1
		for i > start-1 {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				swapped = true
			}
			i--
		}
		start++
	}

	return numbers
}

func main() {
	// 0~999までのランダムな数値を10桁まで生成
	nums := rand.Perm(1000)[:10]

	fmt.Println(cocktailSort(nums))
}
