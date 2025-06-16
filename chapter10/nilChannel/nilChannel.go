package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go add(ch)
	go send(ch)

	time.Sleep(3 * time.Second)
}

func add(ch chan int) {
	sum := 0
	t := time.NewTimer(time.Second)

	for {
		select {
		case input := <-ch:
			sum += input
		case <-t.C:
			ch = nil
			fmt.Println(sum)
		}
	}
}

func send(ch chan int) {
	for {
		ch <- rand.Intn(10)
	}
}
