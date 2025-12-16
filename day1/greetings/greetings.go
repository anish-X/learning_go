package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) { // Hello returns a greeting for the names person (string return type)

	if name == "" {
		return "", errors.New("empty name") // returns two value
	}

	message := fmt.Sprintf("Hi, %v. welcome!", name)
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
