package bus

import "testing"

func TestNumber(t *testing.T) {
	t.Run("one stop with no in and out", func(t *testing.T) {
		input := [][2]int{{0, 0}}
		expected := 0

		actual := Number(input)
		assertNumberInBus(t, expected, actual, input)
	})
	t.Run("returns results for three stops", func(t *testing.T) {
		input := [][2]int{{10, 0}, {3, 5}, {5, 8}}
		expected := 5

		actual := Number(input)
		assertNumberInBus(t, expected, actual, input)
	})
	t.Run("returns results for exampl case", func(t *testing.T) {
		input := [][2]int{{3, 0}, {9, 1}, {4, 10}, {12, 2}, {6, 1}, {7, 10}}
		expected := 17

		actual := Number(input)
		assertNumberInBus(t, expected, actual, input)
	})
}

func assertNumberInBus(t testing.TB, expected, actual int, input [][2]int) {
	if actual != expected {
		t.Errorf("expcted %d, but got %d. Input was %v", expected, actual, input)
	}
}
