package foo

import "testing"

func TestBar(t *testing.T) {
	expected := 4
	actual := Bar(5)
	if expected != actual {
		t.Errorf("%v != %v", expected, actual)
	}
}
