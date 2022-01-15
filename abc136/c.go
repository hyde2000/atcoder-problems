package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	if N == 1 {
		fmt.Println("Yes")
		return
	}

	heights := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&heights[i])
	}

	for i := N - 1; i > 0; i-- {
		if heights[i] >= heights[i-1] {
			continue
		}

		if (heights[i-1] - heights[i]) > 1 {
			fmt.Println("No")
			return
		}
		heights[i-1]--
	}
	fmt.Println("Yes")
}
