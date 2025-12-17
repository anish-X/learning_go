package main

import (
	"fmt"
)

// go maps doesn't return error on not finding the key instead sends 0/""/false and false
var timezone = map[string]int{
	"UTC": 0 * 60 * 60,
	"EST": -5 * 60 * 60,
	"CST": -6 * 60 * 60,
	"MST": -7 * 60 * 60,
	"PST": -8 * 60 * 60,
}

/*
if the key is in the map, then seconds gets value and ok get true
if the key is not in the map, then seconds gets 0 and ok gets false
*/
func offset(tz string) (int, bool) {
	var seconds int
	var ok bool
	/*
		seconds, ok := timezone[tz]
		if ok {
		return seconds, ok
		}
	*/
	if seconds, ok = timezone[tz]; ok { //
		return seconds, ok
	}
	fmt.Println("unknown time zone:", tz)
	return seconds, ok
}

func main() {

	fmt.Println(offset("EST"))
	fmt.Println(offset("ESdT"))

}
