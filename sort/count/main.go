package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func countingSort(numbers []int) []int {

	maxNum, err := max(numbers)
	if err != nil {
		panic(err)
	}

	counts := make([]int, maxNum+1)
	result := make([]int, len(numbers))

	for _, v := range numbers {
		counts[v] = counts[v] + 1
	}

	for i, _ := range counts {
		if i == 0 {
			continue
		}
		counts[i] += counts[i-1]
	}

	i := len(numbers) - 1
	for i >= 0 {
		index := numbers[i]
		result[counts[index]-1] = numbers[i]
		counts[index] -= 1
		i--
	}

	return result
}

func main() {
	// 0~999までのランダムな数値を生成
	rand.Seed(time.Now().UnixNano())
	nums := rand.Perm(1000)[:10]

	fmt.Println(countingSort(nums))
}

func max(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("numbers is Nil")
	}
	max := numbers[0]
	for _, v := range numbers {
		if max < v {
			max = v
		}
	}
	return max, nil
}
