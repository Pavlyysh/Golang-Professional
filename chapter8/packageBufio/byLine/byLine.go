// Построчное чтение текстового файла

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: byLine <file1> [<file2> ...]\n")
		return
	}

	for _, file := range flag.Args() {
		err := lineByLine(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func lineByLine(file string) error {
	var err error

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Println(file)
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			fmt.Printf("END OF FILE\n\n")
			break
		} else if err != nil {
			fmt.Printf("Error reading file %s", err)
			return err
		}
		fmt.Print(line)
	}
	return nil
}
