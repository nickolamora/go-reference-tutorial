package main

import "fmt"

/**
Go Lang supports composition relationship in structs 'has a' A computer has a CPU
as opposed to 'is a' relationship you get with inheritance.

Since Golang does not support classes, so inheritance takes place through struct embedding.
We cannot directly extend structs but rather use a concept called composition where the struct is used to form other objects.
So, you can say there is No Inheritance Concept in Golang
*/

type Processor struct {
	name  string
	cores int
}

type Memory struct {
	capacity   int
	memoryType string
}

type Computer struct {
	Processor
	Memory
	price int
}

func main() {

	computer := Computer{}
	computer.price = 1000
	computer.name = "M1"
	computer.cores = 12
	computer.capacity = 32
	computer.memoryType = "DDR5"
	fmt.Println(computer)
}
