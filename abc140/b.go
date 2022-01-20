package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N)
	B := make([]int, N)
	C := make([]int, N-1)

	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}
	for i := 0; i < N; i++ {
		fmt.Scan(&B[i])
	}
	for i := 0; i < N-1; i++ {
		fmt.Scan(&C[i])
	}
	ans := 0
	for i := 0; i < N; i++ {
		a := A[i]
		ans += B[a-1]
		if i > 0 {
			prevA := A[i-1]
			if prevA == a-1 {
				ans += C[prevA-1]
			}
		}
	}
	fmt.Println(ans)
}
