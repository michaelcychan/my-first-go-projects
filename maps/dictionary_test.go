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
		expected := "no result from your keyword"

		if err == nil {
			t.Fatal("Expected an error, but did not receive one")
		}

		assertStrings(t, err.Error(), expected)
	})

}

func assertStrings(t testing.TB, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected %s, but got %s. Given %s", expected, actual, "test")
	}
}
