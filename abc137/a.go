package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	plus := A + B
	minus := A - B
	times := A * B

	max := plus
	if minus > max {
		max = minus
	}
	if times > max {
		max = times
	}

	fmt.Println(max)
}
