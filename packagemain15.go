package main

import "fmt"

const (
    greeting = "Hello"
    age = 21
)

func main() {
    const country = "India"
    fmt.Println(greeting, "from", country)
    fmt.Println("I am", age, "years old")
    
    const married = false
    fmt.Println("Am I married?", married)
}
