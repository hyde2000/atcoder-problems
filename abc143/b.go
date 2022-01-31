package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	d := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&d[i])
	}
	res := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			res += d[i] * d[j]
		}
	}
	fmt.Println(res)
}
