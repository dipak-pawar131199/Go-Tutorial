package main

import (
	"fmt"
	"time"
)

func main() {
	num := 5
	// Normal switch
	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")

	default:
		fmt.Println("Other")
	}
	// Switch with conditions

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Work day")
	}

	// type chacking switch

	checkType := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("Integer type value")
		case string:
			fmt.Println("String type value")
		case float32:
			fmt.Println("Float type value")
		default:
			fmt.Println("Other")
		}
	}
	checkType(10)
	checkType("10")
	checkType(10.0)

}
