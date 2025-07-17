package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.HelloMessage("Joyston")
	fmt.Println(message)
}
