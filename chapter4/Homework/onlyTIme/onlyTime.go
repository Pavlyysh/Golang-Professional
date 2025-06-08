package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("no file")
		return
	}

	filename := args[1]

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file: %v", err)
			return
		}

		r1 := regexp.MustCompile(`.*\[(\d\d/\w+/\d\d\d\d:\d\d:\d\d.*)\] .*`)
		if r1.MatchString(line) {
			match := r1.FindStringSubmatch(line)
			d1, err := time.Parse("02/Jan/2006:15:04:05 -0700", match[1])
			if err == nil {
				newFormat := d1.Format(time.TimeOnly)
				fmt.Print(strings.Replace(line, match[1], newFormat, 1))
			}

		}
	}
}
