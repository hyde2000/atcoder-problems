package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	ans := N - 1
	for i := 1; i*i <= N; i++ {
		if N%i == 0 {
			ans = min(ans, i+N/i-2)
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
