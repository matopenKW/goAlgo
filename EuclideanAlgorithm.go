package main

import "fmt"

func main() {
	print(12, 8)
	print(125, 25)
}

func print(a, b int) {
	fmt.Printf("%vと%vの最大公約数→%v\n", a, b, gcd(a, b))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}
