package main

import (
	"fmt"
)

// Range is used to iterate over slices, maps and strings
func main() {
	// Iterate over slice int
	nums := []int{1, 2, 3, 4, 5}

	for index, value := range nums {
		fmt.Println(index, "-", value)
	}

	// Iterate over maps
	m := map[string]string{"name": "Golang", "version": "go1.24.2 windows/amd64"}

	for key, value := range m {
		fmt.Println("Key:", key, " value:", value)
	}

	// Iterate over string
	s := "golang"

	for i, charUnicode := range s {
		fmt.Println("i:", i, " ch:", string(charUnicode), "unicode:", charUnicode)
	}
}
