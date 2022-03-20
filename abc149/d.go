package main

import "fmt"

func main() {
	var N, K, R, S, P int
	var T string
	fmt.Scan(&N, &K)
	fmt.Scan(&R, &S, &P)
	fmt.Scan(&T)

	var dp []bool
	ans := 0
	ts := []rune(T)
	for i := 0; i < N; i++ {
		if i >= K {
			if ts[i] == ts[i-K] && dp[i-K] {
				dp = append(dp, false)
				continue
			}
		}
		switch ts[i] {
		case 'r':
			ans += P
		case 's':
			ans += R
		case 'p':
			ans += S
		}
		dp = append(dp, true)
	}
	fmt.Println(ans)
}
