package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.HelloMessage("Joyston")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"Mark", "Luke", "John", "Mathew"}

	messages, msgErr := greetings.Hellos(names)
	if msgErr != nil {
		log.Fatal(msgErr)
	}

	fmt.Println(messages)
}
