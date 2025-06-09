// Посимвольное чтение текстового файла

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
		fmt.Printf("usage: byCharacter <file1> [<file2> ...]\n")
		return
	}

	for _, file := range flag.Args() {
		err := charByChar(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func charByChar(file string) error {
	var err error

	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("error opening file %s", err)
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
			fmt.Printf("error reading file %s", err)
			return err
		}

		for _, x := range line {
			fmt.Println(string(x))
		}
	}

	return nil
}
