package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		print(i)
	}
}

func print(n int) {
	fmt.Printf("%vの場合：%v\n", n, cal(n))
}

/*
 f(n)
 n = 1 → 1
 n > 1 → 2 * f(n-1) + 1
*/

func cal(n int) int {
	if n > 1 {
		return 2*cal(n-1) + 1
	} else {
		return 1
	}

}
