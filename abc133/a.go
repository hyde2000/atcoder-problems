package main

import "fmt"

var N, A, B int

func main() {
	_, _ = fmt.Scan(&N, &A, &B)
	train := N * A
	taxi := B
	if train < taxi {
		fmt.Println(train)
	} else {
		fmt.Println(taxi)
	}
}
