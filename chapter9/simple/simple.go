// В этом подразделе показаны два способа создания горутин. Первый из них заключа-
// ется в использовании обычных функций, а второй — в использовании анонимных
// функций. Эти способы эквивалентны.

package main

import (
	"fmt"
	"time"
)

func main() {
	go function()

	go func() {
		for i := 10; i < 20; i++ {
			fmt.Print(i, " ")
		}
	}()
	time.Sleep(time.Second)
	fmt.Println()
}

func function() {
	for i := range 10 {
		fmt.Print(i)
	}

}
