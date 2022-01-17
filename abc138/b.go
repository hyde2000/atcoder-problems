package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	reciprocal := 0.0
	for i := 0; i < N; i++ {
		var A float64
		fmt.Scan(&A)
		reciprocal += 1.0 / A
	}
	fmt.Println(1.0 / reciprocal)
}
