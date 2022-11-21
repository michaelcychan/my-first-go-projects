package calculator

func Sum(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	numberOfSlices := len(slices)             // get the number of slices in the spread variables
	outputSums := make([]int, numberOfSlices) // create a slice with []int, and number of which
	for index, numbers := range slices {
		outputSums[index] = Sum(numbers)
	}
	return outputSums
}
