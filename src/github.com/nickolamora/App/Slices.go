package main

import "fmt"

//overcoming the limit of fixed size arrays  - use Slices! is a data type that internally uses arrays.
//When you create a slice it creates an array internally that copies over the data and adds any new data you want. The slice POINTS to the internal array

//arrays copy data when copying from another array
//slices point to data

func main() {

	//creating a slice
	var a []int = []int{1, 2, 3}
	var b []int = a
	fmt.Println(a)
	fmt.Println(b)
	b[0] = 4
	fmt.Println(a)
	fmt.Println(b)

	//slice functions
	fmt.Println(len(b))
	fmt.Println(cap(b))

	//making a slice (type, length, capacity)
	var s []int = make([]int, 3, 10)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	//appending a slice
	var x []int = []int{1, 2, 3}
	fmt.Println(x)
	var y []int = append(x, 6, 7, 5, 3, 0, 9)
	fmt.Println(y)
	var z []int = append(y[1:4], 9, 9, 9, 9, 9, 9)
	fmt.Println(z)
	//append data of x to y and put it in z (can use spread ... to get the values)
	var r []int = append(y, x...)
	fmt.Printf("%v, %T\n\n", r, r)

}
