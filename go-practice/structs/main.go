package main

import (
	"fmt"

	"github.com/bto/go-practice/structs/computer"
)

func main() {
	var spec computer.Spec
	spec.Maker = "apple"
	spec.Price = 50000
	spec.model = "Mac Mini"
	fmt.Println("Spec:", spec)
}
