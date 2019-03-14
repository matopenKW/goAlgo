package main

func main() {

}

func binarysearch(high, low, target int, list []int) int {

	if high == target && low == target {
		return target
	} else if high >= low {
		return target
	}
	return target

}
