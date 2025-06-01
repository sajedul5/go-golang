package main

import "fmt"

var a = 10

func main() {

	fmt.Println(a)
	
}


func init() {	
	fmt.Println("init function called")
	fmt.Println(a)
	a = 20

}