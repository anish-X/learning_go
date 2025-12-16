package main 

import (
	"fmt"
	"day1/greetings"
)

func main() {
	message := greetings.Hello("Anish");
	fmt.Println(message)
}