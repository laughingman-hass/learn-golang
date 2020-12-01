package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		actual, _ := dictionary.Search("test")
		expected := "this is just a test"

		assertStrings(t, actual, expected)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, actual := dictionary.Search("unknown")
		expected := ErrNotFound

		assertError(t, actual, expected)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	expected := "this is just a test"
	actual, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word", err)
	}

	if actual != expected {
		t.Errorf("actual %q, expected %q", actual, expected)
	}
}

func assertStrings(t *testing.T, actualString, expectedString string) {
	t.Helper()

	if actualString != expectedString {
		t.Errorf("actual %q, expected %q", actualString, expectedString)
	}
}

func assertError(t *testing.T, actualError, expectedError error) {
	t.Helper()

	if actualError != expectedError {
		t.Errorf("actual error %q expected %q", actualError, expectedError)
	}
}
