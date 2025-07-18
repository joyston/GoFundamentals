package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func HelloMessage(name string) (string, error) {

	if name == "" {
		return "", errors.New("Empty name")
	}
	// message := fmt.Sprintf("Hello Mr. %v", name)
	message := fmt.Sprintf(randomMsg(), name)

	return message, nil
}

// FUnction to send random greeting
func randomMsg() string {
	msg := []string{
		"Howdy %v",
		"Templar %v, nice meeting you!",
		"Deus Vult Mr. %v",
	}

	return msg[rand.Intn(len(msg))]
}
