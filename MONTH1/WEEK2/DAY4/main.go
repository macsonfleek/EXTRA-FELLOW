 package main

 import "fmt"

 func main() {
	for x := 0; x < 7; x++{
		fmt.Println("Dapo")
	}

for val := range 7{
	fmt.Println("Dapo", val)
}

x := 1
for {
	if x == 8 {
		break
	}
	fmt.Println("Dapo")
	x++
}

}
