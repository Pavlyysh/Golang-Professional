// Чтение текстового файла по словам

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage byWord <file1> [<file2> ...]\n")
		return
	}

	for _, file := range flag.Args() {
		err := wordByWord(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func wordByWord(file string) error {
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
			fmt.Printf("error reading file %s", err)
			return err
		}
		r := regexp.MustCompile("[^\\s]+")
		words := r.FindAllString(line, -1)
		for i := 0; i < len(words); i++ {
			fmt.Println(words[i])
		}
	}

	return nil
}
