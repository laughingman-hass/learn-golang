package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Hassan")

	actual := buffer.String()
	expected := "Hello, Hassan"

	if actual != expected {
		t.Errorf("actual %q, expected %q", actual, expected)
	}
}
