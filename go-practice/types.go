package main

import (
	"fmt"
	// "unsafe"
)

func main() {
	i := 10
	var j float64 = float64(i)
	fmt.Println("j", j)

	/*
		i := 55
		j := 67.8
		sum := i + int(j)
		fmt.Println(sum)
	*/

	/*
		first := "Masato"
		last := "Bito"
		name := first + " " + last
		fmt.Println("My name is", name)
	*/

	/*
		c1 := complex(5, 7)
		c2 := 8 + 27i
		cadd := c1 + c2
		fmt.Println("sum:", cadd)
		cmul := c1 * c2
		fmt.Println("product:", cmul)
	*/

	/*
		a, b := 5.67, 8.97
		fmt.Printf("type of a %T b %T\n", a, b)
		sum := a + b
		diff := a - b
		fmt.Println("sum", sum, "diff", diff)

		no1, no2 := 56, 89
		fmt.Println("sum", no1+no2, "diff", no1-no2)
	*/

	/*
		var a int
		var b int8
		var c int16
		var d int32
		var f int64
		var aa uint
		var bb uint8
		var cc uint16
		var dd uint32
		var ff uint64
		fmt.Printf("type of a is %T, size of a is %d\n", a, unsafe.Sizeof(a))
		fmt.Printf("type of b is %T, size of b is %d\n", b, unsafe.Sizeof(b))
		fmt.Printf("type of c is %T, size of c is %d\n", c, unsafe.Sizeof(c))
		fmt.Printf("type of d is %T, size of d is %d\n", d, unsafe.Sizeof(d))
		fmt.Printf("type of f is %T, size of f is %d\n", f, unsafe.Sizeof(f))
		fmt.Printf("type of aa is %T, size of aa is %d\n", aa, unsafe.Sizeof(aa))
		fmt.Printf("type of bb is %T, size of bb is %d\n", bb, unsafe.Sizeof(bb))
		fmt.Printf("type of cc is %T, size of cc is %d\n", cc, unsafe.Sizeof(cc))
		fmt.Printf("type of dd is %T, size of dd is %d\n", dd, unsafe.Sizeof(dd))
		fmt.Printf("type of ff is %T, size of ff is %d\n", ff, unsafe.Sizeof(ff))
	*/

	/*
		a := true
		b := false
		fmt.Println("a:", a, "b:", b)

		c := a && b
		fmt.Println("c:", c)

		d := a || b
		fmt.Println("d:", d)
	*/
}
