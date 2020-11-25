package structs

import "testing"

func TestStructs(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	assertCorrectMessageForGap := func(t *testing.T, shape Shape, got float64, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("Test Area", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		assertCorrectMessage(t, rectangle, 72.0)
	})

	t.Run("Back the same", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		width, height := rectangle.Gap()
		assertCorrectMessageForGap(t, rectangle, width, 12.0)
		assertCorrectMessageForGap(t, rectangle, height, 6.0)
	})

	t.Run("Test Circles", func(t *testing.T) {
		circle := Circle{10}
		assertCorrectMessage(t, circle, 314.1592653589793)
	})

}
