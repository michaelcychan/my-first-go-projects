package calculator

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleSum() {
	actual := Sum([]int{1, 2, 3})
	fmt.Println(actual)
	// Output: 6
}

func TestSum(t *testing.T) {

	t.Run("returns 21 as sum of an int slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		got := Sum(numbers)
		expected := 21

		assertMessage(t, got, expected, numbers)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("takes a number of slice and return a slice of sum of each slice", func(t *testing.T) {
		slice1 := []int{1, 2}
		slice2 := []int{0, 9}
		got := SumAll(slice1, slice2)
		expected := []int{3, 9}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v, but expected %v", got, expected)
		}

	})
}

func assertMessage(t testing.TB, got, expected int, numbers []int) {
	t.Helper()
	if got != expected {
		t.Errorf("got %d but expected %d. Input was %v", got, expected, numbers)
	}
}
