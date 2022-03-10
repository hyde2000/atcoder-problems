package main

import "fmt"

func main() {
	var N int
	var S, T string
	fmt.Scan(&N)
	fmt.Scan(&S, &T)

	ans := ""
	for i := 0; i < N; i++ {
		ans += S[i : i+1]
		ans += T[i : i+1]
	}
	fmt.Println(ans)
}
