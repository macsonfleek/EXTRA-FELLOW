package main

import "fmt"

func main() {
	//CHALLENGE 1
	// fmt.Println("Printing all EVEN Numbers between 20 and 55:")
	// for i := 20; i <= 55; i++ {
	// 	if i % 2 == 0 {
	// 		fmt.Print(i, " ")
	// 	} else {
	// 		continue
	// 	}
	// }
	// fmt.Print("\n")


// CHALLENGE 2
	// myArray := [7]int{1, 2, 3, 4, 5, 6, 7}
	// newArray := [7]int{}
	// for i, val := range myArray {
	// 	result := val * 2
	// 	newArray[i] = result
	// }
	// fmt.Println("new array:")
	// fmt.Println(newArray)


// CHALLENGE 3

mySlice := []int{5, 4, 3, 2, 1}

for i := len(mySlice)-1; i >= 0; i-- {
		fmt.Println(mySlice[i], " ")
	}
	//fmt.Println("")
}
