// Чтение заданного количества данных

package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("<buffer size> <filename>")
		return
	}

	bufferSize, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("invalid buffer size %s", err)
		return
	}

	file := args[2]
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("error opening file %s", err)
		return
	}
	defer f.Close()

	for {
		readData := readSize(f, bufferSize)
		if readData != nil {
			fmt.Println(string(readData))
		} else {
			break
		}
	}

}

func readSize(f *os.File, size int) []byte {
	buffer := make([]byte, size)

	n, err := f.Read(buffer)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return buffer[0:n]
}
