package cascade

func EachCons(input []int, size int) [][]int {
	var outputSlice [][]int

	for i := 0; i <= len(input)-size; i++ {
		outputSlice = append(outputSlice, input[i:i+size])
	}

	return outputSlice
}
