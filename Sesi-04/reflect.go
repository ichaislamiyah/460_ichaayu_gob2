package main

import (
	"fmt"
	"reflect"
)

func (s *student) SetName(name string) {
	s.Name = name
}

type student struct {
	Name  string
	Grade int
}

func main() {
	//Identifying Data Type and Value
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe variabel :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel :", reflectValue.Int())
	}

	fmt.Println("===================================================")

	//Acessing Value Using Interface Method
	/*
		var angka = 25
		var reflectValue = reflect.ValueOf(angka)

		fmt.Println("tipe variabel :", reflectValue.Type())
		fmt.Println("nilai variabel :", reflectValue.Interface())
		fmt.Println("===================================================")

		//Identifying Method Information
		var s1 = &student{Name: "Kang Seungyoon", Grade: 6}
		fmt.Println("name :", s1.Name)

		var reflectValue = reflect.ValueOf(s1)
		var method = reflectValue.MethodByName("SetName")
		method.Call([]reflect.Value{
			reflect.ValueOf("Seungyoon"),
		})

		fmt.Println("name :", s1.Name)
		fmt.Println("===================================================")*/
}
