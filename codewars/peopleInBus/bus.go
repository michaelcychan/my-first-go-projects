package bus

func Number(input [][2]int) int {
	var peopleInBus int
	peopleInBus = 0
	for i := 0; i < len(input); i++ {
		peopleInBus += (input[i][0] - input[i][1])
	}
	return peopleInBus
}
