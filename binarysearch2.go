package main

import (
	"fmt"
)

func main() {

	serch := binarysearch2([]int{1, 2, 3, 4, 5, 6, 8})
	fmt.Println("anser:", serch(5), "番目")
	fmt.Println("anser:", serch(1), "番目")
	fmt.Println("anser:", serch(9), "番目")

	fmt.Println("anser:", serch(8), "番目")
}

func binarysearch2(list []int) func(key int) int {
	return func(key int) int {
		low := 0
		high := len(list)
		for low < high {
			mid := low + (high-low)/2
			fmt.Print("low:", low, " high:", high, " → ")
			if list[mid] == key {
				return mid
			} else if list[mid] > key {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		return -1
	}
}
