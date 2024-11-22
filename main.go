package main

import (
	"fmt"
	"log"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Println("Feel free to type some Monkey code.")
	repl.Start(os.Stdin, os.Stdout)
}
