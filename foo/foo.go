package foo

// Exported
func Bar(a int) int {
	b := 0
	for i := 0; i < a; i++ {
		b = i
	}
	return b
}
