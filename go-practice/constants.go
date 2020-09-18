package main

import (
	"fmt"
	// "math"
)

func main() {
	var a = 5.9 / 8
	fmt.Printf("a's type %T value %v\n", a, a)

	/*
		var i = 5
		var f = 5.6
		var c = 5 + 6i
		fmt.Printf("i's type %T, f's type %T, c's type %T\n", i, f, c)
	*/

	/*
		const a = 5
		var intVar int = a
		var int32Var int32 = a
		var float64Var float64 = a
		var complex64Var complex64 = a
		fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)
	*/

	/*
		const trueConst = true
		type myBool bool
		var defaultBool = trueConst
		var customBool myBool = trueConst
		defaultBool = customBool
	*/

	/*
		var defaultName = "bto"
		type myString string
		var customName myString = "bto"
		customName = defaultName
	*/

	/*
		var name = "bto"
		fmt.Printf("type %T value %v\n", name, name)
	*/

	/*
		fmt.Println("Hello, playground")
		var a = math.Sqrt(4)
		const b = math.Sqrt(4)
	*/
}
