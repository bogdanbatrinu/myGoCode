package main

import "fmt"

func main() {
	var favBook string
	var favBookAuthor string
	var favBookPublicationYear int

	fmt.Println("What is your favourite book?: ")

	fmt.Println("Who is the author of your favourite book?: ")

	fmt.Println("What is the publication year of your favourite book?: ")

	fmt.Printf("favBook has value: %v and data type %T\n", favBook, favBook)
	fmt.Printf("favBookAuthor has value: %v and data type %T\n", favBookAuthor, favBookAuthor)
	fmt.Printf("favBookPublicationYear has value: %v and data type %T\n", favBookPublicationYear, favBookPublicationYear)
}
