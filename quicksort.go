package main

import "fmt"

func inputData() []int {
	return []int{1, 2, 3, 4, 5, 6, 7}
}

func main() {

	list := []int{4, 3, 7, 5, 3, 5, 6, 6, 4, 3}
	fmt.Println(list, "→", sort(list))

}

func sort(list []int) (ret []int) {

	if len(list) < 2 {
		return list
	}

	pivot := list[0]

	low := []int{}
	high := []int{}

	for _, v := range list[1:] {
		if v > pivot {
			high = append(high, v)
		} else {
			low = append(low, v)
		}
	}

	low = sort(low)
	high = sort(high)

	ret = append(low, pivot)
	ret = append(ret, high...)
	return ret

}
