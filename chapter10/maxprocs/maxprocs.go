package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("GOMAXPROCS: %d\n", getGOMAXPROCS())
}

func getGOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}
