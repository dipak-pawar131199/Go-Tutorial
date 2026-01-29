package main

import "fmt"

// Variadic function used where num of arugument are not known
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {

	fmt.Println(sum(1, 2, 3, 4, 5))
	nums := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(sum(nums...))
}
