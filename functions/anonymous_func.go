package main

import "fmt"


func main() {
	
	// Anonymous functions are functions that are defined without a name.
	// They can be assigned to variables, passed as arguments, or returned from other functions.
	// IIFEs (Immediately Invoked Functions) are a common use case for anonymous functions.

	func (a int, b int) {
		c := a + b
		fmt.Println("Sum:", c)
	}(10, 20) // This is an anonymous function that adds two numbers and prints the result.

	
}


func init() {	
	fmt.Println("init function called")
	
}