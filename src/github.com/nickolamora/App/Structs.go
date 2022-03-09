package main

import "fmt"

// struct is a collection of different datatypes (like a Java class)?

//naming conventions apply to structs and their properties
//structs names in camelcase 'struct' are only scoped within a package
//structs names in pascal 'Struct' can be exported to multiple packages because they are in the global scope

type Student struct {
	Name    string
	Id      int
	Classes []string
}

func main() {
	student1 := Student{
		Name: "Nick",
		Id:   123,
		Classes: []string{
			"art",
			"math",
		},
	}
	fmt.Println(student1.Name)
	student1.Name = "Bob"
	fmt.Println(student1.Name)

	//structs are value types theyll pass on the values not the address but if you want to pass the pointer just
	// add &student1
	student2 := student1
	fmt.Println(student2.Name)

}
