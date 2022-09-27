package main

import (
	"fmt"
	"math"
)

// interface
type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func print(t string, s shape) {
	fmt.Printf("%s area %v\n", t, s.area())
}

func (c circle) volume() float64 { //Interface (type assertion)
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func main() {
	var c1 shape = circle{radius: 5}
	var r1 shape = rectangle{width: 3, height: 2}

	fmt.Printf("Type of c1: %T\n", c1)
	fmt.Printf("Type of r1: %T\n", r1)
	fmt.Println("=============================================")

	var a1 shape = circle{radius: 5}
	var b1 shape = rectangle{width: 3, height: 2}

	fmt.Println("Circle area", a1.area())
	fmt.Println("Circle perimeter", a1.perimeter())

	fmt.Println("Rectangle area", b1.area())
	fmt.Println("Rectangle perimeter", b1.perimeter())
	fmt.Println("=============================================")

	var e1 shape = circle{radius: 5}
	var d1 shape = rectangle{width: 3, height: 2}

	print("Rectangle", e1)
	print("Circle", d1)
	fmt.Println("=============================================")

	var f1 shape = circle{radius: 8}

	value, ok := f1.(circle)

	if ok == true {
		fmt.Printf("Circle value: %+v\n", value)
		fmt.Printf("Circle volume: %f\n", value.volume())
		fmt.Println("=============================================")
	}
}
