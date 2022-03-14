package main

import (
	"fmt"
)

type rectangle struct {
	width, height int
}

//method for a rectangle that takes in a rectangle struct and returns an int
func (r rectangle) area() int {
	return r.width * r.height
}

func main() {

	writeMessage("This is a message")
	fmt.Println("================")

	//Here we pass the address to a method expecting a pointer
	msg := "Hello world"
	writeMessageWithPointers(&msg)

	fmt.Println("================")

	//a variadic function is a function that accepts zero, one, or more values as a single argument
	// p.s. if adding more parameters, the variadic set needs to be the last parameter
	total := variadicFunction(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(total)

	fmt.Println("================")

	totalD, error := divide(4, 2)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(totalD)

	fmt.Println("================")

	// anonymous functions
	fun := func() {
		fmt.Println("I am an anonymous function")
		//if you're not assigning it to a variable and just stating it - it needs to end with ()
	}
	// you cannot call a function before its declaration, however if its defined outside of the main()
	//then you can call it anywhere since it will be a 'proper function'
	fun()

	fmt.Println("================")

	//Methods  - special functions that work with a certain data type (Struct)
	rec := rectangle{
		width:  20,
		height: 10,
	}
	fmt.Println("The area is ", rec.area())
}

func writeMessage(msg string) {
	fmt.Println(msg)
}

func writeMessageWithPointers(msg *string) {

	//when getting or setting the value, you have to dereference
	*msg = "Hello Aliens"
	fmt.Println(*msg)
}

//the last int means this function returns an int
func variadicFunction(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

// If you have multiple parameters of the same type you can state them this way
//also this is how you can validate parameters and return a value and error
func divide(a, b float64) (float64, error) {
	if b == 0.0 || b == 0 {
		return 0.0, fmt.Errorf("Cannot divide by 0")
	}
	return a / b, nil
}
