package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	actual := Search(dictionary, "test")
	expected := "this is just a test"

	assertStrings(t, actual, expected)
}

func assertStrings(t *testing.T, actualString, expectedString string) {
	t.Helper()

	if actualString != expectedString {
		t.Errorf("actual %q, expected %q", actualString, expectedString)
	}
}
