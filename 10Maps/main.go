package main

import (
	"fmt"
	"maps"
)

// Map used store data in key-value pair

func main() {
	// Create map using make()
	// syntax: make(map[key-type]value-type)

	m := make(map[string]string)

	// add key-value
	m["Name"] = "Golang"
	m["Version"] = "go1.24.2 windows/amd64"

	fmt.Println(m)

	// Delete element from map

	delete(m, "Version")
	fmt.Println(m)

	// Clear : delete all elements from map

	clear(m)
	fmt.Println(m)

	// Compare map
	m1 := map[string]int{"Phone": 1122345, "Age": 27}
	m2 := map[string]int{"Phone": 1122345, "age": 27}
	fmt.Println(maps.Equal(m1, m2))

	// Check key exist in map
	value, ok := m1["Phone"]
	fmt.Println(value)
	if ok {
		fmt.Println("Key present")
	} else {
		fmt.Println("Key not present")
	}

}
