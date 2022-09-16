package iteration

func Repeater(input string) string {
	var output string
	for i := 0; i < 5; i++ {
		output += input
	}
	return output
}
