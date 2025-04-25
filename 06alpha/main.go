package main

import "fmt"

func main() {
	for abLow := 'a'; abLow <= 'z'; abLow++ {
		fmt.Printf("%c and type: %T\n", abLow, abLow)
	}
}
