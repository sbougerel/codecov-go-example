package main

import (
	"fmt"
	"time"
)

func partiallyCovered(a bool, b bool, c bool) bool {
	if !a {
		return false
	}

	if !b {
		return false
	}

	if !c {
		return false
	}

	return true
}

func main() {
	time.Sleep(1 * time.Millisecond)
	fmt.Printf("Is 42 the answer to everything? %v!\n", partiallyCovered(true, false, false))
	return
}
