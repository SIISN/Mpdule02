package main

import (
	"fmt"
)

func main() {
	Week := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Saturday"}
	Wweek := make([]string, 5, 10)
	k := copy(Wweek, Week)
	fmt.Println(k, Wweek)

	Week = Week[len(Wweek):]
	fmt.Println(Week)

	Wweek = append(Wweek, Week...)
	fmt.Println(Wweek)
}
