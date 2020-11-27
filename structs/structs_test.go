package structs

import "testing"

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}

}

//func TestStructs(t *testing.T) {
//
//	assertCorrectMessage := func(t *testing.T, shape Shape, want float64) {
//		t.Helper()
//		got := shape.Area()
//		if got != want {
//			t.Errorf("got %g want %g", got, want)
//		}
//	}
//
//	assertCorrectMessageForGap := func(t *testing.T, shape Shape, got float64, want float64) {
//		t.Helper()
//		if got != want {
//			t.Errorf("got %g want %g", got, want)
//		}
//	}
//
//	t.Run("Test Area", func(t *testing.T) {
//		rectangle := Rectangle{12, 6}
//		assertCorrectMessage(t, rectangle, 72.0)
//	})
//
//	t.Run("Back the same", func(t *testing.T) {
//		rectangle := Rectangle{12, 6}
//		width, height := rectangle.Gap()
//		assertCorrectMessageForGap(t, rectangle, width, 12.0)
//		assertCorrectMessageForGap(t, rectangle, height, 6.0)
//	})
//
//	t.Run("Test Circles", func(t *testing.T) {
//		circle := Circle{10}
//		assertCorrectMessage(t, circle, 314.1592653589793)
//	})
//
//}
