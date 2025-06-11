// Разработайте утилиту, которая бы построчно читала текстовый файл и удаляла
// пробелы из каждой строки, используя функцию strings.TrimSpace().

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("usage: replaceStringsInFile <filename>")
		return
	}

	filename := args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error open file", err)
		return
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("error reading file", err)
			return
		}
		newLine := strings.TrimSpace(line)
		fmt.Println(newLine)
	}
}
