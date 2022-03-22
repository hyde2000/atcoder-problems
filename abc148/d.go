package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	idx := 1
	ans := 0
	for i := 1; i <= N; i++ {
		var a int
		fmt.Scan(&a)

		if a == idx {
			idx++
		} else {
			ans++
		}
	}

	if idx == 1 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}
