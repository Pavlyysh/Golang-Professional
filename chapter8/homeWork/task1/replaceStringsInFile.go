// Напишите программу Go, которая принимает три аргумента: имя текстового
// файла и две строки. Затем эта утилита должна заменить каждое вхождение
// первой строки в файле второй строкой. В целях безопасности окончательный
// результат выведите на экран, чтобы исходный текстовый файл остался без из-
// менений.

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
	if len(args) != 4 {
		fmt.Println("usage: replaceStringsInFile <filename> <oldString> <newString>")
		return
	}

	filename := args[1]
	oldString := args[2]
	newString := args[3]
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
		newLine := strings.Replace(line, oldString, newString, -1)
		fmt.Println(newLine)
	}
}
