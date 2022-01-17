package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	Vs := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&Vs[i])
	}
	sort.Slice(Vs, func(i, j int) bool { return Vs[i] < Vs[j] })
	ans := float64(Vs[0])
	for i := 1; i < N; i++ {
		ans = (ans + float64(Vs[i])) / 2.0
	}
	fmt.Println(ans)
}
