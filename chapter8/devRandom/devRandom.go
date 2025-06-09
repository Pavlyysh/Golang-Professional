// В этом подразделе показано, как читать данные с устройства /dev/random. Назначе-
// ние системного устройства /dev/random — генерировать случайные данные, кото-
// рые можно использовать для тестирования программ или, как в данной ситуации,
// в качестве начального числа для генератора случайных чисел.

package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/dev/random")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	var seed int64
	binary.Read(f, binary.LittleEndian, &seed)
	fmt.Println("seed:", seed)
}
