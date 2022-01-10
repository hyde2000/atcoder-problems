package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	fmt.Scan(&N)

	rr := bufio.NewReader(os.Stdin)
	numbers := make([]int, N, N)

	m1 := 0
	for i := 0; i < N; i++ {
		var W int
		fmt.Fscan(rr, &W)
		numbers[i] = W
		if numbers[i] > m1 {
			m1 = numbers[i]
		}
	}

	c := 0
	m2 := 0
	for j := 0; j < N; j++ {
		if numbers[j] == m1 {
			c++
		}
		if numbers[j] > m2 && numbers[j] < m1 {
			m2 = numbers[j]
		}
	}

	for i := 0; i < N; i++ {
		if numbers[i] == m1 && c == 1 {
			fmt.Println(m2)
		} else {
			fmt.Println(m1)
		}
	}
}
