package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Luke")
	fmt.Println(message)
}
