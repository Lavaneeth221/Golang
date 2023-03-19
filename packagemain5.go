package main

import "fmt"

func multiplyAndSubtract(x, y, z int) int {
	return z - (x * y)
}

func main() {
	fmt.Println(multiplyAndSubtract(2, 3, 10))
}
