package cascade

import (
	"reflect"
	"testing"
)

func TestEachCons(t *testing.T) {
	t.Run("returns slice for step is 1", func(t *testing.T) {
		input := []int{3, 5, 8, 13}
		size := 1
		expected := [][]int{{3}, {5}, {8}, {13}}

		actual := EachCons(input, size)
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})
	t.Run("returns slice for step is 2", func(t *testing.T) {
		input := []int{3, 5, 8, 13}
		size := 2
		expected := [][]int{{3, 5}, {5, 8}, {8, 13}}

		actual := EachCons(input, size)
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected %v, but got %v", expected, actual)
		}
	})
}
