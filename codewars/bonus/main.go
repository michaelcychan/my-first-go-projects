// https://www.codewars.com/kata/56f6ad906b88de513f000d96/go

package kata

import "fmt"

func BonusTime(salary int, bonus bool) string {
	if bonus {
		return fmt.Sprintf("£%d", salary*10)
	} else {
		return fmt.Sprintf("£%d", salary)
	}
}
