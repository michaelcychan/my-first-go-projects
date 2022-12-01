package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

// it will record how many times the SpySleeper is called in the test
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

// Being section to test the sequence of actions
type SpyCountdownOperations struct {
	Calls []string
}

const sleep = "sleep"
const write = "write"

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("should print and sleep with correct number of counts", func(t *testing.T) {
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
	})
	t.Run("should print and sleep alternatively", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}

		Countdown(spySleepPrinter, spySleepPrinter)
		expected := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(expected, spySleepPrinter.Calls) {
			t.Errorf("expected, %v, but got %v", expected, spySleepPrinter.Calls)
		}
	})

}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurationSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("expected to have slept for %v, but actually slept for %v", sleepTime, spyTime.durationSlept)
	}
}
