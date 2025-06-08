package main

import "fmt"

func main() {
	sq := returnPtr(10)
	fmt.Println("sq value:", *sq)
	fmt.Println("sq address:", sq)
}

func returnPtr(x int) *int {
	y := x * x
	return &y
}
