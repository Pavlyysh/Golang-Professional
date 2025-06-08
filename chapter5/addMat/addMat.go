package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Println("need 3 arguments")
		return
	}

	var row, col int
	row, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("need an int", arguments[1])
		return
	}

	col, err = strconv.Atoi(arguments[2])
	if err != nil {
		fmt.Println("need an int", arguments[2])
		return
	}

	fmt.Printf("Using %dx%d arrays\n", row, col)
	if row <= 0 || col <= 0 {
		fmt.Println("Need positive matrix dimensions")
		return
	}

	m1 := make([][]int, row)
	m2 := make([][]int, row)

	rand.Seed(time.Now().Unix())

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			m1[i] = append(m1[i], random(-1, i*j+rand.Intn(10)))
			m2[i] = append(m2[i], random(-1, i*j+rand.Intn(10)))
		}
	}

	fmt.Println("m1", m1)
	fmt.Println("m2", m2)

	r1 := addMatrices(m1, m2)
	r2 := addMatrices(m1, negativeMatrix(m2))
	fmt.Println("r1", r1)
	fmt.Println("r2", r2)

}

func random(min, max int) int {
	return rand.Intn(max-min) + max
}

func negativeMatrix(s [][]int) [][]int {
	for i, x := range s {
		for j := range x {
			s[i][j] = -s[i][j]
		}
	}
	return s
}

func addMatrices(m1, m2 [][]int) [][]int {
	result := make([][]int, len(m1))
	for i, x := range m1 {
		for j := range x {
			result[i] = append(result[i], m1[i][j]+m2[i][j])
		}
	}
	return result
}
