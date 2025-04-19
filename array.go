package main

import "fmt"

// Array
func array_1() {

	var arr1 = [100]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)
}

func array_2() {

	var arr1 = [...]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)
}

func array_string() {

	var cars = [5]string{"BMW", "Mercedes", "Toyota", "Honda", "Ford"}
	cars[2] = "Audi"
	fmt.Println(cars)
}

func prices_int() {

	var prices = [3]int{20, 25, 100}
	prices[1] = 50
	fmt.Println(prices)
}

func array_len() {

	aar1 := [4]string{"BMW", "Mercedes", "Toyota", "Honda"}
	arr2 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(aar1))
	fmt.Println(len(arr2))

}

func main() {

	// array_1()
	// array_2()
	// array_string()
	// prices_int()

	array_len()

}
