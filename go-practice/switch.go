package main

import (
	"fmt"
)

func number() (num int) {
	num = 15 * 5
	return
}

func main() {
	switch num := number(); {
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200\n", num)
	}

	/*
	   num := 75
	   switch {
	   case num >= 0 && num < 50:
	       fmt.Println("num is greater than or equal 0 and less than 50")
	   case num >= 50 && num < 100:
	       fmt.Println("num is greater than or equal 50 and less than 100")
	   case num >= 100:
	       fmt.Println("num is greater than or equal 100")
	   }
	*/

	/*
	   letter := "i"
	   switch letter {
	   case "a", "e", "i", "o", "u":
	       fmt.Println("vowel")
	   default:
	       fmt.Println("not a vowel")
	   }
	*/

	/*
	   switch finger := 8; finger {
	   case 1:
	       fmt.Println("Thumb")
	   case 2:
	       fmt.Println("Index")
	   case 3:
	       fmt.Println("Middle")
	   case 4:
	       fmt.Println("Ring")
	   case 5:
	       fmt.Println("Pinky")
	   default:
	       fmt.Println("Incorrect finger number")
	   }
	*/
}
