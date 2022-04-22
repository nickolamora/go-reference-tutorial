package main

import (
	"fmt"
	"strconv"
)

// This area outside of the main func is the global scope

func main() {
	//Variable Init

	/**
	var anIntVariable int
	anIntVariable = 1
	OR
	var anIntVariable = 1
	OR
	*/

	anIntVariable := 60
	anotherVariable := anIntVariable
	//Print out the value and the of a variablev
	fmt.Printf("%v, %T\n\n", anIntVariable, anIntVariable)
	fmt.Printf("%v, %T\n\n", anotherVariable, anotherVariable)

	/**
	IMPORTANT
	----------
	variables names in camelcase 'aVariableTest' are only scoped within a package
	variables names in pascal 'AVariableTest' can be exported to multiple packages

	if a variable is created but not assigned value, they will be assigned a default value and type
	*/

	//Type casting
	aFloatVariable := float32(anIntVariable)
	fmt.Printf("%v, %T\n\n", aFloatVariable, aFloatVariable)

	//casting as a string this way will get the ASCII of that int, if you want the value, use strconv
	aStringVariable := string(anIntVariable)
	fmt.Printf("%v, %T\n\n", aStringVariable, aStringVariable)

	aStringVariable = strconv.Itoa(anIntVariable)
	fmt.Printf("%v, %T\n\n", aStringVariable, aStringVariable)

	aBoolean := true
	fmt.Printf("%v, %T\n\n", aBoolean, aBoolean)

}
