package shape

import "testing"

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, got, expected float64) {
		t.Helper()
		if got != expected {
			t.Errorf("got %.2f, but expected %.2f.", got, expected)
		}
	}

	t.Run("returns 40.0 as perimeters of a rectangle with all sides 10.0", func(t *testing.T) {
		rect := Rectangle{10.0, 10.0}
		got := rect.Perimeter()
		expected := 40.0

		checkPerimeter(t, got, expected)
	})

	t.Run("returns 62.831... as perimeter of a circle with radius of 10.0", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Perimeter()
		expected := 62.83185307179586

		checkPerimeter(t, got, expected)
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

	t.Run("tests different types with same interface using struct slice", func(t *testing.T) {
		areaTests := []struct {
			name     string
			shape    Shape
			expected float64
		}{
			{name: "Rectangle", shape: Rectangle{12.0, 6.0}, expected: 72.0},
			{name: "Circle", shape: Circle{5.0}, expected: 78.53981633974483},
			//{name: "Triangle", shape: Triangle{6.0, 3.0}, expected: 9.0},
			{name: "Triangle-fail", shape: Triangle{6.0, 2.0}, expected: 9.0}, // this will cause failure, uncomment to see the error message
		}

		for _, testCase := range areaTests {
			t.Run(testCase.name, func(t *testing.T) {
				got := testCase.shape.Area()
				if got != testCase.expected {
					t.Errorf("%#v got %g, but expected %g", testCase.name, got, testCase.expected)
				}
			})
		}
	})
}
