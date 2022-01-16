package main

import "fmt"

func main() {
	var K, X int
	fmt.Scan(&K, &X)

	for i := X - (K - 1); i < X+K; i++ {
		fmt.Print(i, " ")
	}
}
