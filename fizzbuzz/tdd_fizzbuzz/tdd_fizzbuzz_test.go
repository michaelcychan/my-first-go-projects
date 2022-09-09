package main

import "testing"

func Test1(t *testing.T) {
	got := tddFizzBuzz(1)
	want := "1"

	if got != want {
		t.Errorf("got %q, but want %q", got, want)
	}
}

func Test2(t *testing.T) {
	got := tddFizzBuzz(2)
	want := "2"

	if got != want {
		t.Errorf("got %q, but want %q", got, want)
	}
}

func Test3(t *testing.T) {
	got := tddFizzBuzz(3)
	want := "Fizz"

	if got != want {
		t.Errorf("got %q, but want %q", got, want)
	}
}

func Test5(t *testing.T) {
	got := tddFizzBuzz(5)
	want := "Buzz"

	if got != want {
		t.Errorf("got %q, but want %q", got, want)
	}
}

func Test6(t *testing.T) {
	got := tddFizzBuzz(6)
	want := "Fizz"

	if got != want {
		t.Errorf("got %q, but want %q", got, want)
	}
}
