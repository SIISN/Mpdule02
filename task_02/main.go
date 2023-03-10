package main

import "fmt"

func main() {
	A := new(int)
	B := 1
	A = &B
	fmt.Println(*A)
	*A = 2
	fmt.Println(B)
}
