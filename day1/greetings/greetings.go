package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) { // Hello returns a greeting for the names person (string return type)

	if name == "" {
		return "", errors.New("empty name") // returns two value
	}

	message := fmt.Sprintf(randomFormat(), name)
	/*
			:= operator is shortcut for
				{
					var message string
					message = fmt. ......
				}
		  := declare and initialize and uses the value on right
			to determine the variable's type

	*/
	return message, nil // returns two value
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	// a slice of message formats
	formats := []string{
		"Him %v. Welcome", "Great to see you, %v!", "Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
