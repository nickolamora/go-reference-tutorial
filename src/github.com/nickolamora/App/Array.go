package main

import "fmt"

//Array sizes are fixed

func main() {
	//var list [3]int = [3]int{1, 2, 3}
	list := [3]int{1, 2, 3}
	fmt.Printf("Amount: %v\n", list)
	fmt.Printf("Size: %v\n", len(list))

	//if you don't know the size of the array, whatever you add will determine the size
	arr := [...]int{2, 3}
	fmt.Printf("Amount: %v\n", arr)
	//size is 2
	fmt.Printf("Size: %v\n", len(arr))
	//change the value of a specific index in the arr
	arr[0] = 8
	fmt.Printf("Size: %v\n", arr)

	//copy data from one array to another
	anotherArr := arr
	fmt.Printf("Amount: %v\n", anotherArr)

	//point a new array to the existing memory address of another - Pointers!
	// 	arrPointingToAnother points to arr
	arrPointingToAnother := &arr
	arr[0] = 420
	fmt.Printf("Amount: %v\n", arr)
	fmt.Printf("Amount: %v\n", arrPointingToAnother)
	arrPointingToAnother[0] = 69
	fmt.Printf("Amount: %v\n", arrPointingToAnother)

	//slicing an array
	a := [...]int{1, 2, 3, 4, 5}
	b := a[1:3]
	c := a[:len(a)-1]
	fmt.Printf("Amount: %v\n", b)
	fmt.Printf("Amount: %v\n", c)

	//multidimensional array
	var indentityMatrix [3][3]int = [3][3]int{
		[3]int{1, 9, 0},
		[3]int{0, 1, 0},
		[3]int{0, 0, 1},
	}
	fmt.Println(indentityMatrix)
	fmt.Println(indentityMatrix[0][1])

}
