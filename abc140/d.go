package main

import "fmt"

func main() {
	var N, K int
	var S string
	fmt.Scan(&N, &K)
	fmt.Scan(&S)

	sep := 1
	cur := rune(S[0])
	for _, c := range S {
		if cur != c {
			sep++
			cur = c
		}
	}
	for i := 0; i < K; i++ {
		switch {
		case sep >= 3:
			sep -= 2
		case sep >= 2:
			sep--
		}
	}

	fmt.Println(N - 1 - (sep - 1))
}
