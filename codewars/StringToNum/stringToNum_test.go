package main

import "testing"

func Test1(t *testing.T) {
	got := stringToNum("1")
	want := 1

	if got != want {
		t.Errorf("got %v  (type: %T), but want %d (type: %T)", got, got, want, want)
	}
}

func Test2(t *testing.T) {
	got := stringToNum("2")
	want := 2

	if got != want {
		t.Errorf("got %v (type: %T), but want %d (type: %T)", got, got, want, want)
	}
}

func Test3(t *testing.T) {
	got := stringToNum("3")
	want := 3

	if got != want {
		t.Errorf("got %v (type: %T), but want %d (type: %T)", got, got, want, want)
	}
}
