package structs

import "testing"

func TestPerimeter(t *testing.T) {
	ErrorMessage := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("[FAILED]\ngot => %.2f\nwant => %.2f", got, want)
		}
	}

	t.Run("Calculate the perimeter of a rectagle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0

		ErrorMessage(t, got, want)
	})
}

func TestArea(t *testing.T) {
	ErrorMessage := func(t *testing.T, shape Shape, got float64, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("[FAILED]\n%#v\ngot => %g\nwant => %g", shape, got, want)
		}
	}

	t.Run("Test all the allowed shapes", func(t *testing.T) {
		areaTests := []struct {
			shape Shape
			want  float64
		}{
			{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
			{shape: Circle{Radius: 10}, want: 314.1592653589793},
			{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
		}

		for _, tt := range areaTests {
			got := tt.shape.Area()
			ErrorMessage(t, tt.shape, got, tt.want)
		}
	})
}
