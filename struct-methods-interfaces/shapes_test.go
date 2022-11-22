package shape

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("returns 40.0 as perimeters of a rectangle with all sides 10.0", func(t *testing.T) {
		rect := Rectangle{10.0, 10.0}
		got := Perimeter(rect)
		expected := 40.0

		if got != expected {
			t.Errorf("got %.2f, but expected %.2f.", got, expected)
		}
	})
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		got := shape.Area()

		if got != expected {
			t.Errorf("got %.2f, but expected %.2f.", got, expected)
		}
	}

	t.Run("returns 100.0 as area of a rectangle with all sides as 10.0", func(t *testing.T) {
		rect := Rectangle{10.0, 10.0}
		checkArea(t, rect, 100.0)
	})
	t.Run("return 314.159... for an area of circle with radius as 10.0", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}
