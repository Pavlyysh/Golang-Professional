package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var timeout = time.Duration(1 * time.Second)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s URL(required) timeout(optional)\n", filepath.Base(os.Args[0]))
	}
	if len(os.Args) == 3 {
		temp, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Using default timeout")
		} else {
			timeout = time.Duration(time.Duration(temp) * time.Second)
		}
	}

	URL := os.Args[1]

	client := http.Client{
		Timeout: timeout,
	}

	data, err := client.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		defer data.Body.Close()
		_, err := io.Copy(os.Stdout, data.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
