package main

import (
	"fmt"
)

type Employee struct {
	firstName, lastName string
	age, salary         int
}

type Person struct {
	string
	int
}

type Address struct {
	city, state string
}

type Person2 struct {
	name    string
	age     int
	address Address
}

type Person3 struct {
	name string
	age  int
	Address
}

func main() {
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}

	emp2 := Employee{"Thomas", "Paul", 29, 800}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}

	fmt.Println("Employee 3", emp3)

	var emp4 Employee

	fmt.Println("Employee 4", emp4)

	emp5 := Employee{
		firstName: "John",
		lastName:  "Paul",
	}

	fmt.Println("Employee 5", emp5)

	emp6 := Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Println("Salary:", emp6.salary)

	var emp7 Employee
	emp7.firstName = "Jack"
	emp7.lastName = "Adams"
	fmt.Println("Employee 7", emp7)

	emp8 := &Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name", emp8.firstName)
	fmt.Println("Age", emp8.age)

	p := Person{"Naveen", 50}
	fmt.Println(p)

	var p1 Person
	p1.string = "naveen"
	p1.int = 50
	fmt.Println(p1)

	var p2 Person2
	p2.name = "Naveen"
	p2.age = 50
	p2.address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p2.name)
	fmt.Println("Age:", p2.age)
	fmt.Println("City:", p2.address.city)
	fmt.Println("State:", p2.address.state)

	var p3 Person3
	p3.name = "Naveen"
	p3.age = 50
	p3.Address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p3.name)
	fmt.Println("Age:", p3.age)
	fmt.Println("City:", p3.city)
	fmt.Println("State:", p3.state)
}
