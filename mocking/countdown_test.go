package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

// it will record how many times the SpySleeper is called in the test
func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := &SpySleeper{}

	Countdown(buffer, sleeper)

	actual := buffer.String()
	expected := `3
2
1
GO!`

	if actual != expected {
		t.Errorf("expected %q, but got %q", expected, actual)
	}
	if sleeper.Calls != 3 {
		t.Errorf("expected to have called sleeper 3 times, but it was called %d times", sleeper.Calls)
	}
}
