package main

import "fmt"

func main() {
	x := 7.0
	fmt.Printf("Now you have %g problems.\n", mySqrt(x))
}

func mySqrt(x float64) float64 {
	z := x / 2.0
	for i := 0; i < 10; i++ {
		z = z - ((z*z - x) / (2 * z))
	}
	return z
}
