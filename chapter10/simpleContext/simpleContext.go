// Пакет context
//
// 		Главное назначение пакета context — определение типа Context и поддержка ан-
// нулирования. Да, вы все правильно поняли: бывают случаи, когда по какой-либо
// причине необходимо отказаться от того, что вы делаете. Однако очень хорошо,
// когда есть возможность добавить дополнительную информацию о ваших решениях
// об аннулировании. Именно это позволяет сделать пакет context.
//
// 		Если вы посмотрите в исходный код пакета context, то поймете, что его реали-
// зация довольно проста — даже реализация типа Context проста, тем не менее, пакет
// context очень важен.
//
// 		Тип Context — это интерфейс, имеющий четыре метода: Deadline(), Done(),
// Err() и Value(). К счастью, мы не обязаны реализовывать все функции интерфейса
// Context — достаточно лишь изменить переменную Context, используя такие функ-
// ции, как context.WithCancel(), context.WithDeadline() и context.WithTimeout().
// Эти три функции возвращают производный (дочерний) объект Context и функцию
// CancelFunc. Вызов функции CancelFunc удаляет ссылку родителя на дочерний
// объект и останавливает все связанные с ним таймеры. Следовательно, после этого
// сборщик мусора Go может свободно собрать дочерние горутины, у которых больше
// нет связанных с ними родительских горутин. Чтобы сборка мусора работала пра-
// вильно, родительская горутина должна хранить ссылку на все дочерние горутины.
// Если дочерняя горутина завершится без ведома родительской, начинается утечка
// памяти, которая продолжается до тех пор, пока родительская горутина также
// не будет аннулирована.

package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("need a delay")
		return
	}
	delay, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Delay:", delay)

	f1(delay)
	f2(delay)
	f3(delay)
}

func f1(t int) {
	c1 := context.Background()
	c1, cancel := context.WithCancel(c1)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c1.Done():
		fmt.Println("f1():", c1.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f1():", r)
	}
}

func f2(t int) {
	c2 := context.Background()
	c2, cancel := context.WithTimeout(c2, time.Duration(t)*time.Second)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c2.Done():
		fmt.Println("f2():", c2.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f2():", r)
	}
}

func f3(t int) {
	c3 := context.Background()
	deadline := time.Now().Add(time.Duration(2*t) * time.Second)
	c3, cancel := context.WithDeadline(c3, deadline)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c3.Done():
		fmt.Println("f3():", c3.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f3():", r)
	}
}
