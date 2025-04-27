package main

import "fmt"

func main() {
	prices := []float64{1.99, 2.99, 3.99, 4.99}

	prices[3] = 5.99

	fmt.Println(prices[2])
	fmt.Println(prices[3])

	appendSlice()

}

func appendSlice() {
	// Append to a slice
	prices := []float64{1.99, 2.99, 3.99, 4.99}

	prices = append(prices, 5.99)
	fmt.Println(prices)
	// Append multiple values to a slice
	prices = append(prices, 6.99, 7.99)
	fmt.Println(prices)
	// Append a slice to another slice
	prices2 := []float64{8.99, 9.99}
	prices = append(prices, prices2...)
	fmt.Println(prices)

	//another slice append

	mySlice1 := []int{1, 2, 3, 4, 5}
	mySlice2 := []int{4, 5, 6}

	mySlice3 := append(mySlice1, mySlice2...)

	fmt.Println(mySlice3)
	fmt.Printf("len(mySlice3): %d\n", len(mySlice3))
	fmt.Printf("cap(mySlice3): %d\n", cap(mySlice3))
}
