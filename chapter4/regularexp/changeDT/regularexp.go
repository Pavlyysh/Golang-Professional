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

// Пример посложнее с двумя регулярными выражениями
func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("give me a file")
		os.Exit(1)
	}
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file: %s\n", err)
		os.Exit(1)
	}
	defer f.Close()

	notAMatch := 0

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file: %s", err)
		}

		r1 := regexp.MustCompile(`.*\[(\d\d/\w+/\d\d\d\d:\d\d:\d\d.*)\] .*`)
		if r1.MatchString(line) {
			match := r1.FindStringSubmatch(line)
			d1, err := time.Parse("02/Jan/2006:15:04:05 -0700", match[1])
			if err == nil {
				newFormat := d1.Format(time.Stamp)
				fmt.Print(strings.Replace(line, match[1], newFormat, 1))
			} else {
				notAMatch++
			}
			continue
		}

		r2 := regexp.MustCompile(`.*\[(\w+\-\d\d-\d\d:\d\d:\d\d:\d\d.*)\] .*`)
		if r2.MatchString(line) {
			match := r2.FindStringSubmatch(line)
			d1, err := time.Parse("Jan-02-06:15:04:05 -0700", match[1])
			if err == nil {
				newFormat := d1.Format(time.Stamp)
				fmt.Print(strings.Replace(line, match[1], newFormat, 1))
			} else {
				notAMatch++
			}
			continue
		}
	}
	fmt.Println(notAMatch, "lines did not match")
}

/* Простой пример
func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("not enough arguments\n")
		os.Exit(1)
	}

	temp, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("first arg must be an int")
		os.Exit(1)
	}

	column := temp
	if column <= 0 {
		fmt.Println("column cannot be less than 1")
		os.Exit(1)
	}

	for _, filename := range os.Args[2:] {
		fmt.Println("\t\t", filename)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening file %s\n", err)
			continue
		}
		defer f.Close()

		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file %s\n", err)
			}
			data := strings.Fields(line)
			if len(data) >= column {
				fmt.Println(data[column-1])
			}
		}
	}
}
*/
