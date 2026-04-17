package main

import "fmt"


func main() {
	// userAge := 20
	// AddTenYears(&userAge)
	// fmt.Println(userAge)
	userName := "Dapo"
	AddSureName(&userName)
	fmt.Println(userName)

}


func AddTenYears(age *int) {

	*age += 10

}

func AddSureName(name *string) {
	*name += " Amuda"
}
