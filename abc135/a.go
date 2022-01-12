package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	sum := A + B
	if sum%2 == 0 {
		fmt.Println(sum / 2)
	} else {
		fmt.Println("IMPOSSIBLE")
	}
}
