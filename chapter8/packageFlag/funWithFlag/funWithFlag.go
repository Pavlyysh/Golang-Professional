// Утилита funWithFlag.go распознает разные типы
// аргументов, в том числе такой, который принимает несколько значений, раз-
// деленных запятыми. Кроме того, на примере этой программы вы увидите, как
// можно получить доступ к аргументам командной строки, которые расположены
// в конце строки запуска исполняемого файла и не соответствуют ни одному из
// возможных вариантов.

package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

type NamesFlag struct {
	Names []string
}

func (s *NamesFlag) GetNames() []string {
	return s.Names
}

func (s *NamesFlag) String() string {
	return fmt.Sprint(s.Names)
}

func (s *NamesFlag) Set(v string) error {
	if len(s.Names) > 0 {
		return errors.New("cannot use names flag more than once")
	}

	names := strings.Split(v, ",")
	s.Names = append(s.Names, names...)

	return nil
}

func main() {
	var manyNames NamesFlag
	minusK := flag.Int("k", 0, "An int")
	minusO := flag.String("o", "Pasha", "The name")
	flag.Var(&manyNames, "names", "Comma-separated list")

	flag.Parse()
	fmt.Println("-k", *minusK)
	fmt.Println("-o", *minusO)

	for i, item := range manyNames.GetNames() {
		fmt.Println(i, item)
	}

	fmt.Println("Remaining command line arguments:")
	for i, val := range flag.Args() {
		fmt.Println(i, ":", val)
	}
}
