package main

import "fmt"

func main() {
	var N, K, Q int
	fmt.Scan(&N, &K, &Q)

	scores := make([]int, N)
	for i := 0; i < Q; i++ {
		var A int
		fmt.Scan(&A)
		scores[A-1]++
	}

	for i := range scores {
		if scores[i]+K > Q {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
