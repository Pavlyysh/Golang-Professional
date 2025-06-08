// YAML — это еще один популярный текстовый формат, который используется для
// конфигурационных файлов. В этом подразделе показано, как читать конфигураци-
// онные файлы формата YAML с помощью пакета viper. На этот раз имя конфигу-
// рационного YAML-файла будет передаваться утилите в виде аргумента командной
// строки. Кроме того, в утилите будет использоваться функция viper.AddConfigPath()
// для добавления трех путей поиска, по которым viper будет автоматически ис-
// кать конфигурационные файлы.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func main() {
	var configFile *string = flag.String("c", "myConfig", "Setting the configuration file")
	flag.Parse()

	_, err := os.Stat(*configFile)

	if err == nil {
		fmt.Println("Using User Specified Configuration File")
		viper.SetConfigFile(*configFile)
	} else {
		viper.SetConfigName(*configFile)
		viper.AddConfigPath("/tmp")
		viper.AddConfigPath("$HOME")
		viper.AddConfigPath(".")

		err = viper.ReadInConfig()
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

		if viper.IsSet("item1.k1") {
			fmt.Println("item1.val1:", viper.Get("item1.k1"))
		} else {
			fmt.Println("item1.k1 not set!")
		}
		if viper.IsSet("item1.k2") {
			fmt.Println("item1.val2:", viper.Get("item1.k2"))
		} else {
			fmt.Println("item1.k2 not set!")
		}
	}
	if !viper.IsSet("item3.k1") {
		fmt.Println("item3.k1 is not set!")
	}
}
