package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filename := ""
	args := os.Args
	if len(args) == 1 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	for i := 1; i < len(args); i++ {
		filename = args[i]
		err := printFile(filename)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func printFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error open file %s\n", err)
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		io.WriteString(os.Stdout, scanner.Text())
		io.WriteString(os.Stdout, "\n")
	}
	return nil
}
