package main

import (
	"fmt"
	"math"
)

func main() {
	R := new(float64)
	var S float64
	const (
		L = 35
		p = 3.14
	)
	*R = L / (2 * p)
	*R = math.Round(*R*100) / 100
	fmt.Println(*R)

	S = p * math.Pow(*R, 2)
	fmt.Println(math.Round(S*100) / 100)
}
