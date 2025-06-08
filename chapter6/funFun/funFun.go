package main

import "fmt"

func main() {
	fmt.Println("function1:", funFun(function1, 123))
	fmt.Println("function2:", funFun(function2, 123))
	fmt.Println("Inline:", funFun(func(i int) int { return i * i * i }, 123))
}

func function1(x int) int {
	return x + x
}

func function2(x int) int {
	return x * x
}

func funFun(f func(int) int, v int) int {
	return f(v)
}
