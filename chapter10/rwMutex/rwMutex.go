package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var Password = secret{password: "myPassword"}

type secret struct {
	RWM      sync.RWMutex
	M        sync.Mutex
	password string
}

func main() {
	var showFunction = func(c *secret) string { return "" }
	if len(os.Args) != 2 {
		fmt.Println("using sync.RWMutex!")
		showFunction = show
	} else {
		fmt.Println("using sync.Mutex!")
		showFunction = showWithLock
	}

	var wg sync.WaitGroup

	fmt.Println("Pass:", showFunction(&Password))

	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Go Pass:", showFunction(&Password))
		}()
	}

	go func() {
		wg.Add(1)
		defer wg.Done()
		Change(&Password, "123456")
	}()

	wg.Wait()
	fmt.Println("password", showFunction(&Password))
}

func Change(c *secret, pass string) {
	c.RWM.Lock()
	fmt.Println("LChange")
	time.Sleep(10 * time.Second)
	c.password = pass
	c.RWM.Unlock()
}

func show(c *secret) string {
	c.RWM.RLock()
	defer c.RWM.RUnlock()
	fmt.Print("show")
	time.Sleep(3 * time.Second)
	return c.password
}

func showWithLock(c *secret) string {
	c.RWM.Lock()
	defer c.RWM.Unlock()
	fmt.Print("showWithLock")
	time.Sleep(3 * time.Second)
	return c.password
}
