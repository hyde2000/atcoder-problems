package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	numbers := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&numbers[i])
	}
	temp := make(map[int]int, N)
	for i, v := range numbers {
		temp[v-1] = i + 1
	}
	for i := 0; i < N; i++ {
		fmt.Print(temp[i], " ")
	}
}
