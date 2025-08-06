package main

import "fmt"

func main() {
	//create an array
	score := [3]int{10, 9, 8}

	// access elements through index
	fmt.Println("The first element is", score[0])

	//modify the last element
	score[2] = 7    //now the value of the last element is changed
	score[1] = 8    //now the value of the second element is changed

	fmt.Println("The last element is", score[2])
	fmt.Println("The second element is", score[1])

	//append the first element
	
}