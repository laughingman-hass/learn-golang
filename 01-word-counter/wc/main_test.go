package main

import (
	"bytes"
	"testing"
)

// TestCountwords tests the count function set to count words
func TestCountwords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	expected := 4
	actual := count(b)

	if actual != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, actual)
	}
}
