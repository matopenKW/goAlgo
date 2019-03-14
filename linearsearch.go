package main

import "fmt"

func main() {

	r := []int{1, 2, 3, 4, 5}
	fmt.Println(search(9, r), "番目")
	fmt.Println(search(9, []int{1, 2, 3}), "番目")
	fmt.Println(search(3, []int{1, 2, 3}), "番目")
	fmt.Println(search(1, []int{1, 2, 3}), "番目")
}

func search(target int, cin []int) int {
	list := append(cin, target)
	indexNo := 0
	for _, val := range list {
		if val == target {
			break
		}
		indexNo++
	}
	if indexNo >= len(list)-1 {
		return -1
	} else {
		return indexNo + 1
	}
}
