package main

import "fmt"

func partiallyCovered(a bool, b bool) bool {
	if !a && b {
		return false
	}

	if !b {
		return false
	}

	return true
}

func main() {
	fmt.Printf("Is 42 the answer to everything? %v!\n", partiallyCovered(true, false))
}
