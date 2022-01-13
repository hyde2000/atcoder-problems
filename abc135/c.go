package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	a := make([]int, N+1)
	b := make([]int, N)
	for i := 0; i < N+1; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < N; i++ {
		fmt.Scan(&b[i])
	}

	ans := 0
	for i := 0; i < N; i++ {
		if a[i] <= b[i] {
			ans += a[i]
			b[i] -= a[i]
		} else {
			ans += b[i]
			continue
		}

		if a[i+1] <= b[i] {
			ans += a[i+1]
			a[i+1] = 0
		} else {
			ans += b[i]
			a[i+1] -= b[i]
		}
	}
	fmt.Println(ans)
}
