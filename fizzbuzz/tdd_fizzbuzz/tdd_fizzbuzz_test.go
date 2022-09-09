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

func Test9(t *testing.T) {
	got := tddFizzBuzz(9)
	want := "Fizz"

	if got != want {
		t.Errorf("got %q, but want %q", got, want)
	}
}

func Test10(t *testing.T) {
	got := tddFizzBuzz(10)
	want := "Buzz"

	if got != want {
		t.Errorf("got %q, but want %q", got, want)
	}
}

func Test13(t *testing.T) {
	got := tddFizzBuzz(13)
	want := "13"

	if got != want {
		t.Errorf("got %q, but want %q", got, want)
	}
}

func Test15(t *testing.T) {
	got := tddFizzBuzz(15)
	want := "FizzBuzz"

	if got != want {
		t.Errorf("got %q, but want %q", got, want)
	}
}
