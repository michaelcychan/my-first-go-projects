package main

import "strconv"

func stringToNum(str string) int {
	output, err := strconv.Atoi(str)
	if err == nil {
		return output
	}
	panic(err)
}
