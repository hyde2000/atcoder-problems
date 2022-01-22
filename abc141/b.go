package main

import (
	"fmt"
	"os"
)

func main() {
	var S string
	fmt.Scan(&S)
	for i, s := range S {
		str := fmt.Sprintf("%c", s)
		if i%2 == 0 {
			if str == "L" {
				fmt.Println("No")
				os.Exit(0)
			}
		} else {
			if str == "R" {
				fmt.Println("No")
				os.Exit(0)
			}
		}
	}
	fmt.Println("Yes")
}
