package main

import "fmt"

func inputData() []int {
	return []int{1, 2, 3, 4, 5, 6, 7}
}

func main() {
	var list = inputData()
	list = quickSort(list)
	print(list)
}

func quickSort(list []int) []int {
	after := []int{}
	sort(&list, 0, len(list))
	return after
}

func sort(list *[]int, fir, end int) {
	if fir == end {
		return
	}
	p := pivot(list, fir, end)
	fmt.Print(p)
}

func pivot(list *[]int, fir, end int) int {
	k := fir + 1
	for k <= end {
		k++
	}
	return fir
}

func print(list []int) {
	for i, num := range list {
		fmt.Printf("%d: %v\n", i, num)
	}
}
