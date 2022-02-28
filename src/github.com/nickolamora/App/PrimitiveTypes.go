package main

import "fmt"

func main() {

	//boolean
	foo := false
	fmt.Printf("%v, %T\n\n", foo, foo)

	//strings
	s := "This is a string"
	s1 := "This is another string"
	fmt.Printf("%v, %T\n\n", s+s1, s)
	fmt.Printf("%v, %T\n\n", string(s[1]), s)
	//Strings are saved as a byte array in go
	sArr := []byte(s)
	fmt.Printf("%v, %T\n\n", sArr, sArr)
	//rune
	r := 'R'
	fmt.Printf("%v, %T\n\n", r, r)

	//numerics
	i := 12
	j := 10
	fmt.Printf("%v, %T\n\n", i, i)
	fmt.Printf("%v, %T\n\n", j, j)
	fmt.Println(i + j)
	fmt.Println(i - j)
	fmt.Println(i * j)
	fmt.Println(i / j)
	fmt.Println(i % j)
	//binary operators
	// TODO brush up on binary operators
	fmt.Println(i & j)
	fmt.Println(i | j)
	fmt.Println(i ^ j)
	fmt.Println(i &^ j)
	//floating point
	d := 3.14
	fmt.Printf("%v, %T\n\n", d, d)
	t := 3.1e4
	fmt.Printf("%v, %T\n\n", t, t)
	//complex
	// TODO brush up on complex numbers
	var c complex128 = 1 + 2i
	fmt.Printf("%v, %T\n\n", c, c)
	fmt.Printf("%v, %T\n\n", real(c), imag(c))
	var cC = complex(3, 4)
	fmt.Printf("%v, %T\n\n", real(cC), imag(cC))
	fmt.Println(c + cC)
	fmt.Println(c - cC)
	fmt.Println(c * cC)
	fmt.Println(c / cC)

}
