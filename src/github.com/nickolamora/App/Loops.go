package main

import (
	"fmt"
)

/**
In go there is only one loop: for loop
to do any kind of other kind of loop we need to hack the for loop
*/

func main() {

	//for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("================")

	// for loop with multiple variables
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}
	fmt.Println("================")

	//or
	i, j := 0, 0
	for i < 5 {
		i, j = i+1, j+1
		if i == 2 && j == 2 {
			continue
		}
		fmt.Println(i, j)
	}
	fmt.Println("================")

	//while
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	fmt.Println("================")
	//do while loop
	x, y := 0, 0
	for {
		fmt.Println(x, y)
		x, y = x+1, y+1
		if x == 2 && y == 2 {
			// break - stops further execution of a loop construct.
			// continue - skips the rest of the loop body and starts the next iteration.
			break

		}
	}
	fmt.Println("================")

	//nested loops
	for i := 0; i < 5; i = i + 1 {
		for j := 0; j < 5; j = j + 1 {
			fmt.Println(i, j)
		}
	}

	fmt.Println("================")

	//Loop through collections

	//array
	a := []int{1, 2, 3, 4, 5, 6}
	// index, value or ignore by using _
	for k, v := range a {
		fmt.Println(k, v)
	}
	fmt.Println("================")

	//Map
	shoppingCart := make(map[string]int)
	shoppingCart = map[string]int{
		"keyboard": 100,
		"mouse":    50,
		"laptop":   1000,
	}

	for k, v := range shoppingCart {
		fmt.Println(k, v)
	}
}
