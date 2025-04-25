package main

import "fmt"

func main() {
	var student1 string = "John"
	var student2 = "Jane"
	var age1 int = 28
	age2 := 32
	var balance1 float64 = 5.5
	var balance2 float64 = 7.5
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(age1)
	fmt.Println(age2)
	fmt.Println(student1, "has the account balance of", balance1)
	fmt.Println(student2, "has the account balance of", balance2)
	fmt.Println("The combined income of", student1, "and", student2, "is", balance1+balance2)
}
