package main

import "fmt"



func main() {
	// var num int = 37

	// if num > 0 {
	// 	fmt.Printf("%v is a Positive number \n", num)
	// } else if num < 0 {
	// 	fmt.Printf("%v is a negative number \n", num)
	// } else {
	// 		fmt.Printf("%v is a neutral number \n", num)
	// }
	const age = -4

	// if age < 0 {
	// 	fmt.Println("invalid age!")
	// }

	if age >= 0 && age <= 12 {
		fmt.Println("User is a Child")
	} else if age >= 13 && age <= 19 {
		fmt.Println("User is a Teenager")
	} else if age >= 20 {
		fmt.Println("User is an adult")
	}
}
