package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	mutex sync.Mutex
	v1    int
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("need an int")
		return
	}

	numGR, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup

	fmt.Printf("%d ", read())

	for i := 0; i < numGR; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			change(i)
			fmt.Printf("-> %d", read())
		}(i)
	}
	wg.Wait()
	fmt.Printf("-> %d\n", read())
}

func change(i int) {
	mutex.Lock()
	time.Sleep(time.Second)
	v1 += 1
	if v1%10 == 0 {
		v1 -= 10 * i
	}
	mutex.Unlock()
}

func read() int {
	mutex.Lock()
	a := v1
	mutex.Unlock()
	return a
}
