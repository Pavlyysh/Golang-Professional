package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("not enough args")
		return
	}

	arg1 := args[1]

	pattern := `^(2[0-9][0-9]|3[0-9][0-9]|400)$`

	matched, err := regexp.MatchString(pattern, arg1)
	if err != nil {
		fmt.Printf("error regexp: %v", err)
		return
	}

	if matched {
		fmt.Println("input is in gap 200-400")
	} else {
		fmt.Println("not in gap")
	}
}
