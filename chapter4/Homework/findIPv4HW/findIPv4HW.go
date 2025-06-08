package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"regexp"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("not enough args")
		os.Exit(1)
	}

	countIPMap := make(map[string]int)

	for _, filename := range args[1:] {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening file: %s", err)
			os.Exit(-1)
		}
		defer f.Close()

		reader := bufio.NewReader(f)
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file: %s", err)
				break
			}

			ip := findIP(line)
			trial := net.ParseIP(ip)
			if trial.To4() == nil {
				continue
			}

			countIPMap[ip]++
		}
	}

	max := 0
	addr := make([]string, 0)
	for key, val := range countIPMap {
		if val > max {
			max = val
			addr = []string{key}
			continue
		}
		if val == max {
			addr = append(addr, key)
			continue
		}
	}

	for _, address := range addr {
		fmt.Println("key", address, "value", countIPMap[address])
	}

}

func findIP(input string) string {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(input)
}
