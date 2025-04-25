package main

import "fmt"

type Student struct {
	Name   string
	School string
	Year   int
}

func main() {
	student1 := Student{
		Name:   "Alex",
		School: "Founders High School",
		Year:   2004,
	}

	fmt.Println("Student Information:")
	fmt.Println("Name:", student1.Name)
	fmt.Println("School:", student1.School)
	fmt.Println("Year:", student1.Year)
}
