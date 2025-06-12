// Задача, выполняемая программой pipeline.go, состоит
// в том, чтобы генерировать случайные числа в заданном диапазоне и останавли-
// ваться, когда любое число в этой последовательности встретится во второй раз.
// При этом, прежде чем завершить работу, программа выводит на экран сумму всех
// случайных чисел, сгенерированных до того момента, когда впервые появилось
// повторяющееся случайное число. Для подключения каналов в программе нам
// понадобятся три функции. В этих функциях заключается логика программы. Данные
// будут передаваться по каналам конвейера.

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var CLOSEA = false

var DATA = make(map[int]bool)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Need two integer parameters!")
		return
	}
	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])
	if n1 > n2 {
		fmt.Printf("%d should be smaller than %d\n", n1, n2)
		return
	}

	rand.Seed(time.Now().UnixNano())
	A := make(chan int)
	B := make(chan int)

	go first(n1, n2, A)
	go second(B, A)
	third(B)
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func first(min, max int, out chan<- int) {
	for {
		if CLOSEA {
			close(out)
			return
		}
		out <- random(min, max)
	}
}

func second(out chan<- int, in <-chan int) {
	for x := range in {
		fmt.Print(x, " ")
		_, ok := DATA[x]
		if ok {
			CLOSEA = true
		} else {
			DATA[x] = true
			out <- x
		}
	}
	fmt.Println()
	close(out)
}

func third(in <-chan int) {
	var sum int
	for x2 := range in {
		sum += x2
	}
	fmt.Printf("The sum of the random numbers is %d\n", sum)
}
