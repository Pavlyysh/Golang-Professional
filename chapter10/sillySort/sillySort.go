package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println(os.Args[0], "n1, n2, [n]")
		return
	}

	var wg sync.WaitGroup

	for _, arg := range args[1:] {
		num, err := strconv.Atoi(arg)
		if err != nil || num < 0 {
			fmt.Print(". ")
			continue
		}
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			time.Sleep(time.Duration(n) * time.Second)
			fmt.Print(n, " ")
		}(num)
	}
	wg.Wait()
	fmt.Println()
}
