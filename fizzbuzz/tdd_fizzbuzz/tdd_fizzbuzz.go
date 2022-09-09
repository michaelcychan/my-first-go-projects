package main

import "fmt"

func tddFizzBuzz(num int) string {
	if num == 3 || num == 6 {
		return "Fizz"
	} else if num == 5 {
		return "Buzz"
	} else {
		return fmt.Sprint(num)
	}
}
