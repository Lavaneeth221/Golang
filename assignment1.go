package main

import "fmt"

func main() {
	str := "Anil Kumle"
	count := make(map[rune]int)

	for _, c := range str {
		count[c]++
	}

	for c, n := range count {
		fmt.Printf("%c: %d\n", c, n)
	}
}