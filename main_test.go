package main

import "testing"

func TestPartiallyCovered(t *testing.T) {
	t.Parallel()
	var expected = false
	var actual = partiallyCovered(true, false)
	if expected != actual {
		t.Errorf("%v != %v\n", expected, actual)
	}
}
