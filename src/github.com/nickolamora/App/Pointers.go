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

	fmt.Println("================")
	/**
	This is when to or why use pointers in these situations
	1. Are good when you have a large chunk of data so that the only thing you’re passing around is the address of that struct
	2. Another way is when you need to change something that’s at a certain location
	*/
	x := 0

	fmt.Println("x before ", x)
	pointerTest(&x)
	fmt.Println("x after ", x)
}

func pointerTest(y *int) {
	fmt.Println(y)
	//the value of the address of y, change it to 43
	*y = 43
	fmt.Println(y)
}
