package main

import (
	"fmt"
	"math/rand"
	"time"
)

func shellSort(numbers []int) []int {
	len := len(numbers)
	gap := len / 2
	for gap > 0 {
		for i, _ := range numbers[gap-1:] {
			temp := numbers[i]
			j := i
			for j >= gap && numbers[j-gap] > temp {
				numbers[j] = numbers[j-gap]
				j -= gap
			}
			numbers[j] = temp
		}
		gap /= 2
	}

	return numbers
}

func main() {
	// 0~999までのランダムな数値を生成
	rand.Seed(time.Now().UnixNano())
	nums := rand.Perm(1000)[:10]

	fmt.Println(shellSort(nums))
}
