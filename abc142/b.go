package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	res := 0
	for i := 0; i < N; i++ {
		var H int
		fmt.Scan(&H)
		if H >= K {
			res++
		}
	}
	fmt.Println(res)
}
