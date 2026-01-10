package main

// Constant : Are the variable we declear with const keyword.
// Once constant variable are assign some value, we can't modify / change value

import "fmt"

// Gloabal constant
const pi = 3.14

func main() {
	fmt.Println("Pi:", pi)

	// We can group the constant variables

	const (
		numDayInWeek    = 7
		numMonthsInYear = 12
	)

	fmt.Println(numDayInWeek, numMonthsInYear)
}
