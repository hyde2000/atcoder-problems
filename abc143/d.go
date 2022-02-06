package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)
	L := make([]int, N)
	for i := range L {
		fmt.Scan(&L[i])
	}
	sort.Ints(L)

	count := 0
	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			ab := L[i] + L[j]
			r := sort.SearchInts(L, ab)
			l := i + 1
			count += r - l
		}
	}

	fmt.Println(count)
}
