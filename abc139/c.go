package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	heights := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&heights[i])
	}

	var max, count int
	for i := 1; i < N; i++ {
		if heights[i] > heights[i-1] {
			if count > max {
				max = count
			}
			count = 0
		} else {
			count++
		}
		if count > max {
			max = count
		}
	}
	fmt.Println(max)
}
