package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("returns 3 when adds 1 and 2", func(t *testing.T) {
		sum := Adder(1, 2)
		expected := 3

		assertionMessage(t, sum, expected)
	})
	t.Run("returns 4 when adds 2 to 2", func(t *testing.T) {
		sum := Adder(2, 2)
		expected := 4
		assertionMessage(t, sum, expected)
	})
	t.Run("returns 5 when adds 2 to 3", func(t *testing.T) {
		sum := Adder(2, 3)
		expected := 5
		assertionMessage(t, sum, expected)
	})
}

func assertionMessage(t testing.TB, sum, expected int) {
	t.Helper()
	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}
