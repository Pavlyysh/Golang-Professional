package main

import "fmt"

func main() {
	v := 12.12
	fmt.Println(getPtr(&v))
}

func getPtr(x *float64) float64 {
	return *x * *x
}
