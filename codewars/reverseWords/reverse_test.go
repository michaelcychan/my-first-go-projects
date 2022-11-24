package reverse

import "testing"

func TestReverseWords(t *testing.T) {
	t.Run("reverse a single word", func(t *testing.T) {
		input := "apple"
		expected := "elppa"

		actual := ReverseWords(input)
		assertStrings(t, expected, actual)
	})
	t.Run("reverse a sentence with all single char word", func(t *testing.T) {
		input := "a b c d"
		expected := "a b c d"

		actual := ReverseWords(input)
		assertStrings(t, expected, actual)
	})
	t.Run("test criteria sentence", func(t *testing.T) {
		input := "The quick brown fox jumps over the lazy dog."
		expected := "ehT kciuq nworb xof spmuj revo eht yzal .god"

		actual := ReverseWords(input)
		assertStrings(t, expected, actual)
	})
}

func assertStrings(t testing.TB, expected, actual string) {
	if actual != expected {
		t.Errorf("expected %s, but got %s", expected, actual)
	}
}
