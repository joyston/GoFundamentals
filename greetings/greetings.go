package greetings

import "fmt"

func HelloMessage(name string) string {
	message := fmt.Sprintf("Hello Mr. %v", name)
	return message
}
