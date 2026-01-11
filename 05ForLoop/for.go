package main

// Golang support only one looping construct i.e. for loop

import "fmt"

func main() {
	// while loop
	fmt.Println("while loop like for")

	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
	// classic for loop
	fmt.Println("Classic for loop")
	for j := 0; j <= 4; j++ {
		fmt.Println(j)
	}
	// for loop with break and continue
	fmt.Println("for loop with break and continue")

	for j := 0; j <= 4; j++ {
		if j == 2 {
			continue
		}

		if j == 3 {
			break
		}

		fmt.Println(j)
	}
	// infinite loop
	fmt.Println("Infinite for loop")

	for {
		fmt.Println("hi")
	}
}
