package reflection

import "testing"

func TestWalk(t *testing.T) {
	expected := "LaughingMan"
	var actual []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		actual = append(actual, input)
	})

	if len(actual) != 1 {
		t.Errorf("wrong number of function calls, actual %d, expected %d", len(actual), 1)
	}

	if actual[0] != expected {
		t.Errorf("actual %q, expected %q", actual[0], expected)
	}
}
