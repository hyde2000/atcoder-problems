package main

import (
	"fmt"
)

func main() {
	var A, B, K int
	fmt.Scan(&A, &B, &K)

	// A
	if A > K {
		A -= K
		K = 0
	} else {
		K -= A
		A = 0
	}
	// B
	if B < K {
		B = 0
	} else {
		B -= K
	}
	fmt.Println(A, B)
}
