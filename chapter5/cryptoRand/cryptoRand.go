package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

func generateBytes(n int64) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generatePass(s int64) (string, error) {
	b, err := generateBytes(s)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func main() {
	var LENGTH int64 = 8

	arguments := os.Args

	switch len(arguments) {
	case 2:
		LENGTH, _ = strconv.ParseInt(os.Args[1], 10, 64)
		if LENGTH <= 0 {
			LENGTH = 8
		}
	default:
		fmt.Println("Using default values")
	}

	myPass, err := generatePass(LENGTH)
	if err != nil {
		fmt.Printf("error creating password: %v\n", err)
		return
	}

	fmt.Println(myPass[0:LENGTH])
}
