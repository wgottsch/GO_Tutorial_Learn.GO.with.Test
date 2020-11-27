package structs

import "testing"

func TestComputing(t *testing.T) {

	assertCorrectMessageArea := func(t *testing.T, computing Compute, want float64) {
		t.Helper()
		got := computing.Flaeche()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	assertCorrectMessagePerimeter := func(t *testing.T, computing Compute, want float64) {
		t.Helper()
		got := computing.Umfang()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	//assertCorrectMessageHochUndBreit := func(t *testing.T, computing Compute, want []float64) {
	assertCorrectMessageHochUndBreit := func(t *testing.T, computing Compute, w1 float64, w2 float64) {
		//var got []float64
		t.Helper()
		g1, g2 := computing.HochUndBreit()
		if g1 != w1 && g2 != w2 {
			t.Errorf("got %g %g want %g %g", g1, g2, w1, w2)
		}
	}

	// Fläche
	t.Run("Test Fläche des Rechtecks", func(t *testing.T) {
		rechteck := Rechteck{12, 6}
		assertCorrectMessageArea(t, rechteck, 72.0)
	})

	t.Run("Test Fläche des Kreises", func(t *testing.T) {
		kreis := Kreis{10}
		assertCorrectMessageArea(t, kreis, 314.1592653589793)
	})

	// Umfang
	t.Run("Test Umfang des Rechtecks", func(t *testing.T) {
		rechteck := Rechteck{12, 6}
		assertCorrectMessagePerimeter(t, rechteck, 36)
	})
	t.Run("Test Umfang des Kreises", func(t *testing.T) {
		kreis := Kreis{10}
		assertCorrectMessagePerimeter(t, kreis, 62.83185307179586)
	})

	//// Höhe und Breite (Frage
	//t.Run("Test RECHTECK: Gebe Höhe und Breite zurück", func(t *testing.T) {
	//	rechteck := Rechteck{6, 12}
	//	want := []float64{12, 6}
	//	assertCorrectMessageHochUndBreit(t, rechteck, want )
	//})
	// Höhe und Breite (Frage
	t.Run("Test RECHTECK: Gebe Höhe und Breite zurück", func(t *testing.T) {
		rechteck := Rechteck{6, 12}
		//want := []float64{12, 6}
		assertCorrectMessageHochUndBreit(t, rechteck, 12, 6)
	})

	t.Run("Test KREIS: Gebe Höhe und Breite zurück", func(t *testing.T) {
		kreis := Kreis{10}
		//want := []float64{20, 20}
		assertCorrectMessageHochUndBreit(t, kreis, 20, 20)
	})

	//t.Run("Back the same", func(t *testing.T) {
	//	rechteck := Rechteck{12, 6}
	//	width, height := rechteck.Gap()
	//	assertCorrectMessageForGap(t, rechteck, width, 12.0)
	//	assertCorrectMessageForGap(t, rechteck, height, 6.0)
	//})
	//
	//t.Run("Test Circles", func(t *testing.T) {
	//	circle := Circle{10}
	//	assertCorrectMessage(t, circle, 314.1592653589793)
	//})

}
