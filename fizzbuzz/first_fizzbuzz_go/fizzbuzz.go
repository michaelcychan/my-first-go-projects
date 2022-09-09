package main

import "fmt"

func fizzbuzz(number int) string {
	if divisibleBy(number, 15) {
		return "FizzBuzz"
	} else if divisibleBy(number, 5) {
		return "Buzz"
	} else if divisibleBy(number, 3) {
		return "Fizz"
	} else {
		return fmt.Sprint(number)
	}
}

func divisibleBy(dividend, divisor int) bool {
	if dividend%divisor == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var testNum int
	fmt.Println("Enter a number:")
	fmt.Scan(&testNum)
	fmt.Println(fizzbuzz(testNum))
}
