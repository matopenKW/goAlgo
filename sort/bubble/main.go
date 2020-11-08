package main

import (
	"fmt"
	"math/rand"
)

func bubbleSort(numbers []int) []int {
	for i, _ := range numbers {
		for j, _ := range numbers[0 : len(numbers)-1-i] {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	return numbers
}

func main() {
	// 0~999までのランダムな数値を生成
	rand.Seed(time.Now().UnixNano())
	nums := rand.Perm(1000)[:10]

	fmt.Println(bubbleSort(nums))
}
