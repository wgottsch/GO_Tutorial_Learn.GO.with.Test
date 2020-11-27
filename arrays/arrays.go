package arrays

//package main

import (
	"strings"
)

func Sum(numbers []int) int {
	sum := 0
	//for i := 0; i < len(numbers); i++ {
	//	sum += numbers[i]
	//} REFACTOR

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	//lengthOfNumbers := len(numbersToSum)
	//sums := make([]int, lengthOfNumbers)

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	//lengthOfNumbers := len(numbersToSum)
	//sums := make([]int, lengthOfNumbers)

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

func LastX(lx int) ([]int, []string) {
	s := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ss := [10]string{"00", "11", "22", " 33", "44", "55", "66", "77", "88", "99"}
	return s[len(s)-lx:], ss[len(ss)-lx:]

}

func LastX1(lx string) bool {
	ss := "BRAIN-Omega_dev"
	return strings.HasSuffix(ss, "_dev")

}

//Zum ausprobieren. Package dann auch "Main" nennen
//func main() {
//	got := SumAll([]int{1, 2}, []int{0, 9})
//	fmt.Println(got)
//	g2, g3 := LastX(4)
//	fmt.Println(g2, g3)
//	g4 := LastX1("BRAIN-Omega_dev")
//	fmt.Println(g4)
//}
