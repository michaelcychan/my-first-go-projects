package calculator

func Sum(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	var outputSums []int
	for _, numbers := range slices {
		outputSums = append(outputSums, Sum(numbers))
	}
	return outputSums
}

func SumAllTails(slices ...[]int) []int {
	var outputSums []int
	for _, numbers := range slices {
		if len(numbers) == 0 {
			outputSums = append(outputSums, 0)
		} else {
			tail := numbers[1:]
			outputSums = append(outputSums, Sum(tail))
		}
	}
	return outputSums
}
