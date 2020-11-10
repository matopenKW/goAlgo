package main

import (
	"errors"
	"fmt"
	// "math/rand"
	// "time"
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

func bucketSort(numbers []int) []int {
	max, err := max(numbers)
	len := len(numbers)
	if err != nil {
		panic(err)
	}

	size := max / len
	buckets := make([][]int, 0, 0)
	for _, v := range numbers {
		i := v / size
		fmt.Println(i)
		if i != size {
			buckets[i] = append(buckets[i], v)
		} else {
			buckets[i-1] = append(buckets[i], v)
		}
	}

	for i := 0; i <= size; i++ {
		insertionSort(buckets[i])
	}

	result := make([]int, len, len)
	for i := 0; i <= size; i++ {
		result = append(result, buckets[i]...)
	}
	return numbers
}

func main() {
	// 0~999までのランダムな数値を生成
	// rand.Seed(time.Now().UnixNano())
	// nums := rand.Perm(1000)[:10]

	//fmt.Println(insertionSort(nums))

	fmt.Println(bucketSort([]int{1, 5, 3, 4, 2}))
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
