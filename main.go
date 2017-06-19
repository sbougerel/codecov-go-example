package main

import "fmt"

func partiallyCovered(a bool, b bool) bool {
	if !a && b {
		// This is a comment, and below an intentional empty line

		return false
	}

	if !b {
		// This is a comment, and below an intentional empty line

		return false
	}

	return true
}

func main() {
	fmt.Printf("Is 42 the answer to everything? %v!\n", partiallyCovered(true, false))
}
