package main

import "fmt"

// A pointer points to a specific place in memory where the value of a variable is stored

type Foo struct {
	bar int
}

func main() {
	var a int = 12
	var b *int = &a
	fmt.Println(a)
	// this will only print the memory address that 'a' is in. To get the value in that memory address you need to DEREFERENCE
	fmt.Println(b)
	// dereference the value
	fmt.Println(*b)

	fmt.Println("================")

	//struct example using pointers

	//created a memory address
	var foo *Foo
	//assign a memory address to our struct
	foo = new(Foo)
	//should print the address and value of our empty/default struct
	fmt.Println(foo)
	//dereference
	(*foo).bar = 46
	// or go compiler is smart enough to dereference by itself
	foo.bar = 65
	fmt.Println((*foo).bar)

}
