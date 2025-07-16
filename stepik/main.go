package main

import (
	"fmt"
	"strings"
)

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go separateGopher(c0, c1)
	printGopher(c1)
}

func sourceGopher(downstream chan string) {
	vSlice := []string{"hello world", "a bad apple", "goodbye all"}
	for _, v := range vSlice {
		downstream <- v
	}
	close(downstream)
}

func separateGopher(upstream, downstream chan string) {
	for item := range upstream {
		for _, word := range strings.Fields(item) {
			downstream <- word
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}
