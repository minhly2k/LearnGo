package main

import "fmt"

const name1 = "Ngân"

func main() {
	var (
		array1 = [2]int{0, 1}
		slice1 = []float32{3.14, 23.32}
	)
	intNum := make([]int, 2, 4)
	const NAME = "Hung" //UPPERCASE
	fmt.Println(array1)
	fmt.Println(slice1)
	fmt.Println(intNum)
	fmt.Println(NAME)
	fmt.Println(name1)

	//3 ways to create slices

	sliceExample1 := [2]bool{true, false} //infered   1
	//create slice from array

	//use make() function
	//This is the most idiomatic and generally preferred way to declare slices in Go.
	//It's the standard approach and clearly communicates the intent without initializing its elements.
	//  It's a placeholder – it allocates the memory but doesn't give you any data in that memory yet.
	//It returns a pointer to the newly allocated slice.
	//Crucially, it doesn't return a value itself. This is the most important point.
	sliceExample2 := make([]int, 1, 2) //1 element, 2 is the length

	fmt.Println(sliceExample1, sliceExample2)
}
