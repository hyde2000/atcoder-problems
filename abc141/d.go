package main

import (
	"container/heap"
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	items := &Heap{}
	for i := 0; i < N; i++ {
		var A int
		fmt.Scan(&A)
		heap.Push(items, A)
	}

	for i := 0; i < M; i++ {
		top := (*items)[0]
		heap.Pop(items)
		heap.Push(items, top/2)
	}

	res := 0
	for items.Len() > 0 {
		res += (*items)[0]
		heap.Pop(items)
	}
	fmt.Println(res)
}

type Heap []int

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i] > h[j] }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
