package main

import "fmt"

// Go does not support the ternary operator

func main() {
	age := 18
	// Normal if
	if age >= 18 {
		fmt.Println("You are adult")
	}

	// If-else
	if age >= 18 {
		fmt.Println("You are adult")

	} else if age >= 12 {
		fmt.Println("You are teenager")
	} else {
		fmt.Println("You are kid")
	}

	// If-else && / ||
	hasVoterId := true
	if age >= 18 && hasVoterId {
		fmt.Println("You are eligibile for voting")
	} else {
		fmt.Println("You are not eligibile")
	}

	role := "admin"
	hasPermissions := true

	if role == "admin" || hasPermissions {
		fmt.Println("yes")
	}

	// if with variable decleration

	if num := 2; num == 2 {
		fmt.Println(num, "Is the only prime number is even")
	}

}
