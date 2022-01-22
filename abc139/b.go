package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	ans := 0
	hole := 1
	for {
		if hole >= B {
			break
		}
		hole += A - 1
		ans++
	}
	fmt.Println(ans)
}
