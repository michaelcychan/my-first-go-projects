package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("returns string for a known key", func(t *testing.T) {
		actual, err := dictionary.Search("test")
		expected := "this is just a test"

		assertNoError(t, err)
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
		newKey := "anotherKey"
		newValue := "New value"
		dict.Add(newKey, newValue)

		expected := "New value"
		actual, err := dict.Search("anotherKey")

		assertNoError(t, err)
		assertStrings(t, actual, expected)
	})
	t.Run("throws a warning if adding an existing key", func(t *testing.T) {
		key := "key"
		value := "value"
		dict := Dictionary{key: value}
		err := dict.Add(key, "new value")
		assertErrors(t, err, ErrWordExists)

		actual, err := dict.Search(key)
		assertStrings(t, actual, value)
		assertNoError(t, err)
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

func assertNoError(t testing.TB, actualError error) {
	t.Helper()
	if actualError != nil {
		t.Fatal("Expected no error, but got one")
	}
}
