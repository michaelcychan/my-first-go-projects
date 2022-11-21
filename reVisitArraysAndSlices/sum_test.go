package calculator

import "testing"

func TestSum(t *testing.T) {

	t.Run("returns 21 as sum of an int slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		got := Sum(numbers)
		expected := 21

		assertMessage(t, got, expected, numbers)
	})
}

func assertMessage(t testing.TB, got, expected int, numbers []int) {
	t.Helper()
	if got != expected {
		t.Errorf("got %d but expected %d. Input was %v", got, expected, numbers)
	}
}
