package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int
	fmt.Scan(&N)

	ans := 0
	for i := 1; i <= N; i++ {
		s := strconv.Itoa(i)
		if len(s)%2 == 1 {
			ans++
		}
	}
	fmt.Println(ans)
}
