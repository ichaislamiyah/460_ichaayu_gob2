package main

import (
	"fmt"
	"strings"
)

// Struct 1
type Employee struct {
	name     string
	age      int
	division string
}

// Struct 2
type Person struct {
	name string
	age  int
}

// Struct 3
type eployz struct {
	division string
	person   Person
}

func main() {
	// Struct (giving value to struct)
	var employee Employee

	employee.name = "Mark"

	employee.age = 23

	employee.division = "Singer"

	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)
	fmt.Println("=============================================")

	// Struct (initializing struct)
	var employee1 = Employee{}
	employee1.name = "Mark"

	employee1.age = 23

	employee1.division = "Singer"

	var employee2 = Employee{name: "Mark", age: 23, division: "Singer"}

	fmt.Printf("Employee1 %+v\n", employee1)
	fmt.Printf("Employee2 %+v\n", employee2)
	fmt.Println("=============================================")

	// Struct (Pointer to a struct)
	var dataEmployee1 = Employee{name: "Mark", age: 23, division: "Singer"}

	var dataEmployee2 *Employee = &dataEmployee1

	fmt.Println("Employee1 name:", dataEmployee1.name)
	fmt.Println("Employee2 name:", dataEmployee2.name)

	fmt.Println(strings.Repeat("#", 50))

	dataEmployee2.name = "Lee"

	fmt.Println("Employee1 name:", dataEmployee1.name)
	fmt.Println("Employee2 name:", dataEmployee2.name)
	fmt.Println("=============================================")

	// Struct (embedded struct)
	var employee3 = eployz{}

	employee3.person.name = "Taeyeong"
	employee3.person.age = 22
	employee3.division = "Backend Developer"

	fmt.Printf("%+v\n", employee3)
	fmt.Println("=============================================")

	// Struct (Anonymous struct)

	var emply1 = struct {
		person   Person
		division string
	}{}

	emply1.person = Person{name: "Mark", age: 23}
	emply1.division = "Singer"

	//Anonymous struct dengan pengisian field

	var emply2 = struct {
		person   Person
		division string
	}{
		person:   Person{name: "Jaehyun", age: 23},
		division: "Dancer",
	}

	fmt.Printf("Employee1: %+v\n", employee1)
	fmt.Printf("Employee2: %+v\n", emply2)
	fmt.Println("=============================================")

	// Struct (Slice of struct)
	var people = []Person{
		{name: "Mark", age: 23},
		{name: "Jaehyun", age: 26},
		{name: "Taeyeong", age: 28},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}
	fmt.Println("=============================================")

	// Struct (slice of anonymous struct)
	var emp = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "Mark", age: 23}, division: "Singer"},
		{person: Person{name: "Jaehyun", age: 26}, division: "Dancer"},
		{person: Person{name: "Taeyeing", age: 28}, division: "Leader"},
	}

	for _, v := range emp {
		fmt.Printf("%+v\n", v)
	}
	fmt.Println("=============================================")
}
