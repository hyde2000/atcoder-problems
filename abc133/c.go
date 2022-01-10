package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var L, R int
	rr := bufio.NewReader(os.Stdin)
	fmt.Fscan(rr, &L, &R)

	min := 2020
	if R-L > 2019 {
		fmt.Println(0)
	} else {
		for l := L; l <= R-1; l++ {
			for r := l + 1; r <= R; r++ {
				val := (l * r) % 2019
				if val < min {
					min = val
				}
			}
		}
		fmt.Println(min)
	}
}
