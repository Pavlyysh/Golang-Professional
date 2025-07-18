// Возможно, у вас уже есть программа Go, в которой используется пакет flag и ко-
// торую вы хотели бы перевести на пакет viper. Например, рассмотрим следующий
// код Go, где используется пакет flag:
// package main
//
// import (
// 	   "flag"
//     "fmt"
// )
//
// func main() {
// 		minusI := flag.Int("i", 100, "i parameter")
// 		flag.Parse()
// 		i := *minusI
// 		fmt.Println(i)
// }

package main

import (
	"flag"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	flag.Int("i", 100, "i parameter")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)
	i := viper.GetInt("i")
	fmt.Println(i)
}
