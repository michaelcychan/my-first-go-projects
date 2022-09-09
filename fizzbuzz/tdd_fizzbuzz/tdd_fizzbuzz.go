package main

import "fmt"

func tddFizzBuzz(num int) string {
	var outputString string
	var divisible bool
	if divisibleBy(num, 3) {
		outputString = outputString + "Fizz"
		divisible = true
	}
	if divisibleBy(num, 5) {
		outputString = outputString + "Buzz"
		divisible = true
	}
	if !divisible {
		outputString = fmt.Sprint(num)
	}
	return outputString
}

func divisibleBy(dividend, divisor int) bool {
	if dividend%divisor == 0 {
		return true
	} else {
		return false
	}
}
