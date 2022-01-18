package main

import "fmt"

// 0 1 3 6 10 ...
func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(N * (N - 1) / 2)
}
