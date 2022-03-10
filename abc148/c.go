package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	g := gds(A, B)
	ans := A * B / g
	fmt.Println(ans)
}

func gds(a, b int) int {
	for b > 0 {
		r := a % b
		a = b
		b = r
	}

	return a
}
