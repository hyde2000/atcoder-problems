package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	count := 0
	numbers := make([]int, N)
	for i := range numbers {
		fmt.Scan(&numbers[i])
		if numbers[i] != i+1 {
			count += 1
		}
	}

	if count < 3 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
