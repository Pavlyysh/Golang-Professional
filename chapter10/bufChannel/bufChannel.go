package main

import "fmt"

func main() {
	numbers := make(chan int, 5)
	count := 10

	for i := 0; i < count; i++ {
		select {
		case numbers <- i:
		default:
			fmt.Println("Not enough space for", i)
		}
	}

	for i := 0; i < count+5; i++ {
		select {
		case num := <-numbers:
			fmt.Println(num)
		default:
			fmt.Println("Nothing more to be done!")
		}
	}
}
