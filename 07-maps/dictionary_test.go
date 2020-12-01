package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	actual := Search(dictionary, "test")
	expected := "this is just a test"

	if actual != expected {
		t.Errorf("actual %q, expected %q, given %q", actual, expected, "test")
	}
}
