// Запись в файл
// Для записи данных в файлы на диске можно использовать функциональность
// интерфейса io.Writer. Тем не менее, на примере кода Go из программы save.go
// рассмотрим пять разных способов записи данных в файл.

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	s := []byte("Data to write\n")

	// 1
	// Обратите внимание, что мы будем использовать байтовый срез s в каждой строке
	// этой программы Go, где производится запись. Кроме того, применяемая здесь
	// функция fmt.Fprintf() позволяет записывать данные в ваши собственные файлы
	// журнала, используя нужный формат. В нашем случае fmt.Fprintf() записывает
	// данные в файл, обозначенный как f1
	f1, err := os.Create("f1.txt")
	if err != nil {
		fmt.Println("error create file", err)
		return
	}
	defer f1.Close()
	fmt.Fprintf(f1, string(s))

	// 2
	// В данном случае для записи данных в файл используется функция f2.Wri­teString()
	f2, err := os.Create("f2.txt")
	if err != nil {
		fmt.Println("error create file", err)
		return
	}
	defer f2.Close()

	n, err := f2.WriteString(string(s))
	if err != nil {
		fmt.Println("error write to file", err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)

	// 3
	// В данном случае функция bufio.NewWriter() открывает файл для записи, а bufio.WriteString() записывает данные
	f3, err := os.Create("f3.txt")
	if err != nil {
		fmt.Println("error create file", err)
		return
	}
	defer f3.Close()

	w := bufio.NewWriter(f3)
	n, err = w.WriteString(string(s))
	if err != nil {
		fmt.Println("error write to file", err)
		return
	}
	fmt.Printf("write %d bytes\n", n)
	w.Flush()

	// 4
	// Этот метод требует только одного вызова функции ioutil.WriteFile() для записи данных, и не требует использования os.Create()
	f4 := "f4.txt"
	err = ioutil.WriteFile(f4, s, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 5
	// В этом методе для записи данных в файл использована функция io.Wri­teString()
	f5, err := os.Create("f5.txt")
	if err != nil {
		fmt.Println("error create file", err)
		return
	}
	defer f5.Close()

	n, err = io.WriteString(f5, string(s))
	if err != nil {
		fmt.Println("error write to file", err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)
}
