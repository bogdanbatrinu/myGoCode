package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var favBook string
	var favBookAuthor string
	var favBookPublicationYear int
	memory := bufio.NewReader(os.Stdin)

	fmt.Println("What is your favourite book?: ")
	favBook, _ = memory.ReadString('\n')

	fmt.Println("Who is the author of your favourite book?: ")
	favBookAuthor, _ = memory.ReadString('\n')
	fmt.Println("What is the publication year of your favourite book?: ")
	fmt.Scanln(&favBookPublicationYear)

	fmt.Printf("favBook has value: %v and data type %T\n", favBook, favBook)
	fmt.Printf("favBookAuthor has value: %v and data type %T\n", favBookAuthor, favBookAuthor)
	fmt.Printf("favBookPublicationYear has value: %v and data type %T\n", favBookPublicationYear, favBookPublicationYear)
}
