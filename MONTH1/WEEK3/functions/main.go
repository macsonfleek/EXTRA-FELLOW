package main

import "fmt"


func main() {
	//classifyAge(5)

	fmt.Println(positiveSum(5))
}

func classifyAge(age int) {
	if age >= 0 && age <= 12 {
		fmt.Println("Child")
	} else if age >= 13 && age <= 19 {
		fmt.Println("Teenager")
	} else if age >= 20 {
		fmt.Println("Adult")
	}
}

func positiveSum(num int) int {

	sum := 0

	for i := 1 ; i <= num ; i++ {
		sum += i
	}
	return sum
}
