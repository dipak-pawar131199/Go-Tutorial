package main

import "fmt"

// Functions are used for reused
// In golang main is entry point function not need call manually it is called by golang
// In golang functions can return multiple values
// In golang functions can take as funcation arugrument and also return functions
/*
  Syntax:
	 func func_name([argumentslist with data type])(return data-type){
	    // funcation body
	 }
*/

// function retun single value
func add(num1, num2 int) int {
	return num1 + num2
}

// function returns muliple values
func division(num1 float32, num2 float32) (float32, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("Divide by zero error")
	} else {
		return num1 / num2, nil
	}
}

// function return func
func Operation(num1 int) func() int {
	return func() int {
		return num1
	}
}

// func as argument in function
func funcArugument(func(int) int) int {
	return 10
}
func main() {
	fmt.Println(add(10, 5))

	fmt.Println(division(10, 3))
	fmt.Println(division(10, 0))

	f := Operation(10)
	fmt.Println(f())

	fmt.Println(funcArugument(func(num int) int {
		return num
	}))

}
