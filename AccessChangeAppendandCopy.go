package main

import "fmt"

func main() {
	//create an array
	score := [3]int{10, 9, 8}

	// access elements through index
	fmt.Println("The first element is", score[0])
fmt.Println("The lenght of the array is", len(score))
	
	
	
	//modify the last element
	score[2] = 7    //now the value of the last element is changed
	score[1] = 8    //now the value of the second element is changed

	fmt.Println("The last element is", score[2])
	fmt.Println("The second element is", score[1])


	//create a slice
	scoreSlide := []int{0,1,2,3}
	//append the element =  add new elements
	scoreSlide = append(scoreSlide, 4)	

	//slice range
	scoreSlide = scoreSlide[1:]
	//remain from the first index and cut from the next index 
	// if the next index after the : is empty, it will not slice
	fmt.Println("The score slice has",scoreSlide)

	
}