package main

import (
	"fmt"
	"os"
)

func main() {
	var N int
	fmt.Scan(&N)

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if i*j == N {
				fmt.Println("Yes")
				os.Exit(0)
			}
		}
	}
	fmt.Println("No")
}
