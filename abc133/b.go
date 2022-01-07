package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func dist(a, b *[]int) int {
	var d int
	for i, v := range *a {
		dd := v - (*b)[i]
		d += dd * dd
	}

	return d
}

func isSquare(v int) bool {
	f := math.Sqrt(float64(v))

	return f == math.Round(f)
}

func main() {
	var N, D int
	rr := bufio.NewReader(os.Stdin)
	fmt.Fscan(rr, &N, &D)

	x := make([][]int, N)
	for i := 0; i < N; i++ {
		xx := make([]int, D)
		for j := 0; j < D; j++ {
			fmt.Fscan(rr, &xx[j])
		}
		x[i] = xx
	}

	var v int
	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			if isSquare(dist(&x[i], &x[j])) {
				v++
			}
		}
	}
	fmt.Println(v)
}
