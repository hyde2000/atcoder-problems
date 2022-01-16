package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	ans := 0
	anagrams := map[string]int{}
	for i := 0; i < N; i++ {
		var s string
		fmt.Scan(&s)
		anagram := []rune(s)
		sort.Slice(anagram, func(i, j int) bool { return anagram[i] < anagram[j] })
		ans += anagrams[string(anagram)]
		anagrams[string(anagram)]++
	}
	fmt.Println(ans)
}
