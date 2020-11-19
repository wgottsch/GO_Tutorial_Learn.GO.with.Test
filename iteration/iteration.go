package iteration

const repeatCont = 5

func Repeat(character string, count int) string {
	var repeated string
	for i := 1; i <= count; i++ {
		repeated += character
	}
	return repeated
}
