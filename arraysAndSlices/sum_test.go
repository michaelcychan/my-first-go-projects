package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("adds 1 to 5 and gets 15", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("got %d, but wanted %d, given %v", got, expected, numbers)
		}
	})
	t.Run("adds any number of integers to give results", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		if got != expected {
			t.Errorf("got %d, but wanted %d, given %v", got, expected, numbers)
		}
	})
	t.Run("adds a longer slice numbe of integers to give results", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		got := Sum(numbers)
		expected := 55

		if got != expected {
			t.Errorf("got %d, but wanted %d, given %v", got, expected, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v, but wanted %v.", got, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v, but wanted %v.", got, expected)
		}
	}

	t.Run("normal adding all tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		checkSums(t, got, expected)
	})
	t.Run("empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}

		checkSums(t, got, expected)
	})

}
