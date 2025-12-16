package greetings

import "fmt"

func Hello(name string) string { // Hello returns a greeting for the names person (string return type)
	message := fmt.Sprintf("Hi, welcome!", name)  /*
																										:= operator is shortcut for 
																											{
																												var message string
																												message = fmt. ......
																											}
																									  := declare and initialize and uses the value on right 
																										to determine the variable's type

																											*/
	return message;
}