package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/cch123/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hi %s! This is monkey language\n", user.Name)
	fmt.Println("type some cmds")
	repl.Start(os.Stdin, os.Stdout)
}
