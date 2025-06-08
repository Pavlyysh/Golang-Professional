package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args
	var sum, count float64

	for i := range args {
		if args[i] == "END" {
			break
		}
		x, err := strconv.ParseFloat(args[i], 64)
		if err != nil {
			continue
		}
		sum += x
		count++
	}
	fmt.Println(sum / count)
}
