package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	actual := dictionary.Search("test")
	expected := "this is just a test"

	assertStrings(t, actual, expected)
}

func assertStrings(t *testing.T, actualString, expectedString string) {
	t.Helper()

	if actualString != expectedString {
		t.Errorf("actual %q, expected %q", actualString, expectedString)
	}
}
