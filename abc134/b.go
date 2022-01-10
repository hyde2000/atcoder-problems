package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var N, D float64
	rr := bufio.NewReader(os.Stdin)
	fmt.Fscan(rr, &N, &D)

	i := D*2 + 1
	fmt.Println(math.Ceil(N / i))
}
