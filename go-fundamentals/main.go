package main

import "fmt"

func main() {
	fmt.Println("           GoLang Fundamentals...        ")
	//variableFuncWithGreeting("Samuel Turay", "Sierra Leone")
	//userId := userInputFunction()
	//fmt.Println(userId)

	// pointer
	//i := 10
	//fmt.Printf("%T %v \n", &i, &i)
	//fmt.Printf("%T %v \n", *(&i), *(&i))

	c := Circle{radius: 5}
	c.calcArea()
	fmt.Printf("%+v \n", c)

	r := rect{width: 10, height: 5}
	d := square{side: 5}
	printData(r)
	printData(d)
}

type Circle struct {
	radius float64
	area   float64
}

func (c *Circle) calcArea() {
	c.area = 3.14 * c.radius * c.radius
}

type shape interface {
	area() float64
	perimeter() float64
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (s square) perimeter() float64 {
	return 4 * s.side
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func printData(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}
