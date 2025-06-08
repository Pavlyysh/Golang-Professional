package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	args := os.Args
	if len(args) == 1 {
		myString = "Give me one arg!"
	} else {
		myString = args[1]
	}

	io.WriteString(os.Stdout, "This is Standart output\n")
	io.WriteString(os.Stderr, myString)
	//io.WriteString(os.Stderr, "\n")
}
