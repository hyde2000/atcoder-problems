package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	B := make([]int, N-1)
	for i := range B {
		fmt.Scan(&B[i])
	}

	res := 0
	for i := 0; i < N-2; i++ {
		res += min(B[i], B[i+1])
	}
	fmt.Println(res + B[0] + B[N-2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
