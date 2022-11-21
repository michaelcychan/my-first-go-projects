package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	output := Repeat("ABC", 3)
	fmt.Println(output)
	// Output:
	// ABCABCABC
}

func TestRepeat(t *testing.T) {
	t.Run("repeats a for 5 times", func(t *testing.T) {
		actual := Repeat("a", 5)
		expected := "aaaaa"
		if actual != expected {
			t.Errorf("expected %q, but got %q", expected, actual)
		}
	})
	t.Run("repeats a for 10 times", func(t *testing.T) {
		actual := Repeat("a", 10)
		expected := "aaaaaaaaaa"
		if actual != expected {
			t.Errorf("expected %q, but got %q", expected, actual)
		}
	})
}

func TestCountNumberOf(t *testing.T) {
	t.Run("returns the number of a in a sentence as 1", func(t *testing.T) {
		inputSentence := "This is a test"
		actual := CountNumberOf("a", inputSentence)
		expected := 1
		if actual != expected {
			t.Errorf("expected %d, but got %d", expected, actual)
		}
	})
	t.Run("returns the number of i in a sentence as 2", func(t *testing.T) {
		inputSentence := "This is a test"
		actual := CountNumberOf("i", inputSentence)
		expected := 2
		if actual != expected {
			t.Errorf("expected %d, but got %d", expected, actual)
		}
	})
	t.Run("returns the number of is in a sentence as 2", func(t *testing.T) {
		inputSentence := "This is a test"
		actual := CountNumberOf("is", inputSentence)
		expected := 2
		if actual != expected {
			t.Errorf("expected %d, but got %d", expected, actual)
		}
	})
}

func BenchmarkRepate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
