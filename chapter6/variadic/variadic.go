package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	varFunc(arguments...)
	sum := oneByOne("Adding numbers", 1, 2, 3, 4, 5, -12, 21)
	fmt.Println("Sum:", sum)
	s := []int{1, 2, 35, -40}
	sum = oneByOne("Adding numbers 2", s...)
	fmt.Println(s)
}

func varFunc(input ...string) {
	fmt.Println(input)
}

func oneByOne(message string, s ...int) int {
	fmt.Println(message)
	sum := 0
	for i, a := range s {
		fmt.Println(i, a)
		sum += a
	}
	s[0] = -1000
	return sum
}
