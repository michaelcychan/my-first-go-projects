package main

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Run("saying hello world for empty input", func(t *testing.T) {
		got := HelloWorld("", "en")
		want := "Hello World!"
		assertionMessage(t, got, want)
	})
	t.Run("say hello to Mike if input is Mike", func(t *testing.T) {
		got := HelloWorld("Mike", "en")
		want := "Hello Mike!"

		assertionMessage(t, got, want)
	})
	t.Run("Say hello to Tom if input is Tom", func(t *testing.T) {
		got := HelloWorld("Tom", "en")
		want := "Hello Tom!"

		assertionMessage(t, got, want)
	})
	t.Run("Say こにちは to ゆき if lang is japanese", func(t *testing.T) {
		got := HelloWorld("ゆき", "jp")
		want := "こにちは、ゆき!"

		assertionMessage(t, got, want)
	})
	t.Run("Say 你好 to 偉仔 if lang is zh", func(t *testing.T) {
		got := HelloWorld("偉仔", "zh")
		want := "你好，偉仔!"

		assertionMessage(t, got, want)
	})
}

func assertionMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
