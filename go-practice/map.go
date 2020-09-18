package main

import (
	"fmt"
)

func main() {
	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	map2 := map1
	if map1 == map2 {
	}

	/*
		personSalary := map[string]int{
			"steve": 12000,
			"jamie": 15000,
			"mike":  9000,
		}
		fmt.Println("Original person salary", personSalary)
		newPersonSalary := personSalary
		newPersonSalary["mike"] = 18000
		fmt.Println("Person salary changed", personSalary)
	*/

	/*
		personSalary := map[string]int{
			"steve": 12000,
			"jamie": 15000,
			"mike":  9000,
		}
		fmt.Println("lenght is", len(personSalary))
	*/

	/*
		personSalary := map[string]int{
			"steve": 12000,
			"jamie": 15000,
			"mike":  9000,
		}
		fmt.Println("map before deletion", personSalary)
		delete(personSalary, "steve")
		fmt.Println("map after deletion", personSalary)
	*/

	/*
		personSalary := map[string]int{
			"steve": 12000,
			"jamie": 15000,
			"mike":  9000,
		}
		fmt.Println("All items of a map")
		for key, value := range personSalary {
			fmt.Printf("personSalary[%s] = %d\n", key, value)
		}
	*/

	/*
		personSalary := map[string]int{
			"steve": 12000,
			"jamie": 15000,
			"mike":  9000,
		}
		newEmp := "jamie"
		value, ok := personSalary[newEmp]
		if ok == true {
			fmt.Println("Salary of", newEmp, "is", value)
		} else {
			fmt.Println(newEmp, "not found")
		}
		newEmp = "joe"
		value, ok = personSalary[newEmp]
		if ok == true {
			fmt.Println("Salary of", newEmp, "is", value)
		} else {
			fmt.Println(newEmp, "not found")
		}
	*/

	/*
		personSalary := map[string]int{
			"steve": 12000,
			"jamie": 15000,
			"mike":  9000,
		}
		employee := "jamie"
		fmt.Println("Salary of", employee, "is", personSalary[employee])
		fmt.Println("Salary of joe is", personSalary["joe"])
	*/

	/*
		personSalary := map[string]int{
			"steve": 12000,
			"jamie": 15000,
		}
		personSalary["mike"] = 9000
		fmt.Println("personSalary map contents:", personSalary)
	*/

	/*
		personSalary := make(map[string]int)
		personSalary["steve"] = 12000
		personSalary["jamie"] = 15000
		personSalary["mike"] = 9000
		fmt.Println("personSalary map contents:", personSalary)
	*/

	/*
		var personSalary map[string]int
		if personSalary == nil {
			fmt.Println("map is nil. Going to make one.")
			personSalary = make(map[string]int)
		}
	*/
}
