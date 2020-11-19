package arrays

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
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

// Zum ausprobieren. Package dann auch "Main" nennen
//func main() {
//	got := SumAll([]int{1, 2}, []int{0, 9})
//	fmt.Println(got)
//}
