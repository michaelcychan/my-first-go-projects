package calculator

import (
	"fmt"
	"reflect"
	"testing"
)

// reflect.DeepEqual would be useful to compare slices or arrays

func ExampleSum() {
	actual := Sum([]int{1, 2, 3})
	fmt.Println(actual)
	// Output: 6
}

func TestSum(t *testing.T) {

	checkSum := func(t testing.TB, got, expected int) {
		t.Helper()
		if got != expected {
			t.Errorf("got %d but expected %d.", got, expected)
		}
	}

	t.Run("returns 21 as sum of an int slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		got := Sum(numbers)
		expected := 21

		checkSum(t, got, expected)
	})
}

func ExampleSumAll() {
	actual := SumAll([]int{1, 2, 3}, []int{3, 4, 5}, []int{6, 7, 8})
	fmt.Println(actual)
	// Output: [6 12 21]
}

func TestSumAll(t *testing.T) {

	checkSums := func(t testing.TB, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v, but expected %v", got, expected)
		}
	}

	t.Run("takes a number of slice and return a slice of sum of each slice", func(t *testing.T) {
		slice1 := []int{1, 2}
		slice2 := []int{0, 9}
		got := SumAll(slice1, slice2)
		expected := []int{3, 9}

		checkSums(t, got, expected)
	})

	t.Run("returns sums safely with empty input slice", func(t *testing.T) {
		slice1 := []int{}
		slice2 := []int{1, 2, 9}
		got := SumAll(slice1, slice2)
		expected := []int{0, 12}

		checkSums(t, got, expected)
	})
}

func ExampleSumAllTails() {
	actual := SumAllTails([]int{1, 2, 3}, []int{3, 4, 5}, []int{6, 7, 8})
	fmt.Println(actual)
	// Output: [5 9 15]
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v, but expected %v", got, expected)
		}
	}

	t.Run("returns sums of each slice except the first", func(t *testing.T) {
		slice1 := []int{1, 2, 3}
		slice2 := []int{0, 9}
		got := SumAllTails(slice1, slice2)
		expected := []int{5, 9}

		checkSums(t, got, expected)
	})

	t.Run("returns sums of each slices even for empty slice", func(t *testing.T) {
		slice1 := []int{}
		slice2 := []int{0, 9}
		got := SumAllTails(slice1, slice2)
		expected := []int{0, 9}

		checkSums(t, got, expected)
	})
}
