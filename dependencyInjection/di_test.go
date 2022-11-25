package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("makes a greeting", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Chris")

		actual := buffer.String()
		expected := "Hello, Chris"

		if actual != expected {
			t.Errorf("expected, %q, but got %q", expected, actual)
		}
	})
}
