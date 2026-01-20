package main

import "fmt"

// Array are fixed size collection of squential homogenious elements
// If declare the array in golang and print it without adding elements in it then it shows zeroed value. For int -> 0 , string -> "" , bool -> false , float -> 0.0
//Array used where the length is already known
func main() {
	// 1-D array
	var nums [5]int
	fmt.Println(nums) // array with zeroed value

	// add element in array
	nums[0] = 1
	fmt.Println(nums)

	// Find length of array
	fmt.Println("Length:", len(nums))

	// 2-D array
	names := [2][2]string{{"Golang", "Go"}, {"C++", "CPP"}}
	fmt.Println(names)
}
