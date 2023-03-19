package main

import "fmt"

func divide(dividend, divisor int) (quotient, remainder int) {
	quotient, remainder = div(dividend, divisor)
	return
}

func main() {
	fmt.Println(divide(17, 4))
}

func div(dividend, divisor int) (quotient, remainder int) {
    quotient = dividend / divisor
    remainder = dividend % divisor
    return quotient, remainder
}
