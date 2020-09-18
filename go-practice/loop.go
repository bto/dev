package main

import (
	"fmt"
)

func main() {
	i := 0
	for i <= 10 {
		fmt.Printf("%d ", i)
		i += 2
	}
	fmt.Printf("\nline after for loop\n")

	/*
		outer:
			for i := 0; i < 3; i++ {
				for j := 1; j < 4; j++ {
					fmt.Printf("i = %d, j = %d\n", i, j)
					if i == j {
						break outer
					}
				}
			}
	*/

	/*
	   n := 5
	   for i := 0; i < n; i++ {
	       for j := 0; j < i; j++ {
	           fmt.Print("*")
	       }
	       fmt.Println()
	   }
	*/

	/*
	   for i := 1; i <= 10; i++ {
	       if i % 2 == 0 {
	           continue
	       }
	       fmt.Printf("%d ", i)
	   }
	   fmt.Printf("\nline after for loop\n")
	*/

	/*
	   for i := 1; i <= 10; i++ {
	       if i < 5 {
	           break
	       }
	       fmt.Printf("%d ", i)
	   }
	   fmt.Printf("\nline after for loop\n")
	*/

	/*
	   for i := 1; i <= 10; i++ {
	       fmt.Printf("%d ", i)
	   }
	   fmt.Printf("\nline after for loop\n")
	*/
}
