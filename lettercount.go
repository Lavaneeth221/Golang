package main

import "fmt"

func main() {
	input := "This is an example sentence."

	// Initialize counters to 0
	wordCount := 0
	letterCount := 0

	// Iterate over each character in the input string
	for i := 0; i < len(input); i++ {
		char := input[i]

		// If the current character is a space or the last character in the string, increment the word count
		if char == ' ' || i == len(input)-1 {
			wordCount++
		}

		// If the current character is a letter or digit, increment the letter count
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			letterCount++
		}
	}

	fmt.Printf("Input: %s\nWord count: %d\nLetter count: %d\n", input, wordCount, letterCount)
}