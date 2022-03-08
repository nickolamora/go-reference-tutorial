package main

import "fmt"

//Once constants are created you can't change them
//Constants don't need to be used once you initialized them

//constants names in camelcase 'a_CONSTANT' are only scoped within a package
//constants names in pascal 'A_CONSTANT' can be exported to multiple packages because they are in the global scope

//These constants are globally scoped
const (
	A string = "something"
	B string = "else"
)

//iota is an incrementer  So essentially it can be used to create effective constant in Go .
const (
	a = iota
	// used to skip a value
	_
	b
	c
)

func main() {
	const i int = 12
	fmt.Println(i)
}
