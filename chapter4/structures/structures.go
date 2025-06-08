package main

import "fmt"

type myStruct struct {
	Name    string
	Surname string
	Height  int32
}

func main() {
	s1 := createStruct("Paul", "Gazukin", 185)
	fmt.Println(s1.Name)

	x1 := 1
	x2 := 2
	x3 := 3
	x1, x2, x3 = x3, x1, x2
	fmt.Println(x1, x2, x3)
}

func createStruct(n, s string, h int32) *myStruct {
	if h > 300 {
		h = 0
	}
	return &myStruct{n, s, h}
}
