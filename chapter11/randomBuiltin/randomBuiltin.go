// Сначала с помощью функции rand.New() мы создаем генератор случайных чи-
// сел, а затем используем рефлексию, чтобы получить информацию о типе point3D.
// После этого мы вызываем функцию quick.Value() из пакета testing/quick с дескрип-
// тором типа и генератором случайных чисел, чтобы поместить в переменную myValues
// некоторые случайные данные. Обратите внимание, что для генерации произвольных
// значений структур необходимо экспортировать все поля структуры.

package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing/quick"
	"time"
)

func main() {
	type point3D struct {
		X, Y, Z int8
		S       float32
	}

	ran := rand.New(rand.NewSource(time.Now().Unix()))
	myValues := reflect.TypeOf(point3D{})
	x, _ := quick.Value(myValues, ran)
	fmt.Println(x)
}
