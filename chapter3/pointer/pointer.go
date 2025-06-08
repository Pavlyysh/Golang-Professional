package main

import "fmt"

func main() {
	a := -10
	b := 25

	pA := &a
	pB := &b

	fmt.Println("pA memory:", pA)
	fmt.Println("pB memory:", pB)
	fmt.Println("pA value:", *pA)
	fmt.Println("pB value:", *pB)

	*pA = 123456
	*pA--
	fmt.Println("a", a)

	getPointer(pB)
	fmt.Println("b", b)

	k := returnPointer(12)
	fmt.Println(k)
	fmt.Println(*k)
}

func getPointer(n *int) {
	*n = *n * *n
}

func returnPointer(n int) *int {
	v := n * n
	return &v
}
