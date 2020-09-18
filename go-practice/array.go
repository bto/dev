package main

import (
	"fmt"
)

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function", num)
}

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Println()
	}
}

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries)
	return countriesCpy
}

func main() {
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)

	/*
		pls := [][]string {
			{"C", "C++"},
			{"Javascript"},
			{"Go", "Rust"},
		}
		for _, v1 := range pls {
			for _, v2 := range v1 {
				fmt.Printf("%s ", v2)
			}
			fmt.Println()
		}
	*/

	/*
		nos := []int{8, 7, 6}
		fmt.Println("slice before function call", nos)
		subtactOne(nos)
		fmt.Println("slice after function call", nos)
	*/

	/*
		veggies := []string{"potatoes", "tomatoes", "brinjal"}
		fruits := []string{"oranges", "apples"}
		food := append(veggies, fruits...)
		fmt.Println("food:", food)
	*/

	/*
		var names []string
		if names == nil {
			fmt.Println("slice is nil going to append")
			names = append(names, "John", "Sebastian", "Vinay")
			fmt.Println("names contents:", names)
		}
	*/

	/*
		cars := []string{"Ferrari", "Honda", "Ford"}
		fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))
		cars = append(cars, "Toyota")
		fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars))
	*/

	/*
		i := make([]int, 5, 5)
		fmt.Println(i)
	*/

	/*
		fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
		fruitslice := fruitarray[1:3]
		fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice))
		fruitslice = fruitslice[:cap(fruitslice)]
		fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice))
	*/

	/*
		numa := [3]int{78, 79, 80}
		nums1 := numa[:]
		nums2 := numa[:]
		fmt.Println("array before change 1", numa)
		nums1[0] = 100
		fmt.Println("array after modification to slice nums1", numa)
		nums2[1] = 101
		fmt.Println("array after modification to slice nums2", numa)
	*/

	/*
		darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
		dslice := darr[2:5]
		fmt.Println("array before", darr)
		for i := range dslice {
			dslice[i]++
		}
		fmt.Println("array after", darr)
	*/

	/*
		c := []int{6, 7, 8}
		fmt.Println(c)
	*/

	/*
		a := [5]int{76, 77, 78, 79, 80}
		var b []int = a[1:4]
		fmt.Println(b)
	*/

	/*
		a := [3][2]string{
			{"lion", "tiger"},
			{"cat", "dog"},
			{"pigeon", "peacock"},
		}
		printarray(a)
		fmt.Println()

		var b [3][2]string
		b[0][0] = "apple"
		b[0][1] = "samsung"
		b[1][0] = "microsoft"
		b[1][1] = "google"
		b[2][0] = "AT&T"
		b[2][1] = "T-Mobile"
		printarray(b)
	*/

	/*
		a := [...]float64{67.7, 89.8, 21, 78}
		sum := float64(0)
		for i, v := range a {
			fmt.Printf("%d th element of a is %.2f\n", i, v)
			sum += v
		}
		fmt.Println("sum of all elements of a is", sum)
	*/

	/*
		a := [...]float64{67.7, 89.8, 21, 78}
		for i := 0; i < len(a); i++ {
			fmt.Printf("%d th element of a is %.2f\n", i, a[i])
		}
	*/

	/*
		a := [...]float64{67.7, 89.8, 21, 78}
		fmt.Println("length of a is", len(a))
	*/

	/*
		num := [...]int{5, 6, 7, 8, 9}
		fmt.Println("before passing to function", num)
		changeLocal(num)
		fmt.Println("after passing to function", num)
	*/

	/*
		a := [...]string{"USA", "China", "India", "Germany", "France"}
		b := a
		b[0] = "Singapore"
		fmt.Println("a is", a)
		fmt.Println("b is", b)
	*/

	/*
		a := [3]int{5, 78, 8}
		var b [5]int
		b = a
	*/

	/*
		a := [...]int{12, 78, 50}
		fmt.Println(a)
	*/

	/*
		a := [3]int{12, 78, 50}
		fmt.Println(a)
	*/

	/*
		var a [3]int
		a[0] = 12
		a[1] = 78
		a[2] = 50
		fmt.Println(a)
	*/
}
