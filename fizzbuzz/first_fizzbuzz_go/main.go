package main

import "fmt"

func fizzbuzz(number int) string {
	if number%15 == 0 {
		return "FizzBuzz"
	} else if number%5 == 0 {
		return "Buzz"
	} else if number%3 == 0 {
		return "Fizz"
	} else {
		return fmt.Sprint(number)
	}
}

func main() {
	var testNum int
	fmt.Println("Enter a number:")
	fmt.Scan(&testNum)
	fmt.Println(fizzbuzz(testNum))
}
