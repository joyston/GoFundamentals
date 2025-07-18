package greetings

import (
	"errors"
	"fmt"
)

func HelloMessage(name string) (string, error) {

	if name == "" {
		return "", errors.New("Empty name")
	}
	message := fmt.Sprintf("Hello Mr. %v", name)
	return message, nil
}
