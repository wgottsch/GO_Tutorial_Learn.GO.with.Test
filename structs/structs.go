package structs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Gap() (float64, float64) {
	return r.Width, r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// wenn ich via Shape() auf circle zugreifen will, dann muss ich auch für jede Methode in Shape() für circle auch
// eine entsprechende Funktion anlegen.....
func (c Circle) Gap() (v1 float64, v2 float64) {
	return
}

type Shape interface {
	Area() float64
	Gap() (float64, float64)
}

//func Perimeter(rectangle Rectangle) float64 {
//	return 2 * (rectangle.Width + rectangle.Height)
//}
//func Area(rectangle Rectangle) float64 {
//	return rectangle.Width * rectangle.Height
//}
