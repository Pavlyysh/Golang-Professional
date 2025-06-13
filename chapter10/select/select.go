// Оператор select в Go похож на switch, но только для каналов.
// На практике это означает, что select позволяет горутине
// дождаться завершения нескольких операций коммуникации. Поэтому основное
// преимущество select состоит в том, что этот оператор дает возможность работать
// с несколькими каналами, используя один блок select. В результате, построив
// соответствующие блоки select, можно выполнять неблокирующие операции
// с каналами.
// Самая большая проблема при использовании нескольких каналов и ключевого
// слова select — это взаимоблокировки. Это означает, что необходимо
// быть особенно осторожными при проектировании и разработке,
// чтобы избежать таких взаимоблокировок

// Главным преимуществом оператора select является то, что он может
// подключать, распределять нагрузку и управлять несколькими каналами.
// Поскольку каналы служат для связи между горутинами, select соединяет
// каналы, которые соединяют горутины. Таким образом, оператор select
// является одной из наиболее важных, если не самой важной, частью
// модели конкурентности в Go.

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	createNumber := make(chan int)
	end := make(chan bool)
	if len(os.Args) != 2 {
		fmt.Println("give an integer")
		return
	}
	n, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Creating %d random nums\n", n)
	go gen(0, 2*n, createNumber, end)

	for i := 0; i < n; i++ {
		fmt.Printf("Number: %d\n", <-createNumber)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Exiting...")
	end <- true
}

func gen(min, max int, createNumber chan int, end chan bool) {
	for {
		select {
		case createNumber <- rand.Intn(max-min) + min:
		case <-end:
			close(end)
			return
		case <-time.After(4 * time.Second):
			fmt.Println("\ntime.After()!.")
		}
	}
}
