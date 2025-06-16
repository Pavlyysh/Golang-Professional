package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("need one int")
		return
	}

	times, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	cc := make(chan chan int)

	for i := 1; i < times+1; i++ {
		f := make(chan bool)
		go f1(cc, f)
		ch := <-cc
		ch <- i
		for sum := range ch {
			fmt.Print("Sum(", i, ")=", sum)
		}
		fmt.Println()
		time.Sleep(time.Second)
		close(f)
	}
}

func f1(cc chan chan int, f chan bool) {
	c := make(chan int)
	cc <- c
	defer close(c)

	sum := 0
	select {
	case x := <-c:
		for i := 0; i <= x; i++ {
			sum += i
		}
		c <- sum
	case <-f:
		return
	}
}
