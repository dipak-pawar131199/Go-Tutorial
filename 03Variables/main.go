package main

import "fmt"

func main() {
	/* Variable decleration*/
	/**
	   syntax:
		    var var_name [type] = [value]
	*/

	var name string = "golang"
	fmt.Println("name:", name)

	// without type
	var name1 = "hello"
	fmt.Println("name1:", name1)

	var isVisible bool
	fmt.Println(isVisible)

	// short hand
	/**
	  syntax: var_name := value
	*/
	// Note:- We can't used with short hand variable decleration for global variable decleration

	age := 30
	fmt.Println("age:", age)

}
