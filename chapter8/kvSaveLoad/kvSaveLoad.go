// При каждом запуске kvSaveLoad.go программа проверяет, есть ли данные для
// чтения, пытаясь прочитать файл данных, предлагаемый по умолчанию. Если файла
// данных для чтения нет, то программа начнет работу с пустого хранилища «ключ —
// значение». Перед завершением работы программа записывает все данные на диск
// с помощью функции save().

package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"os"
	"strings"
)

type myElement struct {
	Name    string
	Surname string
	Id      string
}

var DATAFILE = "data.gob"
var DATA = make(map[string]myElement)

func main() {
	err := load()
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		tokens := strings.Fields(text)

		switch len(tokens) {
		case 0:
			continue
		case 1:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 2:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 3:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 4:
			tokens = append(tokens, "")
		}

		switch tokens[0] {
		case "PRINT":
			PRINT()
		case "STOP":
			err = save()
			if err != nil {
				fmt.Println(err)
			}
			return
		case "DELETE":
			if !DELETE(tokens[1]) {
				fmt.Println("Delete operation failed!")
			}
		case "ADD":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !ADD(tokens[1], n) {
				fmt.Println("Add operation failed!")
			}
		case "LOOKUP":
			n := LOOKUP(tokens[1])
			if n != nil {
				fmt.Printf("%v\n", *n)
			}
		case "CHANGE":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !CHANGE(tokens[1], n) {
				fmt.Println("Update operation failed!")
			}
		default:
			fmt.Println("Unknown command - please try again!")
		}
	}
}

func PRINT() {
	for k, v := range DATA {
		fmt.Printf("key: %s, value: %v\n", k, v)
	}
}

func CHANGE(k string, n myElement) bool {
	DATA[k] = n
	return true
}

func DELETE(k string) bool {
	if LOOKUP(k) != nil {
		delete(DATA, k)
		return true
	}
	return false
}

func LOOKUP(k string) *myElement {
	_, ok := DATA[k]
	if ok {
		n := DATA[k]
		return &n
	}
	return nil
}

func ADD(k string, new myElement) bool {
	if k == "" {
		return false
	}

	if LOOKUP(k) == nil {
		DATA[k] = new
		return true
	}
	return false
}

// Обратите внимание: функция save() прежде всего удаляет существующий файл
// данных с помощью функции os.Remove(), чтобы позже создать новый.
// Одна из самых важных задач, которые выполняет функция save(), — убедиться,
// что мы действительно можем создать и записать нужный файл. Есть много способов
// сделать это, и в save() используется самый простой из них: проверка значения
// error, возвращаемого функцией os.Create(). Если это значение не равно nil, это
// указывает на проблему и функция save() завершается без сохранения каких-либо
// данных.
func save() error {
	fmt.Println("Saving", DATAFILE)

	err := os.Remove(DATAFILE)
	if err != nil {
		fmt.Println(err)
	}

	saveTo, err := os.Create(DATAFILE)
	if err != nil {
		fmt.Println("cannot create", DATAFILE)
		return err
	}
	defer saveTo.Close()

	encoder := gob.NewEncoder(saveTo)
	err = encoder.Encode(DATA)
	if err != nil {
		fmt.Println("cannot save to", DATAFILE)
		return err
	}

	return nil
}

// Одна из задач функции load() — убедиться, что файл, который мы пытаемся
// прочитать, действительно существует и его без проблем можно прочитать.
// В функции load() также используется самый простой подход, который заключается
// в просмотре значения, возвращаемого функцией os.Open(). Если возвращаемое
// значение error равно nil, то все в порядке.
// Важно также закрыть файл после чтения данных, поскольку впоследствии он
// будет перезаписан функцией save(). Освобождение файла выполняется с помощью
// инструкции defer loadFrom.Close().
func load() error {
	fmt.Println("Loading", DATAFILE)
	loadFrom, err := os.Open(DATAFILE)
	if err != nil {
		fmt.Println("Empty key/value store")
		return err
	}
	defer loadFrom.Close()

	decoder := gob.NewDecoder(loadFrom)
	decoder.Decode(&DATA)
	return nil
}
