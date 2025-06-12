package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go writeToChannel(ch, 10)
	time.Sleep(time.Second)
}

func writeToChannel(ch chan int, x int) {
	fmt.Println(x)

	ch <- x // здесь функция блокируется
	close(ch)
	fmt.Println(x)
}
