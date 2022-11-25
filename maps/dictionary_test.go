package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("returns string for a known key", func(t *testing.T) {
		actual, _ := dictionary.Search("test")
		expected := "this is just a test"

		assertStrings(t, actual, expected)
	})
	t.Run("returns a message for a unknown key", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		expected := ErrNotFound

		if err == nil {
			t.Fatal("Expected an error, but did not receive one")
		}

		assertErrors(t, err, expected)
	})
}

func TestAdd(t *testing.T) {
	t.Run("able to add a new key-value pair", func(t *testing.T) {
		dict := Dictionary{}
		dict.Add("anotherKey", "New value")

		expected := "New value"
		actual, err := dict.Search("anotherKey")

		if err != nil {
			t.Fatal("Expected no error, but got one")
		}
		assertStrings(t, actual, expected)
	})
}

func assertStrings(t testing.TB, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected %s, but got %s. Given %s", expected, actual, "test")
	}
}

func assertErrors(t testing.TB, actualError, expectedError error) {
	t.Helper()
	if actualError != expectedError {
		t.Errorf("expected error: %q, but got error: %q", expectedError, actualError)
	}
}
