package main

import (
	"fmt"
	"strings"
)

var (
	g   = make([][]int, 20000)
	res = make([]int, 20000)
)

func dfs(i, p int) {
	for _, x := range g[i] {
		if x != p {
			res[x] += res[i]
			dfs(x, i)
		}
	}
}

func main() {
	var N, Q int
	fmt.Scan(&N, &Q)

	for i := 0; i < N-1; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		a--
		b--
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	for i := 0; i < Q; i++ {
		var p, x int
		fmt.Scan(&p, &x)
		p--
		res[p] += x
	}

	dfs(0, 0)
	fmt.Println(strings.Trim(fmt.Sprint(res[:N]), "[]"))
}
