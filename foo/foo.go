package foo

// Exported
func Bar(a int) int {
	a = 0
	for i := 0; i < a; i++ {
		a = i
	}
	return a
}
