package main

import "fmt"

var a = 10

func main() {

	fmt.Println(a)
	
}





/*

In Go (Golang), the init function is a special function that is automatically executed when a package is initialized, 
before the main() function runs. It's commonly used to:

Set up initial state

Initialize variables


Key Points
Each package can have multiple init() functions, even in different files within the same package.

The order of init() function execution:

Imported packages are initialized first (recursively, depth-first).

Within the same package, init() runs in the order the files are compiled (not guaranteed unless explicitly ordered).

init() takes no parameters and returns no values.

It cannot be called manually—only run automatically during program startup.
*/









func init() {	
	fmt.Println("init function called")
	fmt.Println(a)
	a = 20

}