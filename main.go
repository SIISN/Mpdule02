package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var strtoint string = "104"
	var inttostring int = 35

	I, err := strconv.Atoi(strtoint)
	if err != nil {
		log.Fatal(err)
	}
	S := strconv.Itoa(inttostring)
	fmt.Printf("%T", I)
	fmt.Print(" ")
	fmt.Println(I)
	fmt.Printf("%T", S)
	fmt.Print(" " + S)
}
