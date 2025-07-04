package main

import (
	"container/ring"
	"fmt"
)

var size = 10

func main() {
	myRing := ring.New(size + 1)
	fmt.Println("Empty ring", *myRing)

	for i := 0; i < myRing.Len()-1; i++ {
		myRing.Value = i
		myRing = myRing.Next()
	}

	myRing.Value = 2

	sum := 0
	myRing.Do(func(x interface{}) {
		t := x.(int)
		sum += t
	})
	fmt.Println("sum:", sum)

	for i := 0; i < myRing.Len()+2; i++ {
		myRing = myRing.Next()
		fmt.Print(myRing.Value, " ")
	}
	fmt.Println()

}
