package main

import "fmt"

func main() {
	var N int
	var S string
	fmt.Scan(&N)
	fmt.Scan(&S)
	res := 1
	s := []rune(S)
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			continue
		}
		res++
	}

	fmt.Println(res)
}
