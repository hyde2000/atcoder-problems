package main

import "fmt"

func gcd(x, y int) int {
	for y != 0 {
		r := x % y
		x = y
		y = r
	}
	return x
}

func main() {
	var A, B int
	fmt.Scan(&A, &B)
	g := gcd(A, B)
	ans := 1
	for i := 2; i <= 100010; i++ {
		if g%i == 0 {
			ans++
			for g%i == 0 {
				g /= i
			}
		}
	}
	if g > 1 {
		ans++
	}
	fmt.Println(ans)
}
