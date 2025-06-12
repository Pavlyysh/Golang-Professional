package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go writeToChannel(ch, 10)
	time.Sleep(time.Second)
	fmt.Println("Read:", <-ch)
	time.Sleep(time.Second)

	_, ok := <-ch
	if ok {
		fmt.Println("channel is open")
	} else {
		fmt.Println("channel is closed")
	}
}

func writeToChannel(ch chan int, x int) {
	fmt.Println("1", x)
	ch <- x
	close(ch)
	fmt.Println("2", x)
}
