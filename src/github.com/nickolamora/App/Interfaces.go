package main

import "fmt"

//Interfaces - a way to establish a contract an object must meet

func main() {
	var r geometry = rectangle2{
		height: 2,
		width:  3,
	}
	fmt.Println(r.area())
	fmt.Println(r.perimeter())
}

//defining an interface
type geometry interface {
	//defining the methods and what they return
	area() float64
	perimeter() float64
}

//implementing an interface

type rectangle2 struct {
	height, width float64
}

func (r rectangle2) area() float64 {
	return r.width * r.height
}

func (r rectangle2) perimeter() float64 {
	return 2*r.width + 2*r.height
}
