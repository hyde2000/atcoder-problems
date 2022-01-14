package main

import "fmt"

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)
	bottle1 := A - B
	if bottle1 >= C {
		fmt.Println(0)
	} else {
		fmt.Println(C - bottle1)
	}
}
