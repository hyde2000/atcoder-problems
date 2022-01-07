package main

import "fmt"

func main() {
	var N, A, B int
	_, _ = fmt.Scan(&N, &A, &B)
	train := N * A
	taxi := B
	if train < taxi {
		fmt.Println(train)
	} else {
		fmt.Println(taxi)
	}
}
