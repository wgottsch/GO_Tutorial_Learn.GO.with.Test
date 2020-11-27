package structs

import "math"

type Rechteck struct {
	Breite float64
	Hoehe  float64
}

func (r Rechteck) Flaeche() float64 {
	return r.Breite * r.Hoehe
}

func (r Rechteck) Umfang() float64 {
	return (r.Breite * 2) + (r.Hoehe * 2)
}

func (r Rechteck) Hoch() float64 {
	return r.Hoehe
}
func (r Rechteck) Breit() float64 {
	return r.Breite
}

// old
func (r Rechteck) HochUndBreit() (float64, float64) {
	return r.Hoehe, r.Breite
}

//new
//func (r Rechteck) HochUndBreit() []float64 {
//	arr := []float64{r.Hoehe, r.Breite}
//	//arr[0] = r.Hoehe
//	//arr[1] = r.Breite
//	return arr
//}

type Kreis struct {
	Radi float64
}

func (c Kreis) Flaeche() float64 {
	return math.Pi * c.Radi * c.Radi
}

func (c Kreis) Umfang() (v1 float64) {
	return 2 * math.Pi * c.Radi
}

func (r Kreis) Hoch() float64 {
	return r.Radi * 2
}
func (r Kreis) Breit() float64 {
	return r.Radi * 2
}

// OLD
func (r Kreis) HochUndBreit() (float64, float64) {
	return r.Radi * 2, r.Radi * 2
}

////new
//func (r Kreis) HochUndBreit() []float64 {
//	arr := []float64{r.Radi * 2, r.Radi * 2}
//	//arr[0] = r.Hoehe
//	//arr[1] = r.Breite
//	return arr
//}

type Compute interface {
	Umfang() float64
	Flaeche() float64
	Hoch() float64
	Breit() float64
	HochUndBreit() (float64, float64)
}

//func Perimeter(rectangle Rectangle) float64 {
//	return 2 * (rectangle.Width + rectangle.Height)
//}
//func Area(rectangle Rectangle) float64 {
//	return rectangle.Width * rectangle.Height
//}
