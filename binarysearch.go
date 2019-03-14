package main

import (
	"fmt"
	"strconv"
)

func main() {

	list := []int{1, 2, 3, 4, 5, 6, 8}
	fmt.Println("anser:", binarysearch(5, 0, len(list)-1, list), "番目")
	fmt.Println("anser:", binarysearch(8, 0, len(list)-1, list), "番目")

	list = []int{1, 2, 3, 4, 5, 6, 8, 11, 15, 17, 19, 20, 25, 26, 28, 30, 34, 68}
	fmt.Println("anser:", binarysearch(19, 0, len(list)-1, list), "番目")
	fmt.Println("anser:", binarysearch(0, 0, len(list)-1, list), "番目")
	fmt.Println("anser:", binarysearch(555, 0, len(list)-1, list), "番目")
	fmt.Println("anser:", binarysearch(-5, 0, len(list)-1, list), "番目")

	// 年齢当てゲーム！！
	// 20歳～36歳のAさんの年齢を当てるのに必要な質問回数は？？
	fmt.Println("anser:", aGesCheck(0, 20, 36))
}

func binarysearch(key, low, high int, list []int) int {
	fmt.Print("low:", low, " high:", high, " → ")
	if high < low {
		return -1
	} else {
		mid := low + (high-low)/2
		if list[mid] == key {
			return mid
		} else if list[mid] > key {
			return binarysearch(key, low, mid-1, list)
		} else if list[mid] < key {
			return binarysearch(key, mid+1, high, list)
		} else {
			return -1
		}
	}
}

func aGesCheck(cnt, low, high int) string {
	fmt.Print("low:", low, " high:", high, " → ")
	if low-high >= 1 {
		return strconv.Itoa(cnt) + "回"
	} else {
		cnt++
		mid := low + (high-low)/2
		if mid > low {
			return aGesCheck(cnt, low, mid-1)
		} else {
			return aGesCheck(cnt, mid+1, high)
		}
	}
}
