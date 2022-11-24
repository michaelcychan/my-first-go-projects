package roundUp

import "testing"

func TestRoundToNext5(t *testing.T) {
	assertIntEqual := func(t testing.TB, expected, actual int) {
		t.Helper()
		if expected != actual {
			t.Errorf("expected %d, but got %d", expected, actual)
		}
	}

	t.Run("turns 0 to 0", func(t *testing.T) {
		input := 0
		expected := 0
		actual := RoundToNext5(input)

		assertIntEqual(t, expected, actual)
	})
	t.Run("takes 2 and return 5", func(t *testing.T) {
		input := 2
		expected := 5
		actual := RoundToNext5(input)

		assertIntEqual(t, expected, actual)
	})
	t.Run("takes -5 and return -5", func(t *testing.T) {
		input := -5
		expected := -5
		actual := RoundToNext5(input)

		assertIntEqual(t, expected, actual)
	})
	t.Run("takes -2 and returns 0", func(t *testing.T) {
		input := -2
		expected := 0
		actual := RoundToNext5(input)

		assertIntEqual(t, expected, actual)
	})

	t.Run("run tests in a table", func(t *testing.T) {
		testcases := []struct {
			input    int
			expected int
		}{
			{input: 0, expected: 0},
			{input: 2, expected: 5},
			{input: 5, expected: 5},
			{input: -5, expected: -5},
			{input: -2, expected: 0},
		}

		for _, test := range testcases {
			actual := RoundToNext5(test.input)
			if actual != test.expected {
				t.Errorf("expected %d, but got %d", test.expected, actual)
			}
		}
	})

}
