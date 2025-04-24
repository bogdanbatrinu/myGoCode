package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your full name")
	name, _ := reader.ReadString('\n')
	fmt.Println("Your name is ", name)
	fmt.Print("Enter your favourite number")
	favNum, _ := reader.ReadString('\n')
	fmt.Println("Your favourite number", favNum)
	fmt.Print("Enter your fun fact")
	funFact, _ := reader.ReadString('\n')
	fmt.Println("Your fun fact is ", funFact)
}
