package main

import (
	"fmt"
)

func main() {
	num := 99
	if num <= 50 {
		fmt.Println("number is less than or equal to 50")
	} else if num >= 51 && num <= 100 {
		fmt.Println("number is between 51 and 100")
	} else {
		fmt.Println("number is greagter than 100")
	}

	/*
		if num := 10; num%2 == 0 {
			fmt.Println("the number is even")
		} else {
			fmt.Println("the number is odd")
		}
	*/
}
