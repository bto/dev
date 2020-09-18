package main

import (
	"fmt"
)

func change(val *int) {
	*val = 55
}

func modify(sls []int) {
	sls[0] = 90
}

func main() {
	b := [...]int{109, 110, 111}
	p := &b
	p++

	/*
		a := [3]int{89, 90, 91}
		modify(a[:])
		fmt.Println(a)
	*/

	/*
		a := 58
		fmt.Println("value of a before function call is", a)
		b := &a
		change(b)
		fmt.Println("value of a after function call is", a)
	*/

	/*
		b := 255
		a := &b
		fmt.Println("address of b is", a)
		fmt.Println("value of b is", *a)
		*a++
		fmt.Println("new value of b is", *a)
	*/

	/*
		a := 25
		var b *int
		if b == nil {
			fmt.Println("b is", b)
			b = &a
			fmt.Println("b after initialization is", b)
		}
	*/

	/*
		b := 255
		var a *int = &b
		fmt.Printf("Type of a is %T\n", a)
		fmt.Println("address of a is", a)
	*/
}
