package main

import "fmt"

const (
    power2 = 2 << iota
    power4
    power8
)

func times10(x int) int { 
    return x * 10 
}

func multiplyByPi(x float64) float64 {
    return x * 3.14
}

func main() {
    fmt.Println(times10(power2))
    fmt.Println(multiplyByPi(float64(power4)))
    fmt.Println(multiplyByPi(float64(power8)))
}
