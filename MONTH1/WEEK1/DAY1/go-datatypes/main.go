package main

import "fmt"

func main() {
	var name string = "Adedapo Amuda"

	var age int = 30

	//var lovesLinux bool = true

	//var hight float64 = 5.84

	//color = "sky blue"

	//tinyInt = 200

	name = "Samuel"

	age = 45

	//lovesLinux = false

	//height = 8.5

	//color = "Purple"

	//tinyInt = 255

	fmt.Print("His name is ", name, ", and he is ", age, " years old. \n")
	fmt.Println("His name is", name+",", "and he is", age, "years old.")
	fmt.Printf("His name is %v, and he is %v years old. \n", name, age )
	fmt.Printf("His name is %s, and he is %d years old. \n", name, age )
}
