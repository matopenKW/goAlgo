package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insertionSort(numbers []int) []int {

	for i, _ := range numbers {
		if i == 0 {
			continue
		}
		temp := numbers[i]
		j := i - 1

		for j >= 0 && numbers[j] > temp {
			numbers[j+1] = numbers[j]
			j--
		}

		numbers[j+1] = temp
	}

	return numbers
}

func main() {
	// 0~999までのランダムな数値を生成
	rand.Seed(time.Now().UnixNano())
	nums := rand.Perm(1000)[:10]

	fmt.Println(insertionSort(nums))
}
