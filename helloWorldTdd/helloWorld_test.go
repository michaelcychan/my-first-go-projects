package main

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Run("saying hello world for empty input", func(t *testing.T) {
		got := HelloWorld("")
		want := "Hello World!"
		assertionMessage(t, got, want)
	})
	t.Run("say hello to Mike if input is Mike", func(t *testing.T) {
		got := HelloWorld("Mike")
		want := "Hello Mike!"

		assertionMessage(t, got, want)
	})
	t.Run("Say hello to Tom if input is Tom", func(t *testing.T) {
		got := HelloWorld("Tom")
		want := "Hello Tom!"

		assertionMessage(t, got, want)
	})
}

func assertionMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but got %q", got, want)
	}
}
