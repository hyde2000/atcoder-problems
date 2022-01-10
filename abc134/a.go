package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var r int

	rr := bufio.NewReader(os.Stdin)
	fmt.Fscan(rr, &r)

	fmt.Println(3 * r * r)
}
