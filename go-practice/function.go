package main

import (
	"fmt"
)

func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return
}

func main() {
	area, _ := rectProps(10.8, 5.6)
	fmt.Printf("Area %f", area)

	/*
		area, perimeter := rectProps(10.8, 5.6)
		fmt.Printf("Area %f, Perimeter %f", area, perimeter)
	*/

	/*
		price, no := 90, 6
		totalPrice := calculateBill(price, no)
		fmt.Println("Total price is", totalPrice)
	*/
}
