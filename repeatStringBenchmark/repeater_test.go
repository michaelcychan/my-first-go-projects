package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeater() {
	actual := Repeater("abc")
	fmt.Println(actual)
	// Output: abcabcabcabcabc
}

func TestRepeater(t *testing.T) {
	t.Run("repeats a for five times", func(t *testing.T) {
		actual := Repeater("a")
		expected := "aaaaa"
		assertionMessage(t, expected, actual)
	})
	t.Run("repeats b for five times", func(t *testing.T) {
		actual := Repeater("b")
		expected := "bbbbb"
		assertionMessage(t, expected, actual)
	})
}

func assertionMessage(t testing.TB, expected, actual string) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

func BenchmarkRepeater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeater("hjguyguiyfiyuotiuytrbityg")
	}
}
