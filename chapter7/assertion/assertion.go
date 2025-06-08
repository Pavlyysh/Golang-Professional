package main

import "fmt"

func main() {
	var myInt interface{} = 123

	k, ok := myInt.(int)
	if ok {
		fmt.Println("Success:", k)
	}

	v, ok := myInt.(float64)
	if ok {
		fmt.Println("Success:", v)
	} else {
		fmt.Println("Failed without panic")
	}

	i := myInt.(int)
	fmt.Println("Without checking:", i)

	b := myInt.(bool)
	fmt.Println(b)
}
