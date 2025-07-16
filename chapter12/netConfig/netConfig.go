package main

import (
	"fmt"
	"net"
)

func main() {
	// Функция net.Interfaces() возвращает все интерфейсы текущего компьютера
	// в виде среза, содержащего элементы типа net.Interface. Этот срез мы будем
	// использовать для получения желаемой информации.
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, i := range interfaces {
		fmt.Printf("Interface: %v\n", i.Name)
		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			fmt.Println(err)
		}

		addresses, err := byName.Addrs()
		for k, v := range addresses {
			fmt.Printf("Interface Address #%v: %v\n", k, v.String())
		}
		fmt.Println()
	}
}
