package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func multiplyMatrix(m1, m2 [][]int) ([][]int, error) {
	if len(m1[0]) != len(m2) {
		return nil, errors.New("cannot multiply given matrixes")
	}

	result := make([][]int, len(m1))

	for i := 0; i < len(m1); i++ {
		result[i] = make([]int, len(m2[0]))
		for j := 0; j < len(m2[0]); j++ {
			for k := 0; k < len(m2); k++ {
				result[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}

	return result, nil
}

func createMatrix(row, col int) [][]int {
	r := make([][]int, row)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			r[i] = append(r[i], random(-5, i*j))
		}
	}
	return r
}
func main() {
	rand.Seed(time.Now().Unix())
	arguments := os.Args
	if len(arguments) != 5 {
		fmt.Println("Wrong number of arguments!")
		return
	}
	var row, col int
	row, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("Need an integer: ", arguments[1])
		return
	}
	col, err = strconv.Atoi(arguments[2])
	if err != nil {
		fmt.Println("Need an integer: ", arguments[2])
		return
	}
	if col <= 0 || row <= 0 {
		fmt.Println("Need positive matrix dimensions!")
		return
	}
	fmt.Printf("m1 is a %dx%d matrix\n", row, col)
	// Инициализация m1 случайными числами
	m1 := createMatrix(row, col)
	row, err = strconv.Atoi(arguments[3])
	if err != nil {
		fmt.Println("Need an integer: ", arguments[3])
		return
	}
	col, err = strconv.Atoi(arguments[4])
	if err != nil {
		fmt.Println("Need an integer: ", arguments[4])
		return
	}
	if col <= 0 || row <= 0 {
		fmt.Println("Need positive matrix dimensions!")
		return
	}
	fmt.Printf("m2 is a %dx%d matrix\n", row, col)
	// Инициализация m2 случайными числами
	m2 := createMatrix(row, col)
	fmt.Println("m1:", m1)
	fmt.Println("m2:", m2)
	// Умножение
	r1, err := multiplyMatrix(m1, m2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("r1:", r1)

}
