// В этом разделе вы узнаете, как найти идентификатор текущего пользователя,
// а также идентификаторы групп, к которым принадлежит этот пользователь.
// Идентификаторы пользователя и группы являются положительными целыми
// числами, которые хранятся в системных файлах UNIX

package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	fmt.Println("User id:", os.Getuid())
	var u *user.User
	u, _ = user.Current()
	fmt.Print("Group ids: ")
	groupIDs, _ := u.GroupIds()
	for _, i := range groupIDs {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
