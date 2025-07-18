package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func HelloMessage(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}
	// message := fmt.Sprintf("Hello Mr. %v", name)
	message := fmt.Sprintf(randomMsg())

	return message, nil
}

// returns greetings for each person
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	///range through names slice
	for _, name := range names {
		message, err := HelloMessage(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
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
