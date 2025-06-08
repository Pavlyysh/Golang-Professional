package main

import (
	"container/heap"
	"fmt"
)

type heapFloat64 []float64

func main() {
	myHeap := &heapFloat64{1.2, 4.2, 3.1, -1.6}
	heap.Init(myHeap)
	size := len(*myHeap)
	fmt.Printf("heap size: %d\n", size)
	fmt.Printf("%v\n", myHeap)

	myHeap.Push(-12.2)
	myHeap.Push(-54.1)
	fmt.Printf("Heap size: %d\n", len(*myHeap))
	fmt.Printf("%v\n", myHeap)
	heap.Init(myHeap)
	fmt.Printf("%v\n", myHeap)
}

func (n *heapFloat64) Pop() interface{} {
	old := *n
	x := old[len(old)-1]
	new := old[0 : len(old)-1]
	*n = new
	return x
}

func (n *heapFloat64) Push(x interface{}) {
	*n = append(*n, x.(float64))
}

func (n heapFloat64) Len() int {
	return len(n)
}

func (n heapFloat64) Less(a, b int) bool {
	return n[a] < n[b]
}

func (n heapFloat64) Swap(a, b int) {
	n[a], n[b] = n[b], n[a]
}
