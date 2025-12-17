package main

import (
	"day1/greetings"
	"fmt"
	"log"
)

func main() {

	log.SetPrefix("greetings: ") // greetings prefix is added to every log message
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	message, err := greetings.Hellos(names)

	if err != nil {
		log.Println("check") // --> greetings: check
		log.Fatal(err)       // error information
	}

	fmt.Println(message, err)
}
