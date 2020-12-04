package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	actual := buffer.String()
	expected := "3\n2\n1\nGo!"

	if actual != expected {
		t.Errorf("actual %q, expected %q", actual, expected)
	}
}
