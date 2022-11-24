package roundUp

func RoundToNext5(input int) int {
	if input%5 == 0 {
		return input
	} else if input > 0 {
		return input + (5 - input%5)
	} else {
		return input - input%5
	}
}
