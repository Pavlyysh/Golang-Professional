package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("This is"))
	fmt.Fprintf(&buffer, " a string!\n")
	buffer.WriteTo(os.Stdout) // Output: This is a string
	buffer.WriteTo(os.Stdout) // ничего не выведет, так как после первого вывода буфер очищается

	buffer.Reset()
	buffer.Write([]byte("Mastering Go!"))
	r := bytes.NewReader(buffer.Bytes())
	fmt.Println(buffer.String())
	for {
		b := make([]byte, 3)
		n, err := r.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Read %s Bytes: %d\n", b, n)
	}
}
