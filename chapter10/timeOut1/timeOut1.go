package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		c1 <- "c1 complete"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("c1 out of time")
	}

	c2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		c2 <- "c2 complete"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(4 * time.Second):
		fmt.Println("c2 out of time")
	}
}
