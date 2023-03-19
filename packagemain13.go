package main

import (
	"fmt"
)

func main() {
	x, y := 3, 4
	z := uint((x*x + y*y)^(1/2))
	fmt.Println(x, y, z)
}