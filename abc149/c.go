package main

import (
	"fmt"
	"math/big"
)

func main() {
	var X int64
	fmt.Scan(&X)

	for {
		if big.NewInt(X).ProbablyPrime(0) {
			break
		}
		X++
	}
	fmt.Println(X)
}
