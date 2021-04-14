package main

import (
	"bytes"
	"testing"
)

// TestCountwords tests the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	expected := 4
	actual := count(b, false)

	if actual != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, actual)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word4")

	expected := 3
	actual := count(b, true)

	if actual != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, actual)
	}
}
