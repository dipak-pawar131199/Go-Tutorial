package main

import (
	"fmt"
	"slices"
)

// Slices : Dynamic arrays
// Slice are used where we don't know the number of elements.

func main() {
	// Create slice without make
	var nums []int
	// add elements in slice using append
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)

	fmt.Println(nums)

	// Create slice with make
	// make([]type,length,capacity)
	names := make([]string, 0, 3)
	names = append(names, "Dipak")
	names = append(names, "Rajendra")
	names = append(names, "Tushar")
	fmt.Println("capacity", cap(names))
	names = append(names, "Prathmesh")
	// Capacity get double when the length slice > capacity
	fmt.Println("capacity", cap(names))
	fmt.Println(names)

	// Copy slice : it copy only when source and destination slices length are same
	num := []int{1, 2, 3}
	num1 := make([]int, len(num))
	copy(num1, num)
	fmt.Println(num, num1)

	// Slices package : check slice are equal
	fmt.Println(slices.Equal(num, num1))

	// 2-D slices

	num2 := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("2-D Slice:", num2)

}
